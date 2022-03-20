package meme

import (
	"image"
	"image/color"
	"strings"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

type HorizontalAlignment int

const (
	Left HorizontalAlignment = iota
	Center
	Right
)

type VerticalAlignment int

const (
	Top VerticalAlignment = iota
	Middle
	Bottom
)

type TextSlot struct {
	Bounds              image.Rectangle
	Font                *opentype.Font
	MaxFontSize         float64
	HorizontalAlignment HorizontalAlignment
	VerticalAlignment   VerticalAlignment
	TextColor           color.Color
	OutlineColor        color.Color
	OutlineWidth        int
	AllUppercase        bool
	Rotation            float64
}

func (s *TextSlot) Render(dc *gg.Context, text string) {
	if s.AllUppercase {
		text = strings.ToUpper(text)
	}
	dc.Push()
	// Compute font size by taking measurement of a text
	face, _, textHeight := faceForSlot(text, s.Font, s.MaxFontSize, s.Bounds.Dx(), s.Bounds.Dy())
	dc.SetFontFace(face)

	rotCenterX := float64((s.Bounds.Min.X + s.Bounds.Dx()) / 2)
	rotCenterY := float64((s.Bounds.Min.Y + s.Bounds.Dy()) / 2)

	dc.RotateAbout(gg.Radians(s.Rotation),
		rotCenterX, rotCenterY)
	dc.SetColor(s.OutlineColor)

	outlineWidth := s.OutlineWidth // "stroke" size
	if outlineWidth == 0 {
		outlineWidth = 8
	}

	textColor := s.TextColor
	if textColor == nil {
		textColor = color.Black
	}

	xStart := float64(s.Bounds.Min.X)
	xAlign := gg.Align(s.HorizontalAlignment)
	yStart := float64(s.Bounds.Min.Y)
	switch s.VerticalAlignment {
	case Top:
		yStart = float64(s.Bounds.Min.Y)
	case Middle:
		yStart = (float64(s.Bounds.Dy())-textHeight)/2.0 +
			float64(s.Bounds.Min.Y)
	case Bottom:
		yStart = float64(s.Bounds.Max.Y) - textHeight
	default:
		break
	}

	// Some black magic to draw outline
	if s.OutlineColor != nil {
		offset := face.Metrics().Height / 256 * fixed.Int26_6(outlineWidth)
		for _, delta := range []fixed.Point26_6{
			{X: offset, Y: offset},
			{X: -offset, Y: offset},
			{X: -offset, Y: -offset},
			{X: offset, Y: -offset},
		} {
			x := xStart + float64(delta.X)/64
			y := yStart + float64(delta.Y)/64
			dc.DrawStringWrapped(text, x, y, 0, 0, float64(s.Bounds.Dx()), 0.8, xAlign)
		}

	}
	// for dy := -outlineWidth; dy <= outlineWidth; dy++ {
	// 	for dx := -outlineWidth; dx <= outlineWidth; dx++ {
	// 		if dx*dx+dy*dy >= outlineWidth {
	// 			// give it rounded corners
	// 			continue
	// 		}

	// 		x := xStart + float64(dx)
	// 		y := yStart + float64(dy)
	// 		dc.DrawStringWrapped(text, x, y, 0, 0, float64(s.Bounds.Dx()), 1.0, xAlign)
	// 	}
	// }

	dc.SetColor(textColor)
	dc.DrawStringWrapped(text,
		xStart,
		yStart,
		0, 0, float64(s.Bounds.Dx()), 0.8, xAlign)

	dc.Pop()
}

func faceForSlot(text string, fontt *opentype.Font, maxFontSize float64, width int, height int) (font.Face, float64, float64) {
	fontSize := maxFontSize
	if fontSize == 0.0 {
		fontSize = 80
	}
	w := 0.0
	h := 0.0
	dc := gg.NewContext(10, 10)
	face, _ := opentype.NewFace(fontt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     fontSize * 0.25,
		Hinting: font.HintingNone,
	})
	for fontSize >= 6.0 {
		face, _ := opentype.NewFace(fontt, &opentype.FaceOptions{
			Size:    fontSize,
			DPI:     fontSize * 0.25,
			Hinting: font.HintingNone,
		})
		dc.SetFontFace(face)
		lines := dc.WordWrap(text, float64(width))

		w, h = dc.MeasureMultilineString(strings.Join(lines, "\n"), 1.0)

		if w > float64(width) || h > float64(height) {
			fontSize -= (fontSize + 9) / 10
			continue
		}
		break
	}

	return face, w, h
}
