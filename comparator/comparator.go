package comparator

import "bytes"

type Comparator interface {
	Compare(rhs, lhs []byte) int
	Name() string
}

var DefaultComparator Comparator = defaultCmp{}

type defaultCmp struct{}

func (defaultCmp) Compare(rhs, lhs []byte) int {
	return bytes.Compare(rhs, lhs)
}

func (defaultCmp) Name() string {
	return "BytewiseComparator"
}