package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gophergala2016/bobblehat/sense/screen"
	"github.com/gophergala2016/bobblehat/sense/screen/color"
	"github.com/gophergala2016/bobblehat/sense/stick"
)

func main() {
	fb := screen.NewFrameBuffer()

	var count int
	for {
		count = stick.ReadEvent()
		fmt.Printf("Count is-->%d", count)

		rand.Seed(time.Now().UnixNano())
		direction := rand.Intn(5)

		switch direction {
		case stick.Right:
			fmt.Println("right")
			for i := 0; i < 8; i++ {
				for j := 4; j < 8; j++ {
					fb.SetPixel(i, j, color.Red)
				}
			}
			screen.Draw(fb)
			break
		case stick.Left:
			fmt.Println("left")
			for i := 0; i < 8; i++ {
				for j := 0; j < 4; j++ {
					fb.SetPixel(i, j, color.Black)
				}
			}
			screen.Draw(fb)
			break
		case stick.Up:
			fmt.Println("up")
			for i := 0; i < 4; i++ {
				for j := 0; j < 8; j++ {
					fb.SetPixel(i, j, color.Blue)
				}
			}
			screen.Draw(fb)
			break
		case stick.Down:
			fmt.Println("down")
			for i := 4; i < 8; i++ {
				for j := 0; j < 8; j++ {
					fb.SetPixel(i, j, color.Green)
				}
			}
			screen.Draw(fb)
			break
		case stick.Enter:
			fmt.Println("enter")
			for i := 0; i < 8; i++ {
				for j := 0; j < 8; j++ {
					fb.SetPixel(i, j, color.White)
				}
			}
			screen.Draw(fb)
			break
		default:
			fmt.Println("waiting to be pressed")
		}
		time.Sleep(10)
	}

}
