package ui

type Root struct {
	Backgroud Color
	Children  []Widget
}

func (r *Root) Draw(c *Canvas) {
	c.Clear(r.Backgroud)

	for _, child := range r.Children {
		child.Draw(c)
	}
}

func (r *Root) Layout() {
	y := 0

	for _, child := range r.Children {
		size := child.Measure()
		margin := child.GetMargin()

		child.Layout(Rect{
			X: margin.L,
			Y: y + margin.T,
			W: size.Width,
			H: size.Height,
		})

		y += margin.T + size.Height + margin.B
	}
}
