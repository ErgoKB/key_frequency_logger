package main

import (
	"fmt"

	"github.com/lschyi/key_frequency_logger/listener"
	"github.com/lschyi/key_frequency_logger/parser"
	"github.com/lschyi/key_frequency_logger/writer"
)

const (
	MaxDisplayKeystroke = 32
)

type logger struct {
	listener listener.Listener
	parser   *parser.Parser
	writer   *writer.Writer
	doneCh   chan struct{}
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
		doneCh:   make(chan struct{}),
	}, nil
}

func (l *logger) run() {
	l.listener.Start()
	go l.listener.Run()
	counter := 0
	for line := range l.listener.GetOutputCh() {
		if event := l.parser.Parse(line); event != nil {
			l.writer.WriteEvent(event)
			if !event.Pressed {
				continue
			}
			fmt.Print(".")
			counter = (counter + 1) % MaxDisplayKeystroke
			if counter == 0 {
				fmt.Printf("%c[2K\r", 27)
			}
		}
	}
	l.doneCh <- struct{}{}
}

func (l *logger) close() {
	l.writer.Flush()
	l.listener.Stop()
}
