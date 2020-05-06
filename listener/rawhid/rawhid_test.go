package rawhid

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	TestErrorTimes = 10
)

type mockDevice struct {
	mock.Mock
	mockRead [][]byte
}

func (m *mockDevice) open() {
	m.MethodCalled("open")
}

func (m *mockDevice) read() ([]byte, error) {
	if len(m.mockRead) == 0 {
		for {
		}
	}
	res := m.mockRead[0]
	m.mockRead = m.mockRead[1:]
	return res, nil
}

func (m *mockDevice) close() {
	m.MethodCalled("close")
}

func TestStart(t *testing.T) {
	m := new(mockDevice)
	m.On("open")
	r := rawHID{hidDevice: m}
	r.Start()
	m.AssertCalled(t, "open")
}

func TestReadOneLine(t *testing.T) {
	m := new(mockDevice)
	mockRead := [][]byte{
		[]byte("first line\n"),
	}
	m.mockRead = mockRead
	r := NewRawHID(m)
	go r.Run()
	outputCh := r.GetOutputCh()
	line := <-outputCh
	assert.Equal(t, string(bytes.TrimSpace(mockRead[0])), line)
}

func TestReadTwoLines(t *testing.T) {
	m := new(mockDevice)
	mockRead := [][]byte{
		[]byte("first line\n"),
		[]byte("second line\n"),
	}
	m.mockRead = mockRead
	r := NewRawHID(m)
	go r.Run()
	outputCh := r.GetOutputCh()
	for _, line := range mockRead {
		read := <-outputCh
		assert.Equal(t, string(bytes.TrimSpace(line)), read)
	}
}

func TestReadTwoNewLines(t *testing.T) {
	m := new(mockDevice)
	mockRead := [][]byte{
		[]byte("first\nsecond\n"),
	}
	m.mockRead = mockRead
	r := NewRawHID(m)
	go r.Run()
	outputCh := r.GetOutputCh()
	read := <-outputCh
	assert.Equal(t, "first", read)
	read = <-outputCh
	assert.Equal(t, "second", read)
}

func TestNotSendEmptyLine(t *testing.T) {
	m := new(mockDevice)
	mockRead := [][]byte{
		[]byte("first\n\nsecond\n"),
	}
	m.mockRead = mockRead
	r := NewRawHID(m)
	go r.Run()
	outputCh := r.GetOutputCh()
	read := <-outputCh
	assert.Equal(t, "first", read)
	read = <-outputCh
	assert.Equal(t, "second", read)
}

func TestClose(t *testing.T) {
	m := new(mockDevice)
	m.On("close")
	r := NewRawHID(m)
	go r.Run()
	r.Stop()
	m.AssertNumberOfCalls(t, "close", 1)
}
