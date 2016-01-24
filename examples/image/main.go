package main

import (
	"image"
	_ "image/png"

	"github.com/gophergala2016/bobblehat/sense/screen"
	"github.com/gophergala2016/bobblehat/sense/screen/color"

	"log"
	"os"
)

func main() {
	f, err := os.Open("gopher8x8.png")
	fatalIfError(err)
	defer f.Close()

	image, format, err := image.Decode(f)
	fatalIfError(err)

	b := image.Bounds().Size()
	log.Println(format, b)

	if b.X != 8 || b.Y != 8 {
		log.Fatal("Image is the wrong size.")
	}

	var fb screen.FrameBuffer
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			r, g, b, _ := image.At(x, y).RGBA()
			fb[y][x] = color.New(uint8(r>>8), uint8(g>>8), uint8(b>>8))
		}
	}
	screen.Draw(&fb)
}

func fatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
