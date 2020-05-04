package listener

type Listener interface {
	Start() error
	GetEventCh() chan []string
}
