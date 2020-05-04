package rawhid

import (
	"strings"
)

const (
	NewLineChar = '\n'
	ChanBuffer  = 100
)

type hidDevice interface {
	open() error
	read() (string, error)
	close() error
}

type rawHID struct {
	hidDevice  hidDevice
	incomplete string
	readCh     chan string
}

func NewRawHID(device hidDevice) *rawHID {
	return &rawHID{
		hidDevice: device,
		readCh:    make(chan string, ChanBuffer),
	}
}

func (r *rawHID) Start() {
	for {
		err := r.hidDevice.open()
		if err == nil {
			break
		}
	}
}

func (r *rawHID) GetReadCh() chan string {
	return r.readCh
}

func (r *rawHID) Close() {
	r.hidDevice.close()
}

func (r *rawHID) Run() error {
	for {
		read, err := r.hidDevice.read()
		if err != nil {
			return nil
		}
		r.handleRead(strings.Split(read, string(NewLineChar)))
	}
}

func (r *rawHID) handleRead(lines []string) {
	if len(lines) == 0 {
		return
	}

	var incompleteLine string
	if lines[len(lines)-1] != "" {
		incompleteLine = lines[len(lines)-1]
		lines = lines[:len(lines)-1]
	}
	if len(lines) == 0 {
		return
	}
	if r.incomplete != "" {
		lines[0] = r.incomplete + lines[0]
	}
	r.sendNonEmptyRead(lines)
	r.incomplete = incompleteLine
}

func (r *rawHID) sendNonEmptyRead(lines []string) {
	for _, line := range lines {
		if line == "" {
			continue
		}
		r.readCh <- line
	}
}
