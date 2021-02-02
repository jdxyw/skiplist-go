package skiplist

import (
	"reflect"
	"skiplist-go/comparator"
	"sync"
	"testing"
)

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
		name string
		args args
		want bool
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
	s := NewSkiplist(10, nil)
	s.Set([]byte("aaaaa"), []byte("bbbb"))
	s.Set([]byte("aaaba"), []byte("cccc"))
	s.Set([]byte("ccccc"), []byte("dddd"))
	s.Set([]byte(""), []byte("abc"))
	s.Set([]byte("a"), []byte("a"))
	s.Set([]byte("abc"), []byte("def"))
	s.Set([]byte("ddd"), []byte("fgb"))
	s.Set([]byte("$$$"), []byte("###"))
	s.Set([]byte(" "), []byte(" "))

	type args struct {
		key interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr error
	}{
		{name: "case1", args: args{key: []byte("a")}, want: []byte("a"), wantErr: nil},
		{name: "case2", args: args{key: []byte("")}, want: []byte("abc"), wantErr: nil},
		{name: "case3", args: args{key: []byte(" ")}, want: []byte(" "), wantErr: nil},
		{name: "case4", args: args{key: []byte("ccccc")}, want: []byte("dddd"), wantErr: nil},
		{name: "case5", args: args{key: []byte("cced")}, want: []byte("a"), wantErr: ErrNotFound},
		{name: "case6", args: args{key: []byte("$")}, want: []byte("a"), wantErr: ErrNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Get(tt.args.key)
			if err != nil && tt.wantErr == ErrNotFound {
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