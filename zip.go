package linq

import "reflect"

// Zip provides applying a specified function to the corresponding elements of
// two sequences, producing a sequence of the results.
func (l *Linq) Zip(coll interface{}, f interface{}) *Linq {
	if !isFunction(f) {
		panic("Zip: input data not a function")
	}
	size := len(l.result)
	data := makeData(coll)
	sizeInp := len(data)
	length := size
	if size > sizeInp {
		length = sizeInp
	}
	tmp := l.result
	l.result = l.result[:0]
	for i := 0; i < length; i++ {
		fResult := invoke2(f, data[i], tmp[i])
		l.result = append(l.result, reflect.ValueOf(fResult[0].Interface()))
	}
	return l
}
