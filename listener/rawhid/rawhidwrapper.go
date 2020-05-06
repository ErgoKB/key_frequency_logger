package rawhid

import (
	"fmt"
	"unsafe"
)

/*
#cgo linux CFLAGS: -DLINUX
#cgo darwin CFLAGS:  -DDARWIN  -arch x86_64
#cgo darwin LDFLAGS: -framework IOKit -framework CoreFoundation
#cgo windows CFLAGS: -DWINDOWS
#cgo windows LDFLAGS: -lhid -lsetupapi
#include "rawhid_wrapper.h"
*/
import "C"

const (
	ReadBufferSize = 1024
)

type hidWrapper struct {
	buf []byte
}

func NewHIDWrapper() *hidWrapper {
	return &hidWrapper{
		buf: make([]byte, ReadBufferSize),
	}
}

func (h *hidWrapper) open() {
	C.hid_start()
}

func (h *hidWrapper) read() ([]byte, error) {
	num := C.hid_read()

	switch int(num) {
	case -2:
		return nil, fmt.Errorf("buffer overflow")
	case -1:
		return nil, fmt.Errorf("no device connected")
	case 0:
		return nil, nil
	default:
		for i := 0; i < int(num); i++ {
			h.buf[i] = *(*byte)(unsafe.Pointer(&(C.returnBuf[i])))
		}
		return h.buf[:int(num)], nil
	}
}

func (h *hidWrapper) close() {
	C.hid_close()
}
