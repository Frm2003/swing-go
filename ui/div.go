package ui

type Div struct {
	Backgroud Color
	Childrens []Widget
	Height    int
	Width     int

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

	for _, child := range d.Childrens {
		child.Draw(c)
	}
}

// qual será meu tamanho e onde vou ficar?
// Recebe o espaço/posição definido pelo pai.
func (d *Div) Layout(rect Rect) {
	d.rect = rect

	y := rect.Y

	for _, child := range d.Childrens {
		size := child.Measure()

		child.Layout(Rect{
			X: rect.X,
			Y: y,
			W: size.Width,
			H: size.Height,
		})

		y += size.Height
	}
}

// Calcula o tamanho do componente baseado nos filhos, conteúdo, padding etc.
func (d *Div) Measure() Size {
	w := d.Width
	h := d.Height

	for _, child := range d.Childrens {
		childSize := child.Measure()

		if childSize.Width > w {
			w = childSize.Width
		}

		h += childSize.Height
	}

	return Size{
		Width:  w,
		Height: h,
	}
}
