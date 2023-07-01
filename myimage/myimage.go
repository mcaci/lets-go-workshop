package myimage

import (
	"image"
	"image/color"
	"image/color/palette"
)

func New(l, h int, c color.RGBA) *image.Paletted {
	rect := image.Rect(0, 0, l, h)
	img := image.NewPaletted(rect, palette.Plan9)
	for i := 0; i < l; i++ {
		for j := 0; j < h; j++ {
			img.Set(i, j, c)
		}
	}
	return img
}
