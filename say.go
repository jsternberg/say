package main

// #include "say.h"
// void free(void *);
import "C"

import (
	"os"
	"unsafe"
)

func say(s string) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.say(cstr)
}

func main() {
	for _, arg := range os.Args[1:] {
		say(arg)
	}
}
