package main

import (
	"flag"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/mcaci/lets-go-workshop/myimage"
)

func main() {
	// flags for input rectangle
	backR := flag.Uint("backR", 0, "Red value for the background color")
	backG := flag.Uint("backG", 0, "Green value for the background color")
	backB := flag.Uint("backB", 0, "Blue value for the background color")
	// flags for the text
	textR := flag.Uint("textR", 255, "Red value for the color of the text")
	textG := flag.Uint("textG", 255, "Green value for the color of the text")
	textB := flag.Uint("textB", 255, "Blue value for the color of the text")
	textFont := flag.String("textFont", "./resources/fonts/Ubuntu-R.ttf", "Blue value for the color of the text")
	textFontSize := flag.Float64("textFontSize", 32, "Blue value for the color of the text")
	// parse flags after defining all of them
	flag.Parse()
	f, err := os.Create("./out/rect.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	text := strings.Join(flag.Args(), " ")
	l, h := myimage.TextBounds(32, len(text), 10)
	r := myimage.New(l, h, color.RGBA{R: uint8(*backR), G: uint8(*backG), B: uint8(*backB), A: 255})
	err = myimage.Write(r, text, color.RGBA{R: uint8(*textR), G: uint8(*textG), B: uint8(*textB), A: 255}, *textFont, *textFontSize)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(f, r)
	if err != nil {
		log.Fatal(err)
	}
}
