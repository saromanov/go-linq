package linq

import (
	"reflect"
)

// Last provides getting of the last element from collection
func (l *Linq) Last(f interface{}) *Linq {
	err := validateFirstFunc(f)
	if err != nil {
		panic(err)
	}

	l.result = last(f, l.result)
	return l
}

func last(f interface{}, data []reflect.Value) []reflect.Value {
	if len(data) == 0 {
		return data
	}

	for i := len(data); i > 0; i-- {
		resp := invoke(f, data[i])
		if len(resp) == 0 {
			continue
		}
		if resp[0].Bool() {
			return []reflect.Value{
				reflect.ValueOf(data[i].Interface()),
			}
		}
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
