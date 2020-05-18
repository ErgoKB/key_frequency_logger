package listener

import (
	"github.com/ErgoKB/key_frequency_logger/listener/rawhid"
)

func NewListener() Listener {
	return rawhid.NewDefaultRawHID()
}
