package linq

import "reflect"

// ThenBy provides sorting by the key
// It should have "sorted": true after call of OrderBy
func (l *Linq) ThenBy(f interface{}) *Linq {
	if !l.sorted {
		return l
	}
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
	l.result = sorting(f, oldResult)
	l.sorted = true
	return l
}
