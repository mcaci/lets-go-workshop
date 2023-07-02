package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/color"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"unicode"

	"github.com/mcaci/lets-go-workshop/mygif"
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

	ext := flag.String("ext", "gif", "file extension (png or gif)")
	in := flag.String("in", "", "input file with list of texts to transform to gif")
	// parse flags after defining all of them
	flag.Parse()

	var r io.Reader
	switch {
	case *in != "":
		var err error
		r, err = os.Open(*in)
		if err != nil {
			log.Fatal(err)
		}
	default:
		text := strings.Join(flag.Args(), " ")
		r = strings.NewReader(text)
	}

	var wg sync.WaitGroup
	for text := range readInput(r) {
		text := text
		wg.Add(1)
		go func() {
			log.Printf("start %q", text)
			f, err := os.Create(fmt.Sprintf("./out/%s.%s", strings.Map(func(r rune) rune {
				if !unicode.IsLetter(r) {
					return '-'
				}
				return r
			}, text), *ext))
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			c1 := color.RGBA{R: uint8(*backR), G: uint8(*backG), B: uint8(*backB), A: 255}
			c2 := color.RGBA{R: uint8(*textR), G: uint8(*textG), B: uint8(*textB), A: 255}
			gif, err := mygif.New(text, c1, c2, *textFont, *textFontSize)
			if err != nil {
				log.Fatal(err)
			}
			err = gif.Save(f)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("done %q", text)
			wg.Done()
		}()
	}
	wg.Wait()
}

func readInput(r io.Reader) chan string {
	texts := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			texts <- s.Text()
		}
		close(texts)
	}()
	return texts
}
