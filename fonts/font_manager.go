package fonts

import (
	"image"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

type Manager struct {
	resolver resolver
}

func NewManager(resolver resolver) *Manager {
	return &Manager{
		resolver: resolver,
	}
}

func (m *Manager) Glyph(x, y int, char rune, query Query) Glyph {
	face, err := m.load(query)

	if err != nil {
		panic(err)
	}

	_, mask, _, advance, _ := face.Glyph(fixed.Point26_6{
		X: fixed.Int26_6(x),
		Y: fixed.Int26_6(y),
	}, char)

	alpha := mask.(*image.Alpha)

	return Glyph{
		Advance: advance.Round(),
		Mask:    alpha.Pix,
		Stride:  alpha.Stride,
		X:       alpha.Rect.Dx(),
		Y:       alpha.Rect.Dy(),
	}
}

func (m *Manager) load(query Query) (font.Face, error) {
	match, err := m.resolver.Match(query)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(match.File)
	if err != nil {
		return nil, err
	}

	collection, err := opentype.ParseCollection(data)
	if err != nil {
		return nil, err
	}

	font, err := collection.Font(match.Index)
	if err != nil {
		return nil, err
	}

	face, err := opentype.NewFace(font, &opentype.FaceOptions{
		Size: 36,
		DPI:  90,
	})

	if err != nil {
		return nil, err
	}

	return face, nil
}

func (m *Manager) Measure(char rune, query Query) (int, int) {
	face, err := m.load(query)

	if err != nil {
		panic(err)
	}

	advance, _ := face.GlyphAdvance(char)
	metrics := face.Metrics()

	return advance.Round(), metrics.Height.Round()
}
