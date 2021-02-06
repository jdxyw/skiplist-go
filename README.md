# skiplist-go
![Go](https://github.com/jdxyw/skiplist-go/workflows/Go/badge.svg?branch=main)
[![GoDoc](https://godoc.org/github.com/jdxyw/skiplist-go?status.svg)](https://godoc.org/github.com/jdxyw/skiplist-go)
[![codecov](https://codecov.io/gh/jdxyw/skiplist-go/branch/main/graph/badge.svg?token=BK9VMLZKHI)](https://codecov.io/gh/jdxyw/skiplist-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/jdxyw/skiplist-go)](https://goreportcard.com/report/github.com/jdxyw/skiplist-go)

Skip list is a probabilistic data structure that allows $O(log n)$ search complexity as well as $O(log n)$ insertion complexity within an ordered sequence of $n$ elements.

This package implements a genetic skip list with `interface{}` type. You could define your expected type you want.
# Install

Install this package through `go get`.

```bash
    go get github.com/jdxyw/skiplist-go
```

# Basic usage

The usage is quite simple.

```go
package main

import (
	"fmt"
	"github.com/jdxyw/skiplist-go"
)

func main() {
	// If you pass the nil to `cmp` parameter, which would use the default comparactor (Bytes wise).
	s := skiplist.NewSkiplist(10, nil)

	// Use the Set to insert/update element in this list.
	// The `value` could be nil.
	s.Set([]byte("Hello"), []byte("world"))
	s.Set([]byte("Python"), []byte("Perl"))
	s.Set([]byte("PHP"), []byte("C++"))
	s.Set([]byte("PyTorch"), []byte("Tensorflow"))
	s.Set([]byte("Java"), nil)

	fmt.Printf("The length of this skiplist is %v.\n", s.Len())

	// Get value from a skiplist.
	val, _ := s.Get([]byte("PyTorch"))
	fmt.Printf("The value of the key PyTorch in this skiplist is %v.\n", string(val.([]byte)))

	if _, err := s.Get([]byte("IBM")); err != nil {
		fmt.Printf("The key IBM is not in this skiplist is.\n")
	}

	if s.Contains([]byte("Hello")) == true {
		fmt.Println("The key Hello is exist in this skiplist.")
	}

	// Remove one key from the skiplist
	s.Delete([]byte("Hello"))
	if s.Contains([]byte("Hello")) == false {
		fmt.Println("The key Hello has been removed from this skiplist.")
	}
}
```

# Use the int instead of []byte as key/value type

This skiplist support any key/value type, as long as you implement your Comparator by yourself like below.

```go
package main

import (
	"fmt"
	"github.com/jdxyw/skiplist-go"
)

type IntCmp struct {}

func (IntCmp) Compare(rhs, lhs interface{}) int {
	rhsint := rhs.(int)
	lhsint := lhs.(int)

	switch result := rhsint-lhsint; {
	case result == 0:
		return 0
	case result > 0:
		return 1
	default:
		return -1
	}
}

func (IntCmp) Name() string {
	return "Int64Comparator"
}

func main() {
	// We implement a Int64 Comaparator and pass it.
	// This skiplist would be use int64 as the key/value type.
	s := skiplist.NewSkiplist(10, IntCmp{})

	// Use the Set to insert/update element in this list.
	// The `value` could be nil.
	s.Set(111, 123)
	s.Set(222, 234)
	s.Set(333, 345)
	s.Set(444, 456)
	s.Set(555, 567)


	fmt.Printf("The length of this skiplist is %v.\n", s.Len())

	// Get value from a skiplist.
	val, _ := s.Get(111)
	fmt.Printf("The value of the key 111 in this skiplist is %v.\n", val.(int))

	if _, err := s.Get(666); err != nil {
		fmt.Printf("The key 666 is not in this skiplist is.\n")
	}

	if s.Contains(444) == true {
		fmt.Println("The key 444 is exist in this skiplist.")
	}

	// Remove one key from the skiplist
	s.Delete(444)
	if s.Contains(444) == false {
		fmt.Println("The key 444 has been removed from this skiplist.")
	}
}

```

# Beachmark result

Run benchmark if you are interested.

```bash
    go test -bench=. -benchmem ./benchmark/
```

The result.

```
BenchmarkLevel6-8        1000000             12724 ns/op             176 B/op          6 allocs/op
BenchmarkLevel8-8        1000000              4052 ns/op             194 B/op          6 allocs/op
BenchmarkLevel10-8       1000000              3677 ns/op             212 B/op          6 allocs/op
BenchmarkLevel12-8       1000000              3844 ns/op             229 B/op          6 allocs/op
```