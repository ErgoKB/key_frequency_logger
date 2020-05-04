package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNonTargetLine(t *testing.T) {
	p := NewParser()
	res := p.Parse("asdfline")
	assert.Nil(t, res)
}

func TestParseEventLine(t *testing.T) {
	p := NewParser()
	event := Event{
		Keycode: 9,
		Column:  5,
		Row:     1,
		Layer:   0,
		Pressed: true,
	}
	res := p.Parse("KL: kc: 9, col: 5, row: 1, pressed: 1, layer: 0")
	assert.Equal(t, &event, res)
}
