package ui

type Direction int

const (
	Column Direction = iota
	Row
)

func mainAxis(d Direction, s Size) int {
	if d == Row {
		return s.Width
	}
	return s.Height
}

func crossAxis(d Direction, s Size) int {
	if d == Row {
		return s.Height
	}
	return s.Width
}

func axisDelta(d Direction, delta int) (dx, dy int) {
	if d == Row {
		return delta, 0
	}
	return 0, delta
}

func measure(d Direction, current, child Size) Size {
	main := mainAxis(d, child)
	cross := crossAxis(d, child)

	// Pega o maior valor oposto a main axis que aparecer nos filhos
	if c := crossAxis(d, current); c > cross {
		cross = c
	}

	if d == Row {
		return Size{Width: current.Width + main, Height: cross}
	}

	return Size{Width: cross, Height: current.Height + main}
}
