package ui

type Edge struct {
	B, L, R, T int
}

type Size struct {
	Width  int
	Height int
}

type Constraint struct {
	MinH, MaxH int
	MinW, MaxW int
}

type Point struct {
	X, Y int
}
