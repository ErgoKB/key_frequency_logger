package main

import (
	"github.com/lschyi/key_frequency_logger/listener"
	"github.com/lschyi/key_frequency_logger/parser"
	"github.com/lschyi/key_frequency_logger/writer"
)

type logger struct {
	listener listener.Listener
	parser   *parser.Parser
	writer   *writer.Writer
	stopCh   chan struct{}
}

func newLogger(path string) (*logger, error) {
	w, err := writer.NewWriter(path)
	if err != nil {
		return nil, err
	}
	return &logger{
		listener: listener.NewListener(),
		parser:   parser.NewParser(),
		writer:   w,
		stopCh:   make(chan struct{}),
	}, nil
}

func (l *logger) run() {
	l.listener.Start()
	go l.listener.Run()
	for {
		select {
		case <-l.stopCh:
			return
		case line := <-l.listener.GetOutputCh():
			if event := l.parser.Parse(line); event != nil {
				l.writer.WriteEvent(event)
			}
		}
	}
}

func (l *logger) close() {
	l.stopCh <- struct{}{}
	l.listener.Stop()
	l.writer.Flush()
}