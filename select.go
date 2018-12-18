package linq

import "reflect"

// Select defines a method for mapping of the representation
// of collection on the new form
func (l *Linq) Select(f interface{}) *Linq {
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
		l.result = append(l.result, x)
	}

	return l
}
