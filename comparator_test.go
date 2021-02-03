package skiplist

import (
	"testing"
)

func Test_defaultCmp_Compare(t *testing.T) {
	type args struct {
		rhs []byte
		lhs []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "BytewiseComparator", args: args{rhs: []byte("aaaa"), lhs: []byte("bbbb")}, want: -1},
		{name: "BytewiseComparator", args: args{rhs: []byte("aaaa"), lhs: []byte("aaaa")}, want: 0},
		{name: "BytewiseComparator", args: args{rhs: []byte("cccc"), lhs: []byte("bbbb")}, want: 1},
		{name: "BytewiseComparator", args: args{rhs: []byte("aaaa"), lhs: nil}, want: 1},
		{name: "BytewiseComparator", args: args{rhs: nil, lhs: nil}, want: 0},
		{name: "BytewiseComparator", args: args{rhs: []byte("aaaa"), lhs: []byte("")}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			de := GetDefaultComparator()
			if got := de.Compare(tt.args.rhs, tt.args.lhs); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if de.Name() != tt.name {
				t.Errorf("invalid comparator name")
			}
		})
	}
}
