package parser

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	TargetPrefix = "KL: "
	SplitToken   = ", "
)

type Parser struct{}

func NewParser() *Parser {
	return new(Parser)
}

func (p *Parser) Parse(line string) *Event {
	if !strings.HasPrefix(line, TargetPrefix) {
		return nil
	}
	trimed := strings.TrimPrefix(line, TargetPrefix)
	tokens := strings.Split(trimed, SplitToken)
	event, err := parseEvent(tokens)
	if err != nil {
		return nil
	}
	return event
}

func parseEvent(tokens []string) (*Event, error) {
	res := new(Event)

	for _, token := range tokens {
		fields := strings.Split(token, ": ")
		if len(fields) != 2 {
			return nil, fmt.Errorf("invalid fields")
		}
		val, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, fmt.Errorf("can not parse field")
		}
		switch fields[0] {
		case "kc":
			res.Keycode = val
		case "col":
			res.Column = val
		case "row":
			res.Row = val
		case "layer":
			res.Layer = val
		case "pressed":
			if val == 1 {
				res.Pressed = true
			} else if val == 0 {
				res.Pressed = false
			}
		default:
			log.Warn("unknown field, skipped")
		}
	}

	return res, nil
}
