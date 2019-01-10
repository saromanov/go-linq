package linq

import (
	"reflect"
)

// Where provides filtering on collection
func (l *Linq) Where(f interface{}) *Linq {
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
	for _, x := range oldResult {
		fResult := invoke(f, x)
		if len(fResult) == 0 {
			continue
		}
		if fResult[0].Bool() {
			l.result = append(l.result, reflect.ValueOf(x.Interface()))
		}
	}
	return l
}

// Invoke provides call of the method
func invoke(fn interface{}, val reflect.Value) []reflect.Value {
	return reflect.ValueOf(fn).Call([]reflect.Value{val})
}

// Invoke provides call of the method with two arguments
func invoke2(fn interface{}, val, val2 reflect.Value) []reflect.Value {
	return reflect.ValueOf(fn).Call([]reflect.Value{val, val2})
}
