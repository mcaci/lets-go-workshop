package main

import (
	"flag"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/mcaci/lets-go-workshop/myimage"
)

func main() {
	// flags for input rectangle
	l := flag.Int("l", 200, "length of the image")
	h := flag.Int("h", 100, "height of the image")
	backR := flag.Uint("backR", 0, "Red value for the background color")
	backG := flag.Uint("backG", 0, "Green value for the background color")
	backB := flag.Uint("backB", 0, "Blue value for the background color")
	// parse flags after defining all of them
	flag.Parse()
	f, err := os.Create("./out/rect.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := myimage.New(*l, *h, color.RGBA{R: uint8(*backR), G: uint8(*backG), B: uint8(*backB), A: 255})
	err = png.Encode(f, r)
	if err != nil {
		log.Fatal(err)
	}
}
