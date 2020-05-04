package rawhid

import (
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	NewLineChar  = '\n'
	ChanBuffer   = 100
	OpenInterval = 200 * time.Millisecond
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
	stopCh     chan struct{}
}

func NewRawHID(device hidDevice) *rawHID {
	return &rawHID{
		hidDevice: device,
		readCh:    make(chan string, ChanBuffer),
		stopCh:    make(chan struct{}),
	}
}

func (r *rawHID) Start() {
	firstOpen := true
	for {
		select {
		case <-r.stopCh:
			return
		default:
			err := r.hidDevice.open()
			if err == nil {
				log.Info("Found Device")
				return
			}
			if firstOpen {
				firstOpen = false
				log.Info("Waiting for device...")
			}
			<-time.After(time.Millisecond)
		}
	}
}

func (r *rawHID) GetReadCh() chan string {
	return r.readCh
}

func (r *rawHID) Close() {
	r.stopCh <- struct{}{}
	r.hidDevice.close()
}

func (r *rawHID) Run() error {
	for {
		select {
		case <-r.stopCh:
			return nil
		default:
			read, err := r.hidDevice.read()
			if err != nil {
				return nil
			}
			r.handleRead(strings.Split(r.incomplete+read, string(NewLineChar)))
		}
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
