package main

func main() {
	returnValue()
	returnPointer()
	returnValueUsePointer()
	noReturnNew()
	returnNew()
}

// noReturnNew use stack
func noReturnNew() {
	nr := new(int)
	if nr == nil {
	}
}

// returnNew use heap
func returnNew() *int {
	nr := new(int)
	return nr
}

// returnValue use stack (i)
func returnValue() int {
	rv := 1
	return rv
}

// returnValueUserPointer use stack
func returnValueUsePointer() int {
	i := 1
	rvup := &i
	return *rvup
}

// returnPointer use heap (*i)
func returnPointer() *int {
	rp := 1
	return &rp
}
