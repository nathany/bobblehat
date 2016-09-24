package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/nathany/bobblehat/sense/screen"
	"github.com/nathany/bobblehat/sense/screen/color"
	"github.com/nathany/bobblehat/sense/stick"
)

var path = flag.String("path", "/dev/input/event0", "path to the event device")

func main() {
	// Parse command line flags
	flag.Parse()

	// Open the input device (and defer closing it)
	input, err := stick.Open(*path)
	if err != nil {
		fmt.Printf("Unable to open input device: %s\nError: %v\n", *path, err)
		os.Exit(1)
	}
	defer input.Close()

	// Clear the screen
	screen.Clear()

	// Print the name of the input device
	fmt.Println(input.Name())

	// Set up a signals channel (stop the loop using Ctrl-C)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)

	// Loop forever
	for {
		select {
		case <-signals:
			fmt.Println("")
			screen.Clear()

			// Exit the loop
			return
		case e := <-input.Events:
			fb := screen.NewFrameBuffer()

			switch e.Code {
			case stick.Enter:
				fmt.Println("⏎ ")
			case stick.Up:
				fmt.Println("↑")
				draw(fb, 0, 0, 8, 4, color.New(255, 255, 0))
			case stick.Down:
				fmt.Println("↓")
				draw(fb, 0, 4, 8, 8, color.New(255, 0, 0))
			case stick.Left:
				fmt.Println("←")
				draw(fb, 0, 0, 4, 8, color.New(0, 0, 255))
			case stick.Right:
				fmt.Println("→")
				draw(fb, 4, 0, 8, 8, color.New(0, 255, 0))
			}

			screen.Draw(fb)
		}
	}
}

func draw(fb *screen.FrameBuffer, a, b, m, n int, c color.Color) {
	for i := a; i < m; i++ {
		for j := b; j < n; j++ {
			fb.SetPixel(i, j, c)
		}
	}
}
