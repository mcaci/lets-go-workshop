package mygif

import (
	"image"
	"image/color"
	"image/gif"
	"io"

	"github.com/mcaci/lets-go-workshop/myimage"
)

type GIF struct {
	frames []*image.Paletted
	delay  int
}

func New(text string, c1, c2 color.RGBA, fontPath string, fontSize float64) (GIF, error) {
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
			return GIF{}, err
		}
		frames = append(frames, frame)
	}
	return GIF{frames: frames, delay: 100}, nil
}

func (g GIF) Save(w io.Writer) error {
	delays := make([]int, len(g.frames))
	for i := range delays {
		delays[i] = g.delay
	}
	return gif.EncodeAll(w, &gif.GIF{
		Image: g.frames,
		Delay: delays,
	})
}
