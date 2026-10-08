package ui

import (
	"swing-go/fonts"
)

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

func (c *Canvas) DrawGlyph(x, y int, glyph fonts.Glyph) {
	for py := 0; py < glyph.Y; py++ {
		for px := 0; px < glyph.X; px++ {
			alpha := glyph.Mask[py*glyph.Stride+px]

			c.blendPixel(
				x+px,
				y+py,
				Color{R: 255, G: 255, B: 255, A: alpha},
			)
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

func (c *Canvas) blendPixel(x, y int, src Color) {
	offset := y*c.Stride + x*4

	dst := Color{
		B: c.Pixels[offset+0],
		G: c.Pixels[offset+1],
		R: c.Pixels[offset+2],
		A: c.Pixels[offset+3],
	}

	a := uint16(src.A)
	invA := uint16(255 - src.A)

	c.Pixels[offset+0] = uint8(
		(uint16(src.B)*a + uint16(dst.B)*invA) / 255,
	)
	c.Pixels[offset+1] = uint8(
		(uint16(src.G)*a + uint16(dst.G)*invA) / 255,
	)
	c.Pixels[offset+2] = uint8(
		(uint16(src.R)*a + uint16(dst.R)*invA) / 255,
	)
	c.Pixels[offset+3] = uint8(
		uint16(src.A) + uint16(dst.A)*invA/255,
	)
}
