package rawhid

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/briandowns/spinner"
	log "github.com/sirupsen/logrus"
)

const (
	NewLineChar = '\n'
	BufferSize  = 1024
)

type hidDevice interface {
	open()
	read() ([]byte, error)
	close()
}

type rawHID struct {
	hidDevice hidDevice

	outputCh chan string

	sendBuf    []byte
	sendBufPtr int

	buffer    [][]byte
	bufferIdx int

	pullingDevice       atomic.Value
	stopPullingDeviceCh chan struct{}
	stopConsumeCh       chan struct{}

	runningGroup sync.WaitGroup
}

func NewDefaultRawHID() *rawHID {
	hidDevice := NewHIDWrapper()
	return NewRawHID(hidDevice)
}

func NewRawHID(device hidDevice) *rawHID {
	res := &rawHID{
		hidDevice:           device,
		outputCh:            make(chan string),
		sendBuf:             make([]byte, 128),
		sendBufPtr:          0,
		buffer:              make([][]byte, BufferSize),
		bufferIdx:           0,
		stopPullingDeviceCh: make(chan struct{}),
		stopConsumeCh:       make(chan struct{}),
	}
	res.pullingDevice.Store(false)
	return res
}

func (r *rawHID) Start() {
	log.Info("Waiting for device...")
	s := spinner.New(spinner.CharSets[43], 200*time.Millisecond)
	s.Start()
	r.hidDevice.open()
	s.Stop()
	log.Info("Device connected, starts listening")
}

func (r *rawHID) GetOutputCh() chan string {
	return r.outputCh
}

func (r *rawHID) Stop() {
	if r.pullingDevice.Load().(bool) {
		r.stopPullingDeviceCh <- struct{}{}
	}
	r.hidDevice.close()
	r.runningGroup.Wait()
}

func (r *rawHID) Run() {
	r.runningGroup.Add(1)
	defer r.runningGroup.Done()

	go r.consume()
	defer func() {
		r.stopConsumeCh <- struct{}{}
	}()

	r.pullingDevice.Store(true)
	defer r.pullingDevice.Store(false)

	for {
		select {
		case <-r.stopPullingDeviceCh:
			return
		default:
			read, err := r.hidDevice.read()
			if err != nil {
				log.Info("Device disconnected")
				return
			}
			if len(read) <= 1 {
				continue
			}
			r.bufferRead(read)
		}
	}
}

func (r *rawHID) bufferRead(read []byte) {
	r.runningGroup.Add(1)
	defer r.runningGroup.Done()

	buf := make([]byte, len(read))
	copy(buf[:len(read)], read[:len(read)])
	r.buffer[r.bufferIdx] = buf
	r.bufferIdx = (r.bufferIdx + 1) % BufferSize
}

func (r *rawHID) consume() {
	defer func() { close(r.outputCh) }()
	lastIdx := 0
	for {
		select {
		case <-r.stopConsumeCh:
			return
		default:
			if lastIdx != r.bufferIdx {
				for _, v := range r.buffer[lastIdx] {
					r.send(v)
				}
				lastIdx = (lastIdx + 1) % BufferSize
			}
		}
	}
}

func (r *rawHID) send(val byte) {
	if val == '\n' {
		output := r.sendBuf[:r.sendBufPtr]
		if len(output) != 0 {
			r.outputCh <- string(r.sendBuf[:r.sendBufPtr])
		}
		r.sendBufPtr = 0
		return
	}
	r.sendBuf[r.sendBufPtr] = val
	r.sendBufPtr++
}
