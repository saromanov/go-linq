package linq

import (
	"errors"
	"reflect"
	"sort"

	"github.com/bradfitz/slice"
)

// sortingInt is helpful method for sorting of
// ints
type sortingInt struct {
	Key   reflect.Value
	Value interface{}
}

// OrderBy provides sorting of elements to increasing
func (l *Linq) OrderBy(f interface{}) *Linq {
	err := validateOrerByInput(f)
	if err != nil {
		panic(err)
	}
	oldResult := l.result
	l.result = l.result[:0]
	l.result = sorting(f, oldResult)
	l.sorted = true
	return l
}

func validateOrerByInput(f interface{}) error {
	if !isFunction(f) {
		return errors.New("OrderBy: input data not a function")
	}
	numIn := reflect.TypeOf(f).NumIn()
	if numIn != 1 {
		return errors.New("OrderBy: input arguments is not 1")
	}
	outIn := reflect.TypeOf(f).NumOut()
	if outIn != 1 {
		return errors.New("OrderBy: output arguments is not 1")
	}
	return nil
}

func sorting(f interface{}, data []reflect.Value) []reflect.Value {
	if len(data) == 0 {
		return data
	}

	first := data[0].Kind()
	switch first {
	case reflect.Int:
		tmp := make([]int, len(data))
		for i, x := range data {
			fResult := invoke(f, x)
			if len(fResult) == 0 {
				continue
			}
			tmp[i] = int(fResult[0].Int())
		}
		sort.Ints(tmp)
		for i := range data {
			data[i] = reflect.ValueOf(tmp[i])
		}
		return data
	case reflect.String:
		tmp := make([]string, len(data))
		helpStr := make([]sortingInt, len(data))
		for i, x := range data {
			fResult := invoke(f, x)
			if len(fResult) == 0 {
				continue
			}

			if fResult[0].Kind() == reflect.String {
				tmp[i] = fResult[0].String()
				continue
			}

			helpStr[i] = sortingInt{
				Key:   x,
				Value: fResult[0].Int(),
			}

		}
		if tmp[0] != "" {
			sort.Strings(tmp)
			for i := range data {
				data[i] = reflect.ValueOf(tmp[i])
			}
			return data
		}
		if len(helpStr) > 0 {
			slice.Sort(helpStr[:], func(i, j int) bool {
				return helpStr[i].Value.(int64) < helpStr[j].Value.(int64)
			})
			for i, x := range helpStr {
				data[i] = reflect.Value(x.Key)
			}
			return data
		}
	case reflect.Float64:
		tmp := make([]float64, len(data))
		for i, x := range data {
			fResult := invoke(f, x)
			if len(fResult) == 0 {
				continue
			}
			if fResult[0].Kind() != reflect.Float64 {
				continue
			}
			tmp[i] = fResult[0].Float()
		}
		sort.Float64s(tmp)
		for i := range data {
			data[i] = reflect.ValueOf(tmp[i])
		}
		return data
	}

	return data
}
