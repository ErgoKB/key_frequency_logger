package parser

import (
	"fmt"
)

type Event struct {
	Keycode int
	Column  int
	Row     int
	Layer   int
	Pressed bool
}

func (e *Event) String() string {
	var pressedStr string
	if e.Pressed {
		pressedStr = "pressed"
	} else {
		pressedStr = "released"
	}
	return fmt.Sprintf("row: %d, col: %d, layer: %d, keycode: %d %s", e.Row, e.Column, e.Layer, e.Keycode, pressedStr)
}
