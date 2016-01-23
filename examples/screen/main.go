package main

import (
	"time"

	"github.com/gophergala2016/bobblehat/sense/screen"
	"github.com/gophergala2016/bobblehat/sense/screen/color"
)

func main() {
	var frontBuffer screen.FrameBuffer
	frontBuffer[0][0] = color.Red
	frontBuffer[0][1] = color.Green
	frontBuffer[0][2] = color.Blue
	frontBuffer[0][3] = color.New(128, 0, 128)
	screen.Draw(&frontBuffer)

	time.Sleep(time.Second * 3)
	screen.Clear()
}
