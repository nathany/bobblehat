package main

import (
	"time"

	"github.com/gophergala2016/bobblehat/sense/screen"
	"github.com/gophergala2016/bobblehat/sense/screen/color"
)

func main() {
	var fb screen.FrameBuffer
	fb[0][0] = color.Red
	fb[0][1] = color.Green
	fb[0][2] = color.Blue
	fb[0][3] = color.New(255, 0, 255) // Magenta
	fb[0][4] = color.New(255, 255, 0) // Yellow
	fb[0][5] = color.New(0, 255, 255) // Cyan
	fb[0][6] = color.White

	// gradients
	var c uint8
	var x int
	for i := 1; i <= 8; i++ {
		x = i - 1
		c = uint8(i*32 - 1)
		fb[1][x] = color.New(c, c, c)
		fb[2][x] = color.New(c, 0, 0)
		fb[3][x] = color.New(0, c, 0)
		fb[4][x] = color.New(0, 0, c)
		fb[5][x] = color.New(c, 0, c)
		fb[6][x] = color.New(c, c, 0)
		fb[7][x] = color.New(0, c, c)
	}

	screen.Draw(&fb)

	time.Sleep(time.Second * 5)
	screen.Clear()
}
