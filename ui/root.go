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

	r.Child.Draw(c)
}

func (r *Root) Layout() {
	if r.Child == nil {
		return
	}

	size := r.Child.Measure()
	margin := r.Child.GetMargin()

	r.Child.Layout(Rect{
		X: margin.L,
		Y: margin.T,
		W: size.Width,
		H: size.Height,
	})
}
