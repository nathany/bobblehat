package main

import (
	_ "image/png"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gophergala2016/bobblehat/sense/screen"
	"github.com/gophergala2016/bobblehat/sense/screen/texture"

	"log"
)

func main() {
	// load a PNG specified on the command line (16x16 recommended).
	if len(os.Args) < 2 {
		log.Fatal("specify a png image to load.")
	}

	tx, err := texture.Load(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// create a new frame buffer
	fb := screen.NewFrameBuffer()

	// setup a channel to handle Ctrl-C events
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM)

	// scroll around the edges of the image.
	var xo, yo, state int
loop:
	for {
		// this is to pause for 33 milliseconds (30fps)
		// except during state changes, which will pause for half-a-second.
		sleep := 33

		switch state {
		case 0: // nothing (initial state)
			state = 1
			sleep = 500
		case 1: // down
			yo++
			if yo == tx.Height()-8 {
				state = 2
				sleep = 500
			}
		case 2: // right
			xo++
			if xo == tx.Width()-8 {
				state = 3
				sleep = 500
			}
		case 3: // up
			yo--
			if yo == 0 {
				state = 4
				sleep = 500
			}
		case 4: // left
			xo--
			if xo == 0 {
				state = 1
				sleep = 500
			}
		}

		// Blit texture to the frame buffer.
		texture.Blit(fb.Texture, 0, 0, tx, xo, yo, 8, 8)

		// Draw frame buffer to the screen.
		screen.Draw(fb)

		// Check if Ctrl-C was pressed
		// while waiting a few milliseconds.
		timer := time.NewTimer(time.Millisecond * time.Duration(sleep))
		select {
		case <-quit:
			timer.Stop()
			break loop
		case <-timer.C:
		}
	}

	screen.Clear()
}
