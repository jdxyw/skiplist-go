package skiplist

import (
	"errors"
	"math/rand"
	"sync"

	"skiplist-go/comparator"
)

const (
	kMaxHeight = 12
	kBranching = 4
)

var (
	ErrNotFound = errors.New("Key not found")
)

type Data struct {
	key   interface{}
	value interface{}
}

func newData(key, value interface{}) *Data {
	return &Data{
		key:   key,
		value: value,
	}
}

type node struct {
	data *Data
	next []*node
}

func (n *node) Key() interface{} {
	return n.data.key
}

func (n *node) Value() interface{} {
	return n.data.value
}

func (n *node) NextWithLevel(i int) *node {
	return n.next[i]
}

func newNode(level int, data *Data) *node {
	return &node{
		data: data,
		next: make([]*node, level, level),
	}
}

type Skiplist struct {
	level    int
	maxLevel int
	length   int
	cmp      comparator.Comparator
	root     *node
	nodes    []*node
	mutex    sync.RWMutex
}

func NewSkiplist(maxlevel int, cmp comparator.Comparator) *Skiplist {
	newCmp := cmp
	if newCmp == nil {
		newCmp = comparator.DefaultComparator
	}

	if maxlevel < 4 {
		panic("skiplist: the maxlevel should be larger than 4!")
	}

	if maxlevel > kMaxHeight {
		maxlevel = kMaxHeight
	}

	return &Skiplist{
		maxLevel: maxlevel,
		root:     newNode(maxlevel, nil),
		nodes:    make([]*node, maxlevel, maxlevel),
		cmp:      newCmp,
	}
}

func (s *Skiplist) Len() int {
	return s.length
}

func (s *Skiplist) Level() int {
	return s.level
}

func (s *Skiplist) MaxLevel() int {
	return s.maxLevel
}

func (s *Skiplist) Get(key interface{}) (interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	n := s.root
	for i := s.level - 1; i >= 0; i-- {
		for n.next[i] != nil && s.cmp.Compare(n.next[i].data.key, key) >= 0 {
			n = n.next[i]
		}
	}

	n = n.next[0]

	if n != nil && n.data.key == key {
		return n.data.value, nil
	}

	return nil, ErrNotFound
}

func (s *Skiplist) Contains(key interface{}) bool {
	if s.Get(key) != nil {
		return true
	}

	return false
}

func (s *Skiplist) Set(key, value interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	prevs := make([]*node, s.maxLevel, s.maxLevel)
	n := s.root
	for i := s.level - 1; i >= 0; i-- {
		for n.next[i] != nil && s.cmp.Compare(n.next[i].data.key, key) >= 0 {
			n = n.next[i]
		}
		prevs[i] = n
	}

	// The key is already in this list, just update the value field.
	if n.next[0] != nil && s.cmp.Compare(n.next[0].data.key, key) == 0 {
		n.next[0].data.value = value
		return nil
	}

	level := s.randLevel()

	if level > s.level {
		for i := s.level; i < level; i++ {
			prevs[i] = s.root
		}
		s.level = level
	}

	newNode := newNode(level, newData(key, value))
	for i := 0; i < level; i++ {
		newNode.next[i] = prevs[i].next[i]
		prevs[i].next[i] = newNode
	}
	s.length += 1
	return nil
}

func (s *Skiplist) Delete(key interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	prevs := make([]*node, s.maxLevel, s.maxLevel)
	head := s.root

	n := head
	for i := s.level - 1; i >= 0; i-- {
		for n.next[i] != nil && s.cmp.Compare(n.next[i].data.key, key) >= 0 {
			n = n.next[i]
		}
		prevs[i] = n
	}

	n = n.next[0]
	if n == nil || s.cmp.Compare(n.next[i].data.key, key) != 0 {
		return ErrNotFound
	}

	for i := 0; i < s.level; i++ {
		if prevs[i].next[i] == n {
			prevs[i].next[i] = n.next[i]
		}
	}

	for s.level > 1 && head.next[s.level-1] == nil {
		s.level -= 1
	}

	s.length -= 1
	return nil
}

func (s *Skiplist) randLevel() int {
	h := 1
	for h < s.maxLevel && rand.Intn(kBranching) == 0 {
		h += 1
	}

	return h
}
