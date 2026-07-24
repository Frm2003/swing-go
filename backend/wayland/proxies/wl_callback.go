package proxies

import "swing-go/backend/wayland/protocol"

type WlCallback struct {
	objectId uint32
	ready    chan struct{}
}

func NewWlCallback(newId uint32) *WlCallback {
	return &WlCallback{
		objectId: newId,
		ready:    make(chan struct{}),
	}
}

func (wl *WlCallback) Handle(message *protocol.Message) {
	close(wl.ready)
}

func (wl *WlCallback) GetId() uint32 {
	return wl.objectId
}

func (w *WlCallback) Wait() {
	<-w.ready
}
