package writer

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/lschyi/key_frequency_logger/parser"
)

type Writer struct {
	csv.Writer
}

func NewWriter(path string) (*Writer, error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	res := newWriter(f)
	if err := res.writeHeader(); err != nil {
		return nil, err
	}
	return res, nil
}

func newWriter(writer io.Writer) *Writer {
	return &Writer{
		Writer: *csv.NewWriter(writer),
	}
}

func (w *Writer) WriteEvent(e *parser.Event) error {
	pressedVal := "true"
	if !e.Pressed {
		pressedVal = "false"
	}
	val := []string{
		strconv.Itoa(e.Row),
		strconv.Itoa(e.Column),
		strconv.Itoa(e.Layer),
		strconv.Itoa(e.Keycode),
		pressedVal,
	}
	return w.Write(val)
}

func (w *Writer) writeHeader() error {
	header := []string{
		"row",
		"col",
		"layer",
		"keycode",
		"pressed",
	}
	return w.Write(header)
}
