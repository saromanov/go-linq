package linq

import (
	"errors"
	"reflect"
)

var errNoSlice = errors.New("input type not slice")

// Enumerable defines main interface for collections
type Enumerable interface {
	GetList() []interface{}
	Compare(interface{}) int
}

// Linq defines main structure on collection
type Linq struct {
	coll interface{}
	result []*reflect.Value
}

// New creates init of the linq
func New(data interface{}) (*Linq, error) {
	kind := getKind(data)
	ok := isSliceType(kind)
	if !ok {
		return nil, errNoSlice
	}
	return &Linq{
		coll: data,
	}, nil
}

// getKind returns type of the input data
func getKind(data interface{}) reflect.Kind {
	return reflect.TypeOf(data).Kind()
}

// isSliceType returns true if input data is slice
func isSliceType(k reflect.Kind) bool {
	if k == reflect.Slice {
		return true
	}
	return false
}

// isFunction returns true if input data is a function
func isFunction(data interface{}) bool {
	if reflect.TypeOf(data).Kind() == reflect.Func {
		return true
	}
	return false
}
