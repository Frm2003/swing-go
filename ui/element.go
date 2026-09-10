package ui

type Element struct {
	Children []Widget

	Background Color
	Display    Direction
	Height     int
	Margin     Edge
	Padding    Edge
	Width      int

	rect Rect
}

func (e *Element) Draw(c *Canvas) {
	// Alpha 0 (valor zero de Color, quando Background não é definido)
	// significa "sem background": deixa o que está atrás aparecer, em vez
	// de pintar um retângulo transparente por cima (Canvas não faz blend).
	if e.Background.A > 0 {
		c.FillRect(
			e.rect.X,
			e.rect.Y,
			e.rect.W,
			e.rect.H,
			e.Background,
		)
	}

	for _, child := range e.Children {
		child.Draw(c)
	}
}

func (e *Element) GetMargin() Edge {
	return e.Margin
}

// qual será meu tamanho e onde vou ficar?
// Recebe o espaço/posição definido pelo pai.
func (e *Element) Layout(rect Rect) {
	e.rect = rect

	y := rect.Y + e.Padding.T
	x := rect.X + e.Padding.L

	for _, child := range e.Children {
		size := child.Measure()
		margin := child.GetMargin()

		child.Layout(Rect{
			X: x + margin.L,
			Y: y + margin.T,
			W: size.Width,
			H: size.Height,
		})

		advance := mainAxis(e.Display, Size{
			Width:  margin.L + size.Width + margin.R,
			Height: margin.T + size.Height + margin.B,
		})

		dx, dy := axisDelta(e.Display, advance)

		x += dx
		y += dy
	}
}

// Calcula o tamanho do componente baseado nos filhos, conteúdo, padding etc.
func (e *Element) Measure() Size {
	size := Size{}

	for _, child := range e.Children {
		childSize := child.Measure()
		margin := child.GetMargin()

		childSize.Width += margin.L + margin.R
		childSize.Height += margin.T + margin.B

		size = measure(
			e.Display,
			size,
			childSize,
		)
	}

	size.Width += e.Padding.L + e.Padding.R
	size.Height += e.Padding.T + e.Padding.B

	if e.Width > size.Width {
		size.Width = e.Width
	}

	if e.Height > size.Height {
		size.Height = e.Height
	}

	return size
}
