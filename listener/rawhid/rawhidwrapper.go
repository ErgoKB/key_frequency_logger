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
#include "rawhid.h"
*/
import "C"

const (
	Vid                = 0
	Pid                = 0
	UsagePage          = 0xFF31
	Usage              = 0x0074
	BufSize            = 64
	TimeoutMillisecond = 200
	NewLineChar        = '\n'
)

type rawHIDWrapper struct {
	hid *C.rawhid_t
	buf []byte
}

func NewHIDWrapper() *rawHIDWrapper {
	return &rawHIDWrapper{
		buf: make([]byte, BufSize),
	}
}

func (r *rawHIDWrapper) open() error {
	hid := C.rawhid_open_only1(C.int(Vid), C.int(Pid), C.int(UsagePage), C.int(Usage))
	if hid == nil {
		return fmt.Errorf("no device found")
	}
	r.hid = (*C.rawhid_t)(unsafe.Pointer(hid))
	return nil
}

func (r *rawHIDWrapper) read() (string, error) {
	numInC := C.rawhid_read(unsafe.Pointer(r.hid), unsafe.Pointer(&r.buf[0]), C.int(BufSize), C.int(TimeoutMillisecond))
	num := int(numInC)
	if num < 0 {
		return "", fmt.Errorf("device disconnected")
	}
	return string(r.buf[:num]), nil
}
