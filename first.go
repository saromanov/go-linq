package linq

// First provides getting of the first element from collection
func (l *Linq) First(f interface{}) *Linq {
	err := validateFirstFunc(f)
	if err != nil {
		panic(err)
	}
	oldResult := l.result
	return l
}

func validateFirstFunc(f interface{}) error {
	if !isFunction(f) {
		return errors.New("First: input data not a function")
	}
	numIn := reflect.TypeOf(f).NumIn()
	if numIn != 1 {
		return errors.New("First: input arguments is not 1")
	}
	outIn := reflect.TypeOf(f).NumOut()
	if outIn != 1 {
		return errors.New("First: output arguments is not 1")
	}
	return nil
}