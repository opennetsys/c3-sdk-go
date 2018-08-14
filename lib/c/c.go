package main

/*
#ifndef C3_WRAPPER
#include <c3_wrapper.h>
#endif
*/

import (
	"C"
	"sync"
	"unsafe"

	sdk "github.com/c3systems/c3-sdk-go"
)

var (
	c3   *sdk.C3
	once sync.Once
)

// Build must be called first and is used to construct the library
//export Build
func Build() {
	once.Do(func() {
		c3 = sdk.NewC3()
	})
}

// RegisterMethod is used to mount a method or function
//export RegisterMethod
func RegisterMethod(methodName *C.char, types **C.char, typesLength C.int, ifn unsafe.Pointer) {
	mName := C.GoString(methodName)
	length := int(typesLength)

	// note: convert the C array to a Go Array so we can index it
	a := (*[1 << 30]*C.char)(unsafe.Pointer(types))[:length:length]

	var tmpTypes []string
	for idx := range a {
		tmpTypes = append(tmpTypes, C.GoString(a[idx]))
	}

	// note: convert to interface
	fn := (interface{})(unsafe.Pointer(ifn))

	if err := c3.RegisterMethod(mName, tmpTypes, fn); err != nil {
		// TODO: how best to handle this? Pass back to C?
		panic(err)
	}
}

// Serve starts the server
//export Serve
func Serve() {
	c3.Serve()
}

// TODO: implement State()

// Set is used to set a value to the state
//export Set
func Set(key, value *C.char) {
	k := C.GoString(key)
	v := C.GoString(value)

	if err := c3.State().Set([]byte(k), []byte(v)); err != nil {
		// TODO: handle this better? See comment, above of same type
		panic(err)
	}
}

// Get is used to retrieve a value from the state
//export Get
func Get(key *C.char) (*C.char, C.int) {
	k := C.GoString(key)

	s := C.CString("")
	b := C.int(0)
	if r, ok := c3.State().Get([]byte(k)); ok {
		s = C.CString(string(r))
		b = C.int(1)
	}

	return s, b
}

func main() {}
