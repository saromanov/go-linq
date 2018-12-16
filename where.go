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
	s := reflect.ValueOf(l.coll)
	for i := 0; i < s.Len(); i++ {
		fmt.Println(s.Index(i).Interface())
	}
	return l
}
