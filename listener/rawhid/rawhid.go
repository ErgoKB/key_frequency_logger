package rawhid

import (
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

	stopPullingDeviceCh chan struct{}
	stopConsumeCh       chan struct{}
}

func NewDefaultRawHID() *rawHID {
	hidDevice := NewHIDWrapper()
	return NewRawHID(hidDevice)
}

func NewRawHID(device hidDevice) *rawHID {
	return &rawHID{
		hidDevice:           device,
		outputCh:            make(chan string),
		sendBuf:             make([]byte, 128),
		sendBufPtr:          0,
		buffer:              make([][]byte, BufferSize),
		bufferIdx:           0,
		stopPullingDeviceCh: make(chan struct{}),
		stopConsumeCh:       make(chan struct{}),
	}
}

func (r *rawHID) Start() {
	log.Info("Waiting for device...")
	r.hidDevice.open()
	log.Info("Device connected, starts listening")
}

func (r *rawHID) GetOutputCh() chan string {
	return r.outputCh
}

func (r *rawHID) Stop() {
	go func() {
		r.stopPullingDeviceCh <- struct{}{}
	}()
	r.hidDevice.close()
}

func (r *rawHID) Run() {
	go r.consume()
	defer func() {
		r.stopConsumeCh <- struct{}{}
	}()

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
