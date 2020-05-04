package rawhid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

const (
	TestErrorTimes = 10
)

type mockDevice struct {
	mock.Mock
}

func (m *mockDevice) open() error {
	m.MethodCalled("open")
	return nil
}

func (m *mockDevice) read() (string, error) {
	return "", nil
}

type mockErrorDevice struct {
	mockDevice
	counter int
}

func (m *mockErrorDevice) open() error {
	m.MethodCalled("open")
	m.counter++
	if m.counter == TestErrorTimes {
		return nil
	}
	return fmt.Errorf("mock error")
}

func TestStart(t *testing.T) {
	m := new(mockDevice)
	m.On("open")
	r := rawHID{hidDevice: m}
	r.Start()
	m.AssertCalled(t, "open")
}

func TestSTartWithError(t *testing.T) {
	m := new(mockErrorDevice)
	m.On("open")
	r := rawHID{hidDevice: m}
	r.Start()
	m.AssertNumberOfCalls(t, "open", TestErrorTimes)
}
