package listener

type Listener interface {
	Start() error
	GetReadCh() chan string
	Run() error
	Stop()
}
