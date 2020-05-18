package main

import (
	"fmt"

	"github.com/ErgoKB/key_frequency_logger/listener"
	"github.com/ErgoKB/key_frequency_logger/parser"
	"github.com/ErgoKB/key_frequency_logger/writer"

	log "github.com/sirupsen/logrus"
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
	defer func() { l.doneCh <- struct{}{} }()
	for line := range l.listener.GetOutputCh() {
		if err := l.handleLine(line); err != nil {
			log.Errorf("get error: %s, can not handle this line: %s", err, line)
			return
		}
	}
}

func (l *logger) handleLine(line string) error {
	event := l.parser.Parse(line)
	if event == nil {
		return nil
	}

	if err := l.writer.WriteEvent(event); err != nil {
		return err
	}

	fmt.Printf("%c[2K\r", 27)
	fmt.Printf("%+v", event)

	return nil
}

func (l *logger) close() {
	l.writer.Flush()
	l.listener.Stop()
}
