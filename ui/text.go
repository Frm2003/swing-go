package ui

import (
	"fmt"
	"swing-go/fonts"
)

type Text struct {
	Color     Color
	Content   string
	Margin    Edge
	Padding   Edge
	TextStyle fonts.Query

	FontManager *fonts.Manager
	rect        Rect
}

type TextStyle struct {
	Family string
	Slant  fonts.Slant
	Weight fonts.Weight
}

func (t *Text) Draw(c *Canvas) {
	cursorX := t.rect.X

	for _, char := range t.Content {
		glyph := t.FontManager.Glyph(
			t.rect.X,
			t.rect.Y,
			char,
			t.TextStyle,
		)

		c.DrawGlyph(cursorX, t.rect.Y, glyph)

		cursorX += glyph.Advance
	}
}

func (t *Text) GetMargin() Edge {
	return t.Margin
}

// qual será meu tamanho t onde vou ficar?
// Recebe o espaço/posição definido pelo pai.
func (t *Text) Layout(rect Rect) {
	t.rect = rect
}

// Calcula o tamanho do componente baseado nos filhos, conteúdo, padding etc.
func (t *Text) Measure() Size {
	var h, w int

	for _, char := range t.Content {
		advance, height := t.FontManager.Measure(char, t.TextStyle)

		fmt.Println(char)

		if height > h {
			h = height
		}

		w += advance
	}

	return Size{
		Width:  w,
		Height: h,
	}
}
