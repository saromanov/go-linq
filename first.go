package linq

import (
	"errors"
	"reflect"
)

// First provides getting of the first element from collection
func (l *Linq) First(f interface{}) *Linq {
	err := validateFirstFunc(f)
	if err != nil {
		panic(err)
	}

	l.result = first(f, l.result)
	return l
}

func first(f interface{}, data []reflect.Value) []reflect.Value {
	if len(data) == 0 {
		return data
	}

	for _, x := range data {
		resp := invoke(f, x)
		if len(resp) == 0 {
			continue
		}
		if resp[0].Bool() {
			return []reflect.Value{
				reflect.ValueOf(x.Interface()),
			}
		}
	}

	return []reflect.Value{}
}

func validateFirstFunc(f interface{}) error {
	if !isFunction(f) {
		return errors.New("first: input data not a function")
	}
	numIn := reflect.TypeOf(f).NumIn()
	if numIn != 1 {
		return errors.New("first: input arguments is not 1")
	}
	outIn := reflect.TypeOf(f).NumOut()
	if outIn != 1 {
		return errors.New("first: output arguments is not 1")
	}
	return nil
}
