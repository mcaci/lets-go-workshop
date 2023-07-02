package myimage

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"os"

	"github.com/golang/freetype"
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

func Write(dst draw.Image, text string, c color.RGBA, fontPath string, fontSize float64) error {
	ctx := freetype.NewContext()
	fontBytes, err := os.ReadFile(fontPath)
	if err != nil {
		return err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return err
	}
	ctx.SetClip(dst.Bounds())
	ctx.SetDPI(72)
	ctx.SetDst(dst)
	ctx.SetFont(f)
	ctx.SetFontSize(fontSize)
	ctx.SetSrc(image.NewUniform(c))
	_, err = ctx.DrawString(text, freetype.Pt(10, 30))
	return err
}
