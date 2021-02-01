package comparator

import "bytes"

type Comparator interface {
	Compare(rhs, lhs interface{}) int
	Name() string
}

var DefaultComparator Comparator = defaultCmp{}

type defaultCmp struct{}

func (defaultCmp) Compare(rhs, lhs interface{}) int {
	return bytes.Compare(rhs.([]byte), lhs.([]byte))
}

func (defaultCmp) Name() string {
	return "BytewiseComparator"
}

type Int32Cmp struct {}

func (Int32Cmp) Compare(rhs, lhs interface{}) int {
	rhsInt := rhs.(int32)
	lhsInt := lhs.(int32)

	switch res := rhsInt-lhsInt; {
	case res == 0:
		return 0
	case res > 0:
		return 1
	default:
		return -1
	}
}

func (Int32Cmp) Name() string{
	return "Int32Comparator"
}

type Int64Cmp struct {}

func (Int64Cmp) Compare(rhs, lhs interface{}) int {
	rhsInt := rhs.(int64)
	lhsInt := lhs.(int64)

	switch res := rhsInt-lhsInt; {
	case res == 0:
		return 0
	case res > 0:
		return 1
	default:
		return -1
	}
}

func (Int64Cmp) Name() string{
	return "Int32Comparator"
}