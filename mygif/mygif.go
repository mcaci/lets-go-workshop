package mygif

import (
	"image"
	"image/color"
	"image/gif"
	"io"

	"github.com/mcaci/lets-go-workshop/myimage"
)

func New(text string, c1, c2 color.RGBA, fontPath string, fontSize float64) ([]*image.Paletted, error) {
	const nFrames = 2
	l, h := myimage.TextBounds(int(fontSize), len(text), 10)
	var frames []*image.Paletted
	for i := 0; i < nFrames; i++ {
		var bgColor, fgColor color.RGBA
		switch i % 2 {
		case 0:
			bgColor, fgColor = c1, c2 // same as params
		default:
			bgColor, fgColor = c2, c1 // switch
		}
		frame := myimage.New(l, h, bgColor)
		err := myimage.Write(frame, text, fgColor, fontPath, fontSize)
		if err != nil {
			return nil, err
		}
		frames = append(frames, frame)
	}
	return frames, nil
}

func Save(w io.Writer, frames []*image.Paletted, delay int) error {
	delays := make([]int, len(frames))
	for i := range delays {
		delays[i] = delay
	}
	return gif.EncodeAll(w, &gif.GIF{
		Image: frames,
		Delay: delays,
	})
}
