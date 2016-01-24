package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gophergala2016/bobblehat/sense/screen"
	"github.com/gophergala2016/bobblehat/sense/screen/color"
)

func main() {
	// create a new frame buffer
	fb := screen.NewFrameBuffer()

	// turn on LEDs on the first row
	fb.SetPixel(0, 0, color.Red)
	fb.SetPixel(1, 0, color.Green)
	fb.SetPixel(2, 0, color.Blue)
	fb.SetPixel(3, 0, color.New(255, 0, 255)) // Magenta
	fb.SetPixel(4, 0, color.New(255, 255, 0)) // Yellow
	fb.SetPixel(5, 0, color.New(0, 255, 255)) // Cyan
	fb.SetPixel(6, 0, color.White)

	// draw gradients
	var c uint8
	for x := 0; x < 8; x++ {
		c = uint8((x+1)*32 - 1)
		fb.SetPixel(x, 1, color.New(c, c, c))
		fb.SetPixel(x, 2, color.New(c, 0, 0))
		fb.SetPixel(x, 3, color.New(0, c, 0))
		fb.SetPixel(x, 4, color.New(0, 0, c))
		fb.SetPixel(x, 5, color.New(c, 0, c))
		fb.SetPixel(x, 6, color.New(c, c, 0))
		fb.SetPixel(x, 7, color.New(0, c, c))
	}

	// draw the frame buffer to the screen
	screen.Draw(fb)

	// wait for Ctrl-C, then clear the screen
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	screen.Clear()
}
