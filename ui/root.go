package ui

// Root é a ponte entre a janela e a árvore de widgets: não participa do
// protocolo de layout (não é medido nem posicionado por ninguém), apenas
// ancora um único Child na origem e delega Measure/Layout/Draw a ele.
type Root struct {
	Child Widget

	Background Color
}

func (r *Root) Draw(c *Canvas) {
	c.Clear(r.Background)

	if r.Child == nil {
		return
	}

	margin := r.Child.GetMargin()

	r.Child.Draw(c, Point{margin.L, margin.T})
}

func (r *Root) Layout(width, height int) {
	if r.Child == nil {
		return
	}

	r.Child.Layout(Constraint{
		MaxH: width,
		MaxW: height,
	})
}
