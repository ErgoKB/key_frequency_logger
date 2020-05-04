package listener

import (
	"github.com/lschyi/key_frequency_logger/listener/rawhid"
)

func NewListener() Listener {
	return rawhid.NewDefaultRawHID()
}
