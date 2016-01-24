package main

import (
	"time"

	"github.com/gophergala2016/bobblehat/sense/screen"
	"github.com/gophergala2016/bobblehat/sense/screen/color"
)

func main() {
	fb := screen.NewFrameBuffer()
	fb.SetPixel(0, 0, color.Red)
	fb.SetPixel(1, 0, color.Green)
	fb.SetPixel(2, 0, color.Blue)
	fb.SetPixel(3, 0, color.New(255, 0, 255)) // Magenta
	fb.SetPixel(4, 0, color.New(255, 255, 0)) // Yellow
	fb.SetPixel(5, 0, color.New(0, 255, 255)) // Cyan
	fb.SetPixel(6, 0, color.White)

	// gradients
	var c uint8
	var x int
	for i := 1; i <= 8; i++ {
		x = i - 1
		c = uint8(i*32 - 1)
		fb.SetPixel(x, 1, color.New(c, c, c))
		fb.SetPixel(x, 2, color.New(c, 0, 0))
		fb.SetPixel(x, 3, color.New(0, c, 0))
		fb.SetPixel(x, 4, color.New(0, 0, c))
		fb.SetPixel(x, 5, color.New(c, 0, c))
		fb.SetPixel(x, 6, color.New(c, c, 0))
		fb.SetPixel(x, 7, color.New(0, c, c))
	}
	screen.Draw(fb)

	time.Sleep(time.Second * 5)
	screen.Clear()
}
