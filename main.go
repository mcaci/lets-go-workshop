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
	flag.Parse()
	f, err := os.Create("./out/rect.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := myimage.New(200, 100, color.RGBA{R: 255, A: 255})
	err = png.Encode(f, r)
	if err != nil {
		log.Fatal(err)
	}
}
