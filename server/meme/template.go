package meme

import (
	"image"

	"github.com/fogleman/gg"
)

type Template struct {
	Name      string
	Image     image.Image
	TextSlots []*TextSlot
}

func (t *Template) Render(text []string) (image.Image, error) {
	dc := gg.NewContextForImage(t.Image)

	write, copy := t.filterSlots()
	for i, slot := range write {
		if i >= len(text) {
			break
		}
		slot.Render(dc, text[i], DEBUG)
	}

	for _, slot := range copy {
		if slot.Copy <= len(text) {
			slot.Render(dc, text[slot.Copy-1], DEBUG)
		}
	}

	return dc.Image(), nil
}

func (t *Template) filterSlots() (writeSlots []*TextSlot, copySlots []*TextSlot) {

	for _, s := range t.TextSlots {
		if s.Copy == 0 {
			writeSlots = append(writeSlots, s)
		} else {
			copySlots = append(copySlots, s)
		}
	}
	return writeSlots, copySlots
}
