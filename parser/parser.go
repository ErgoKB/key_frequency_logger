package parser

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	TargetPrefix = "ErgoKB:"
	SplitToken   = ","
)

type Parser struct{}

func NewParser() *Parser {
	return new(Parser)
}

func (p *Parser) Parse(line string) *Event {
	if !strings.HasPrefix(line, TargetPrefix) {
		return nil
	}
	event, err := parseEvent(line)
	if err != nil {
		log.WithError(err).WithField("input line", line).Warn("Encounter error while parsing, skipped")
		return nil
	}
	return event
}

func parseEvent(line string) (*Event, error) {
	res := new(Event)

	trimed := strings.TrimPrefix(line, TargetPrefix)
	splitted := strings.Split(trimed, SplitToken)
	if len(splitted) != 5 {
		return nil, fmt.Errorf("Malformed line, skipped")
	}
	for idx, valStr := range splitted {
		val, err := strconv.Atoi(valStr)
		if err != nil {
			return nil, err
		}
		switch idx {
		case 0:
			res.Column = val
		case 1:
			res.Row = val
		case 2:
			res.Layer = val
		case 3:
			if val == 1 {
				res.Pressed = true
			} else {
				res.Pressed = false
			}
		case 4:
			res.Keycode = val
		}
	}

	return res, nil
}
