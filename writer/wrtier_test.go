package writer

import (
	"bytes"
	"testing"

	"github.com/lschyi/key_frequency_logger/parser"
	"github.com/stretchr/testify/assert"
)

func TestWriteEvent(t *testing.T) {
	event := parser.Event{
		Keycode: 1,
		Column:  2,
		Row:     3,
		Layer:   4,
		Pressed: false,
	}
	expected := "3,2,4,1,false\n"
	w := bytes.NewBuffer(make([]byte, 0))
	testWriter := newWriter(w)
	err := testWriter.WriteEvent(&event)
	testWriter.Flush()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(w.Bytes()))
}

func TestWriteHeader(t *testing.T) {
	expected := "row,col,layer,keycode,pressed\n"
	w := bytes.NewBuffer(make([]byte, 0))
	testWriter := newWriter(w)
	err := testWriter.writeHeader()
	testWriter.Flush()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(w.Bytes()))
}
