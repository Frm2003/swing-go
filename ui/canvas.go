package ui

type Canvas struct {
	Height int
	Pixels []byte
	Stride int
	Width  int
}

func (c *Canvas) Clear(color Color) {
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			c.setPixel(x, y, color)
		}
	}
}

func (c *Canvas) FillRect(x, y, w, h int, color Color) {
	for py := y; py < y+h; py++ {
		for px := x; px < x+w; px++ {
			c.setPixel(px, py, color)
		}
	}
}

func (c *Canvas) setPixel(x, y int, color Color) {
	offset := y*c.Stride + x*4

	c.Pixels[offset+0] = color.B
	c.Pixels[offset+1] = color.G
	c.Pixels[offset+2] = color.R
	c.Pixels[offset+3] = color.A
}
