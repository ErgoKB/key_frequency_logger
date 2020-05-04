package rawhid

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	TestErrorTimes = 10
)

type mockDevice struct {
	mock.Mock
	mockRead []string
}

func (m *mockDevice) open() error {
	m.MethodCalled("open")
	return nil
}

func (m *mockDevice) read() (string, error) {
	if len(m.mockRead) == 0 {
		for {
		}
	}
	res := m.mockRead[0]
	m.mockRead = m.mockRead[1:]
	return res, nil
}

func (m *mockDevice) close() error {
	m.MethodCalled("close")
	return nil
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

func TestReadOneLine(t *testing.T) {
	m := new(mockDevice)
	mockRead := []string{
		"first line\n",
	}
	m.mockRead = mockRead
	r := NewRawHID(m)
	go r.Run()
	outputCh := r.GetReadCh()
	line := <-outputCh
	assert.Equal(t, strings.TrimSpace(mockRead[0]), line)
}

func TestReadTwoLines(t *testing.T) {
	m := new(mockDevice)
	mockRead := []string{
		"first line\n",
		"second line\n",
	}
	m.mockRead = mockRead
	r := NewRawHID(m)
	go r.Run()
	outputCh := r.GetReadCh()
	for _, line := range mockRead {
		read := <-outputCh
		assert.Equal(t, strings.TrimSpace(line), read)
	}
}

func TestIncompleteLine(t *testing.T) {
	m := new(mockDevice)
	mockRead := []string{
		"first line\nincomplete tail",
	}
	m.mockRead = mockRead
	r := NewRawHID(m)
	go r.Run()
	outputCh := r.GetReadCh()
	read := <-outputCh
	assert.Equal(t, "first line", read)
}

func TestTwoIncompleteLines(t *testing.T) {
	m := new(mockDevice)
	mockRead := []string{
		"first line ",
		"incomplete ",
		"complete tail\n",
	}
	m.mockRead = mockRead
	r := NewRawHID(m)
	go r.Run()
	outputCh := r.GetReadCh()
	read := <-outputCh
	assert.Equal(t, "first line incomplete complete tail", read)
}

func TestComposeIncompleteLine(t *testing.T) {
	m := new(mockDevice)
	mockRead := []string{
		"first line\nincomplete tail",
		" this one complete it\n",
	}
	m.mockRead = mockRead
	r := NewRawHID(m)
	go r.Run()
	outputCh := r.GetReadCh()
	<-outputCh
	read := <-outputCh
	assert.Equal(t, "incomplete tail this one complete it", read)
	assert.Equal(t, "", r.incomplete)
}

func TestClose(t *testing.T) {
	m := new(mockDevice)
	m.On("close")
	r := NewRawHID(m)
	go r.Run()
	r.Close()
	m.AssertNumberOfCalls(t, "close", 1)
}
