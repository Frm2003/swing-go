package ui

type Div struct {
	Backgroud Color
	Children  []Widget

	Height int
	Width  int

	Padding Edge
	Margin  Edge

	rect Rect
}

func (d *Div) Draw(c *Canvas) {
	c.FillRect(
		d.rect.X,
		d.rect.Y,
		d.rect.W,
		d.rect.H,
		d.Backgroud,
	)

	for _, child := range d.Children {
		child.Draw(c)
	}
}

func (d *Div) GetMargin() Edge {
	return d.Margin
}

// qual será meu tamanho e onde vou ficar?
// Recebe o espaço/posição definido pelo pai.
func (d *Div) Layout(rect Rect) {
	d.rect = rect

	y := rect.Y + d.Padding.T
	x := rect.X + d.Padding.L

	for _, child := range d.Children {
		size := child.Measure()
		margin := child.GetMargin()

		child.Layout(Rect{
			X: x + margin.L,
			Y: y + margin.T,
			W: size.Width,
			H: size.Height,
		})

		y += margin.T + size.Height + margin.B
	}
}

// Calcula o tamanho do componente baseado nos filhos, conteúdo, padding etc.
func (d *Div) Measure() Size {
	w := d.Width
	h := d.Height

	for _, child := range d.Children {
		childSize := child.Measure()
		margin := child.GetMargin()

		childWidth := margin.L + childSize.Width + margin.R
		childHeight := margin.T + childSize.Height + margin.B

		if childWidth > w {
			w = childWidth
		}

		h += childHeight
	}

	return Size{
		Width:  w + d.Padding.L + d.Padding.R,
		Height: h + d.Padding.T + d.Padding.B,
	}
}
