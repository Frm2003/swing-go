package fonts

type Glyph struct {
	Advance int
	Mask    []byte
	Stride  int
	X, Y    int
}
