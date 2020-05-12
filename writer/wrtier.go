package writer

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/lschyi/key_frequency_logger/parser"

	log "github.com/sirupsen/logrus"
)

type Writer struct {
	csv.Writer
}

func NewWriter(path string) (*Writer, error) {
	isOutputFileExists := isFileExists(path)
	if isOutputFileExists {
		log.Warnf("%s exists, append results to it", path)
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	res := newWriter(f)
	if !isOutputFileExists {
		if err := res.writeHeader(); err != nil {
			return nil, err
		}
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

func isFileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
