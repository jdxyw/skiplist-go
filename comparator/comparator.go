package comparator

import "bytes"

type Comparator interface {
	Compare(rhs, lhs interface{}) int
	Name() string
}

func GetDefaultComparator() Comparator {
	return BytewiseComparator{}
}

type BytewiseComparator struct{}

func (_ BytewiseComparator) Compare(rhs, lhs interface{}) int {
	return bytes.Compare(rhs.([]byte), lhs.([]byte))
}

func (_ BytewiseComparator) Name() string {
	return "BytewiseComparator"
}