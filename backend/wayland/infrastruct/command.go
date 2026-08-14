package infrastruct

type iCommand interface {
	execute()
}

type run[T any] func() T

type command[T any] struct {
	run  func() T
	done chan T
}

func newCommand[T any](run run[T], done chan T) *command[T] {
	return &command[T]{
		run:  run,
		done: done,
	}
}

func (c *command[T]) execute() {
	c.done <- c.run()
}
