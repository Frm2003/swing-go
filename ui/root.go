package ui

type Root struct {
	Backgroud Color
	Childrens []Widget
}

func (r *Root) Draw(c *Canvas) {
	c.Clear(r.Backgroud)

	for _, child := range r.Childrens {
		child.Draw(c)
	}
}

func (r *Root) Layout() {
	y := 0

	for _, child := range r.Childrens {
		size := child.Measure()

		child.Layout(Rect{
			X: 0,
			Y: y,
			W: size.Width,
			H: size.Height,
		})

		y += size.Height
	}
}
