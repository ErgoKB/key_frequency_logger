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
		Keycode: 13,
		Column:  5,
		Row:     4,
		Layer:   0,
		Pressed: false,
	}
	res := p.Parse("ErgoKB:5,4,0,0,13")
	assert.Equal(t, &event, res)
}

func TestParseConvertFail(t *testing.T) {
	p := NewParser()
	res := p.Parse("ErgoKB:5,4,0,true,13")
	assert.Nil(t, res)
}

func TestParseInvalidLine(t *testing.T) {
	p := NewParser()
	res := p.Parse("ErgoKB:5,4,0,0,13Ergokb:2,2")
	assert.Nil(t, res)
}
