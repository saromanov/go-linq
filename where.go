package linq

import (
	"reflect"
)

// Where provides filtering on collection
func (l *Linq) Where(f func(d interface{}) bool) *Linq {
	s := reflect.ValueOf(l.coll)
	for i := 0; i < s.Len(); i++ {
		f(s.Index(i).Interface())
	}
	return l
}
