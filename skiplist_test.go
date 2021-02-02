package skiplist

import (
	"reflect"
	"skiplist-go/comparator"
	"sync"
	"testing"
)

func TestNewSkiplist(t *testing.T) {
	type args struct {
		maxlevel int
		cmp      comparator.Comparator
	}
	tests := []struct {
		name string
		args args
		want *Skiplist
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSkiplist(tt.args.maxlevel, tt.args.cmp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSkiplist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkiplist_Contains(t *testing.T) {
	s := NewSkiplist(10, nil)
	s.Set([]byte("aaaaa"), []byte("aaaaa"))
	s.Set([]byte("aaaba"), []byte("aaaba"))
	s.Set([]byte("ccccc"), []byte("ccccc"))
	s.Set([]byte(""), []byte(""))
	s.Set([]byte("a"), []byte("a"))
	s.Set([]byte(" "), []byte(" "))

	type args struct {
		key interface{}
	}
	tests := []struct {
		name   string
		args   args
		want   bool
	}{
		{name: "case1", args: args{key: []byte("aaaaa")}, want: true},
		{name: "case2", args: args{key: []byte("aaaa")}, want: false},
		{name: "case3", args: args{key: []byte("")}, want: true},
		{name: "case4", args: args{key: []byte("a")}, want: true},
		{name: "case5", args: args{key: []byte("#$$")}, want: false},
		{name: "case6", args: args{key: []byte("\\0")}, want: false},
		{name: "case7", args: args{key: []byte(" ")}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.Contains(tt.args.key); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkiplist_Delete(t *testing.T) {

	s := NewSkiplist(10, nil)
	s.Set([]byte("aaaaa"), []byte("aaaaa"))
	s.Set([]byte("aaaba"), []byte("aaaba"))
	s.Set([]byte("ccccc"), []byte("ccccc"))
	s.Set([]byte(""), []byte(""))
	s.Set([]byte("a"), []byte("a"))
	s.Set([]byte("abc"), []byte("abc"))
	s.Set([]byte("ddd"), []byte("ddd"))
	s.Set([]byte("$$$"), []byte("$$$"))
	s.Set([]byte(" "), []byte(" "))

	type args struct {
		key interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{name: "case1", args: args{key: []byte("a")}, wantErr: nil},
		{name: "case2", args: args{key: []byte("")}, wantErr: nil},
		{name: "case3", args: args{key: []byte(" ")}, wantErr: nil},
		{name: "case4", args: args{key: []byte("b")}, wantErr: ErrNotFound},
		{name: "case5", args: args{key: []byte("cced")}, wantErr: ErrNotFound},
		{name: "case6", args: args{key: []byte("$")}, wantErr: ErrNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Delete(tt.args.key); err != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSkiplist_Get(t *testing.T) {
	type fields struct {
		level    int
		maxLevel int
		length   int
		cmp      comparator.Comparator
		root     *node
		nodes    []*node
		mutex    sync.RWMutex
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Skiplist{
				level:    tt.fields.level,
				maxLevel: tt.fields.maxLevel,
				length:   tt.fields.length,
				cmp:      tt.fields.cmp,
				root:     tt.fields.root,
				nodes:    tt.fields.nodes,
				mutex:    tt.fields.mutex,
			}
			got, err := s.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkiplist_Len(t *testing.T) {
	type fields struct {
		level    int
		maxLevel int
		length   int
		cmp      comparator.Comparator
		root     *node
		nodes    []*node
		mutex    sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Skiplist{
				level:    tt.fields.level,
				maxLevel: tt.fields.maxLevel,
				length:   tt.fields.length,
				cmp:      tt.fields.cmp,
				root:     tt.fields.root,
				nodes:    tt.fields.nodes,
				mutex:    tt.fields.mutex,
			}
			if got := s.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkiplist_Level(t *testing.T) {
	type fields struct {
		level    int
		maxLevel int
		length   int
		cmp      comparator.Comparator
		root     *node
		nodes    []*node
		mutex    sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Skiplist{
				level:    tt.fields.level,
				maxLevel: tt.fields.maxLevel,
				length:   tt.fields.length,
				cmp:      tt.fields.cmp,
				root:     tt.fields.root,
				nodes:    tt.fields.nodes,
				mutex:    tt.fields.mutex,
			}
			if got := s.Level(); got != tt.want {
				t.Errorf("Level() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkiplist_MaxLevel(t *testing.T) {
	type fields struct {
		level    int
		maxLevel int
		length   int
		cmp      comparator.Comparator
		root     *node
		nodes    []*node
		mutex    sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Skiplist{
				level:    tt.fields.level,
				maxLevel: tt.fields.maxLevel,
				length:   tt.fields.length,
				cmp:      tt.fields.cmp,
				root:     tt.fields.root,
				nodes:    tt.fields.nodes,
				mutex:    tt.fields.mutex,
			}
			if got := s.MaxLevel(); got != tt.want {
				t.Errorf("MaxLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkiplist_Set(t *testing.T) {
	type fields struct {
		level    int
		maxLevel int
		length   int
		cmp      comparator.Comparator
		root     *node
		nodes    []*node
		mutex    sync.RWMutex
	}
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Skiplist{
				level:    tt.fields.level,
				maxLevel: tt.fields.maxLevel,
				length:   tt.fields.length,
				cmp:      tt.fields.cmp,
				root:     tt.fields.root,
				nodes:    tt.fields.nodes,
				mutex:    tt.fields.mutex,
			}
			if err := s.Set(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSkiplist_randLevel(t *testing.T) {
	type fields struct {
		level    int
		maxLevel int
		length   int
		cmp      comparator.Comparator
		root     *node
		nodes    []*node
		mutex    sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Skiplist{
				level:    tt.fields.level,
				maxLevel: tt.fields.maxLevel,
				length:   tt.fields.length,
				cmp:      tt.fields.cmp,
				root:     tt.fields.root,
				nodes:    tt.fields.nodes,
				mutex:    tt.fields.mutex,
			}
			if got := s.randLevel(); got != tt.want {
				t.Errorf("randLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newData(t *testing.T) {
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want *Data
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newData(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newNode(t *testing.T) {
	type args struct {
		level int
		data  *Data
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newNode(tt.args.level, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_Key(t *testing.T) {
	type fields struct {
		data *Data
		next []*node
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				data: tt.fields.data,
				next: tt.fields.next,
			}
			if got := n.Key(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_NextWithLevel(t *testing.T) {
	type fields struct {
		data *Data
		next []*node
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				data: tt.fields.data,
				next: tt.fields.next,
			}
			if got := n.NextWithLevel(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextWithLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_Value(t *testing.T) {
	type fields struct {
		data *Data
		next []*node
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				data: tt.fields.data,
				next: tt.fields.next,
			}
			if got := n.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}