package skiplist

import "bytes"

// The comparator interface.
// rhs > lhs : 1
// rhs == lhs : 0
// rhs < lhs : -1
// In this package, only provides the BytewiseComparator comparator.
// If the key is int32 or int64, you can define the Int32Comparator by yourself.
type Comparator interface {
	Compare(rhs, lhs interface{}) int
	Name() string
}

// Return the default comparator which is the BytewiseComparator.
func GetDefaultComparator() Comparator {
	return BytewiseComparator{}
}

type BytewiseComparator struct{}

// Compare compare two byte slices for BytewiseComparator.
func (BytewiseComparator) Compare(rhs, lhs interface{}) int {
	return bytes.Compare(rhs.([]byte), lhs.([]byte))
}

// Name returns the name of current comparator.
func (BytewiseComparator) Name() string {
	return "BytewiseComparator"
}
