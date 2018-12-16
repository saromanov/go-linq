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
}

// New creates init of the linq
func New(data interface{}) (*Linq, error) {
	ok := isSliceType(data)
	if !ok {
		return nil, errNoSlice
	}
	return &Linq{
		coll: data,
	}, nil
}

// isSliceType returns true if input data is slice
func isSliceType(data interface{}) bool {
	if reflect.TypeOf(data).Kind() == reflect.Slice {
		return true
	}
	return false
}