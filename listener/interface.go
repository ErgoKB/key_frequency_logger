package listener

type Listener interface {
	Start()
	GetOutputCh() chan string
	Run()
	Stop()
}
