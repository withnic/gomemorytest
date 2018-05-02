package main

func main() {
	returnValue()
	returnPointer()
	returnValueUsePointer()
	noReturnNew()
	returnNew()

	sli := []int{1, 2, 3}
	sliceGet(sli)

	i := 1
	pointerGet(&i)
	noReturnPointerGet(&i)
}

// slizeGet has alarting leaking param message
func sliceGet(sl []int) []int {
	return sl[:]
}

// pointerGet has alarting leaking param message
func pointerGet(pt *int) *int {
	*pt = *pt + 1
	return pt
}

// noReturnPointerGet has no alarting
func noReturnPointerGet(pt *int) {
	*pt = *pt + 1
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
