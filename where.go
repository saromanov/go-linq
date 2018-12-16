package linq

import (
	"fmt"
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
	s := reflect.ValueOf(l.coll)
	for i := 0; i < s.Len(); i++ {
		fmt.Println(s.Index(i).Interface())
	}
	return l
}

// Invoke provides call of the method
func invoke(fn interface{}, val reflect.Value) []reflect.Value {
	return reflect.ValueOf(fn).Call([]reflect.Value{val})
}
