// +build dev

package rawhid

import (
	"fmt"
	"testing"
)

func TestRawRead(t *testing.T) {
	device := NewHIDWrapper()
	fmt.Println("waiting for device")
	device.open()
	fmt.Println("device successfully opened, starts test")
	for i := 0; i < 2000; i++ {
		output, _ := device.read()
		if len(output) > 0 {
			fmt.Print(string(output))
		}
	}
	device.close()
}

func TestRawHID(t *testing.T) {
	hid := NewDefaultRawHID()
	hid.Start()
	ch := hid.GetReadCh()
	go hid.Run()
	for i := 0; i < 1000; i++ {
		output := <-ch
		fmt.Println(output)
	}
}
