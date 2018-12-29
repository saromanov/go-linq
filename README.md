Implementation of [LINQ](https://msdn.microsoft.com/en-us/library/bb308959.aspx) queries on Golang

[![Go Report Card](https://goreportcard.com/badge/github.com/saromanov/go-linq)](https://goreportcard.com/report/github.com/saromanov/go-linq)

## Goal

Simplify operations over slices and study of working of reflection

## Queries

Supported queries:
* Where

`Where`

```go
package main

import (
	"fmt"

	"github.com/saromanov/go-linq"
)

func main() {
	l, err := linq.New([]int{1, 4, 5, 8, 9})
	if err != nil {
		panic(err)
	}
	fmt.Println(l.Where(func(x int) bool {
		if x > 4 {
			return true
		}
		return false
	}).Result().([]int)
}
```
``` [5 8 9]```

In this example: First, create a Linq object, then apply of the Where method with filter. Filter function should be ```go func(x int) bool{}```. At the last step, showing of response with `Result` method

`Select`

Select method to modify the elements in an array.

```go
package main

import (
	"fmt"

	"github.com/saromanov/go-linq"
)

type Data struct {
	Key   int
	Value int
}

func main() {
	l, err := linq.New([]Data{Data{
		Key:   1,
		Value: 2,
	},
		Data{
			Key:   3,
			Value: 5,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(l.Select(func(x Data) int {
		return x.Key
	}).Result().([]int))

}
```

