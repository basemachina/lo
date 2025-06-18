# lo

[![Go Reference](https://pkg.go.dev/badge/github.com/basemachina/lo.svg)](https://pkg.go.dev/github.com/basemachina/lo)

A Go utility library providing functional programming helpers for slices, maps, and other data structures. This library is based on and inspired by [samber/lo](https://github.com/samber/lo), offering type-safe operations using Go generics.

## Installation

```bash
go get github.com/basemachina/lo
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/basemachina/lo"
)

func main() {
    // Map operation
    numbers := []int{1, 2, 3, 4, 5}
    doubled := lo.Map(numbers, func(x int) int {
        return x * 2
    })
    fmt.Println(doubled) // [2, 4, 6, 8, 10]

    // Filter operation
    evens := lo.Filter(numbers, func(x int) bool {
        return x%2 == 0
    })
    fmt.Println(evens) // [2, 4]

    // Check for duplicates
    hasDuplicates := lo.HasDuplicates([]int{1, 2, 2, 3})
    fmt.Println(hasDuplicates) // true

    // Invert map
    original := map[string]int{"a": 1, "b": 2}
    inverted := lo.Invert(original)
    fmt.Println(inverted) // map[1:a 2:b]
}
```

## Requirements

- Go 1.24.0 or later
