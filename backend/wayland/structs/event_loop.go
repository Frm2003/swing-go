package structs

type EventLoop struct {
	commands chan iCommand
}

func NewEventLoop() *EventLoop {
	loop := &EventLoop{
		commands: make(chan iCommand),
	}

	go loop.Run()

	return loop
}

func call[T any](e *EventLoop, f func() T) T {
	done := make(chan T)
	e.commands <- newCommand(f, done)
	return <-done
}

func (e *EventLoop) Run() {
	for cmd := range e.commands {
		cmd.execute()
	}
}
