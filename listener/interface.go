package listener

type Listener interface {
	Start()
	GetReadCh() chan string
	Run() error
	Stop()
}
