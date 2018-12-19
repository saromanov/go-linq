package linq

import (
	"reflect"
	"sort"
)

// OrderBy provides sorting of elements to increasing
func (l *Linq) OrderBy(f interface{}) *Linq {
	if !isFunction(f) {
		panic("Where: input data not a function")
	}
	numIn := reflect.TypeOf(f).NumIn()
	if numIn != 1 {
		panic("Where: input arguments is not 1")
	}
	outIn := reflect.TypeOf(f).NumOut()
	if outIn != 1 {
		panic("Where: output arguments is not 1")
	}

	oldResult := l.result
	l.result = l.result[:0]
	sort.Sort(oldResult)
	l.result = oldResult
	return l
}
