package rawhid

type hidDevice interface {
	open() error
	read() (string, error)
}

type rawHID struct {
	hidDevice hidDevice
}

func NewRawHID(device hidDevice) *rawHID {
	return &rawHID{hidDevice: device}
}

func (r *rawHID) Start() {
	for {
		err := r.hidDevice.open()
		if err == nil {
			break
		}
	}
}

func (r *rawHID) Read() (string, error) {
	return "", nil
}
