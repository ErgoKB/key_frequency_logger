// +build dev

package rawhid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRawHID(t *testing.T) {
	device := NewHIDWrapper()
	assert.NotNil(t, device)
}

func TestRead(t *testing.T) {
	device := NewHIDWrapper()
	fmt.Println("waiting for device")
	for err := device.open(); err != nil; err = device.open() {
	}
	fmt.Println("device successfully opened, starts test")
	for i := 0; i < 100; i++ {
		output, err := device.read()
		assert.NoError(t, err)
		fmt.Println(output)
	}
}
