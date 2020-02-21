package safe

// Must will return the first argument as the result, if the error is nil, otherwise will call panic.
func Must(any interface{}, err error) interface{} {
	if !IsNil(err) {
		panic(err)
	}
	return any
}
