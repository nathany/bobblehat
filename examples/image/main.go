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
	if len(os.Args) < 2 {
		log.Fatal("specify a png image to load.")
	}

	fb := screen.NewFrameBuffer()

	tx, err := texture.Load(os.Args[1])
	fatalIfError(err)

	// setup channel to handle Ctrl-C events
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM)

	// pan around the image
	state := 0
	var xo, yo int

loop:
	for {
		sleep := 100

		switch state {
		case 0: // nothing
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

		texture.Blit(fb.Texture, 0, 0, tx, xo, yo, 8, 8)
		screen.Draw(fb)

		// check if Ctrl-C was pressed
		select {
		case <-quit:
			break loop
		default:
		}

		time.Sleep(time.Duration(sleep) * time.Millisecond)
	}

	screen.Clear()
}

func fatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
