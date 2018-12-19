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
	l.result = sorting(oldResult)
	return l
}

func sorting(data []reflect.Value) []reflect.Value {
	if len(data) == 0 {
		return data
	}

	first := data[0].Kind()
	switch first {
	case reflect.Int:
		tmp := make([]int, len(data))
		for i, x := range data {
			tmp[i] = int(x.Int())
		}
		sort.Ints(tmp)
		for i := range data {
			data[i] = reflect.ValueOf(tmp[i])
		}
		return data
	}

	return data
}
