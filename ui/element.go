package ui

type Element struct {
	Children []Widget

	Background Color
	Display    Direction
	Height     int
	Margin     Edge
	Padding    Edge
	Width      int

	offsets []Point
	size    Size
}

func (e *Element) Draw(c *Canvas, origin Point) {
	if e.Background.A > 0 {
		c.FillRect(
			origin.X,
			origin.Y,
			e.size.Width,
			e.size.Height,
			e.Background,
		)
	}

	for index, child := range e.Children {
		offset := e.offsets[index]

		x := offset.X + origin.X
		y := offset.Y + origin.Y

		child.Draw(c, Point{x, y})
	}
}

func (e *Element) GetMargin() Edge {
	return e.Margin
}

func (e *Element) Layout(c Constraint) Size {
	e.offsets = make([]Point, 0, len(e.Children))

	y, x := 0, 0
	content := Size{}

	for _, child := range e.Children {
		childMargin := child.GetMargin()

		childSize := child.Layout(Constraint{
			MaxH: c.MaxH - y - (e.Padding.T + e.Padding.B),
			MaxW: c.MaxW - x - (e.Padding.L + e.Padding.R),
		})

		e.offsets = append(e.offsets, Point{
			x + e.Padding.L + childMargin.L,
			y + e.Padding.T + childMargin.T,
		})

		outerSize := Size{
			Height: childMargin.T + childSize.Height + childMargin.B,
			Width:  childMargin.L + childSize.Width + childMargin.R,
		}

		content = measure(e.Display, content, outerSize)

		dx, dy := axisDelta(e.Display, outerSize)

		x += dx
		y += dy
	}

	e.size = Size{
		Height: e.Padding.T + max(e.Height, content.Height) + e.Padding.B,
		Width:  e.Padding.L + max(e.Width, content.Width) + e.Padding.R,
	}

	return e.size
}
