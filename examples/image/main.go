package main

import (
	_ "image/png"
	"time"

	"github.com/gophergala2016/bobblehat/sense/screen"
	"github.com/gophergala2016/bobblehat/sense/screen/texture"

	"log"
)

func main() {
	fb := screen.NewFrameBuffer()

	tx, err := texture.Load("gopher16x16.png")
	fatalIfError(err)

	// display top corner of image and wait a bit
	texture.Blit(fb.Texture, 0, 0, tx, 0, 0, 8, 8)
	screen.Draw(fb)
	time.Sleep(500 * time.Millisecond)

	// pan around the image
	state := 1
	var xo, yo int
	for {
		texture.Blit(fb.Texture, 0, 0, tx, xo, yo, 8, 8)
		screen.Draw(fb)
		switch state {
		case 1: // down
			yo++
			if yo == tx.Height()-8 {
				state = 2
			}
		case 2: // right
			xo++
			if xo == tx.Width()-8 {
				state = 3
			}
		case 3: // up
			yo--
			if yo == 0 {
				state = 4
			}
		case 4: // left
			xo--
			if xo == 0 {
				state = 1
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func fatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
