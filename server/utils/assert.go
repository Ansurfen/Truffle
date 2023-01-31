package utils

func AssertWithPanic(e error) {
	if e != nil {
		panic(e)
	}
}

func AssertWithHandle(e error, h func(e error)) {
	if e != nil {
		h(e)
	}
}

func SafeAssert(e error, f func(e error)) error {
	if e != nil {
		f(e)
	}
	return e
}
