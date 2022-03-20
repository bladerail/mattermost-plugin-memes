package meme

import (
	"image"
	"image/color"
	"strings"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type HorizontalAlignment gg.Align

const (
	Left   HorizontalAlignment = -1
	Center HorizontalAlignment = 0
	Right  HorizontalAlignment = 1
)

type VerticalAlignment int

const (
	Top    VerticalAlignment = -1
	Middle VerticalAlignment = 0
	Bottom VerticalAlignment = 1
)

type TextSlot struct {
	Bounds              image.Rectangle
	Font                *truetype.Font
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
	face := faceForSlot(text, s.Font, s.MaxFontSize, s.Bounds.Dx(), s.Bounds.Dy())
	dc.SetFontFace(face)

	rotCenterX := float64((s.Bounds.Min.X + s.Bounds.Dx()) / 2)
	rotCenterY := float64((s.Bounds.Min.Y + s.Bounds.Dy()) / 2)

	dc.RotateAbout(gg.Radians(s.Rotation),
		rotCenterX, rotCenterY)
	dc.SetColor(s.OutlineColor)

	outlineWidth := s.OutlineWidth // "stroke" size
	if outlineWidth == 0 {
		outlineWidth = 2
	}
	for dy := -outlineWidth; dy <= outlineWidth; dy++ {
		for dx := -outlineWidth; dx <= outlineWidth; dx++ {
			if dx*dx+dy*dy >= outlineWidth {
				// give it rounded corners
				continue
			}
			x := float64(s.Bounds.Min.X) + float64(dx)
			y := float64(s.Bounds.Min.Y) + float64(dy)
			dc.DrawStringWrapped(text, x, y, 0, 0, float64(s.Bounds.Dx()), 1.0, gg.AlignLeft)
		}
	}

	dc.SetColor(s.TextColor)
	dc.DrawStringWrapped(text,
		float64(s.Bounds.Min.X),
		float64(s.Bounds.Min.Y),
		0, 0, float64(s.Bounds.Dx()), 1.0, gg.AlignLeft)

	dc.Pop()
}

func faceForSlot(text string, font *truetype.Font, maxFontSize float64, width int, height int) font.Face {
	fontSize := maxFontSize
	if fontSize == 0.0 {
		fontSize = 80
	}

	dc := gg.NewContext(10, 10)
	face := truetype.NewFace(font, &truetype.Options{
		Size: fontSize,
	})
	for fontSize >= 6.0 {
		face = truetype.NewFace(font, &truetype.Options{
			Size: fontSize,
		})
		dc.SetFontFace(face)
		lines := dc.WordWrap(text, float64(width))

		w, h := dc.MeasureMultilineString(strings.Join(lines, "\n"), 1.0)

		if w > float64(width) || h > float64(height) {
			fontSize -= (fontSize + 9) / 10
			continue
		}
		break
	}
	return face
}
