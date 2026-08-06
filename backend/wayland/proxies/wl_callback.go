package proxies

import "swing-go/backend/wayland/protocol"

type Wlcallback struct {
	objectId uint32
	ready    chan struct{}
}

func NewWlcallback(newId uint32) *Wlcallback {
	return &Wlcallback{
		objectId: newId,
		ready:    make(chan struct{}),
	}
}

func (wl *Wlcallback) Handle(message *protocol.Message) {
	wl.ready <- struct{}{}
}

func (wl *Wlcallback) GetId() uint32 {
	return wl.objectId
}

func (w *Wlcallback) Wait() {
	<-w.ready
}
