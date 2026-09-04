package ui

type Element interface {
	Draw(*Canvas)
}

type RenderTree struct {
	root Element
}

func (r *RenderTree) Render(c *Canvas) error {
	return nil
}
