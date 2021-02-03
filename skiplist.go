package skiplist

import (
	"errors"
	"math/rand"
	"sync"
)

const (
	kMaxHeight = 12
	kBranching = 4
)

var (
	// ErrNotFound returns if the key is not found in current list.
	ErrNotFound = errors.New("key not found")
)

// The Data load in each node of list.
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

func newNode(level int, data *Data) *node {
	return &node{
		data: data,
		next: make([]*node, level, level),
	}
}

// Skiplist is the struct of skiplist.
type Skiplist struct {
	level    int
	maxLevel int
	length   int
	cmp      Comparator
	root     *node
	mutex    sync.RWMutex
}

// NewSkiplist creates a skiplist. maxlevel specific the max level this skiplist would support.
// cmp is a Comparator object. If nil, the skiplist would use the default Comparator - BytewiseComparator.
// It supports all kinds of type as long as you define it.
func NewSkiplist(maxlevel int, cmp Comparator) *Skiplist {
	newCmp := cmp
	if newCmp == nil {
		newCmp = GetDefaultComparator()
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
		cmp:      newCmp,
	}
}

// Len returns the length of current skiplist.
func (s *Skiplist) Len() int {
	return s.length
}

// Level returns the current level of current skiplist.
func (s *Skiplist) Level() int {
	return s.level
}

// MaxLevel returns the max level allowed by skiplist.
func (s *Skiplist) MaxLevel() int {
	return s.maxLevel
}

func (s *Skiplist) getGreaterOrEqual(key interface{}) (*node, []*node) {
	prevs := make([]*node, s.maxLevel, s.maxLevel)
	n := s.root
	for i := s.level - 1; i >= 0; i-- {
		for n.next[i] != nil && s.cmp.Compare(n.next[i].data.key, key) < 0 {
			n = n.next[i]
		}
		prevs[i] = n
	}

	return n, prevs
}

// Get returns the corresponding value for specific value. error would be nil if the key is exist in the list.
func (s *Skiplist) Get(key interface{}) (interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	n, _ := s.getGreaterOrEqual(key)

	n = n.next[0]

	if n != nil && s.cmp.Compare(n.data.key, key) == 0 {
		return n.data.value, nil
	}

	return nil, ErrNotFound
}

// Contains return true if the key is exist in the list.
func (s *Skiplist) Contains(key interface{}) bool {
	if _, err := s.Get(key); err == nil {
		return true
	}

	return false
}

// Set would insert a new node into the list with key/value. If the key has been exist
// in this list, the value would be updated.
func (s *Skiplist) Set(key, value interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	n, prevs := s.getGreaterOrEqual(key)

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

// Delete would remove specific node from the list based on the key.
func (s *Skiplist) Delete(key interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	head := s.root
	n, prevs := s.getGreaterOrEqual(key)

	n = n.next[0]
	if n == nil || s.cmp.Compare(n.data.key, key) != 0 {
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

// randLevel returns how many level the new node would use.
func (s *Skiplist) randLevel() int {
	h := 1
	for h < s.maxLevel && rand.Intn(kBranching) == 0 {
		h += 1
	}

	return h
}
