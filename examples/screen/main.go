package main

import (
	"encoding/binary"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Color of a pixel (16 bit RGB565)
type Color uint16

// NewColor from RGB components (0-255).
func NewColor(red, green, blue uint8) Color {
	r := uint16(red>>3) & 0x1F
	g := uint16(green>>2) & 0x3F
	b := uint16(blue>>3) & 0x1F
	return Color((r << 11) + (g << 5) + b)
}

// errors
var (
	ErrFrameBufferNotFound = errors.New("frame buffer not found")
)

// frameBuffer finds the named frame buffer
func frameBuffer(name string) (string, error) {
	matches, err := filepath.Glob("/sys/class/graphics/fb*")
	if err != nil {
		return "", err
	}

	for _, dir := range matches {
		b, err := ioutil.ReadFile(filepath.Join(dir, "name"))
		if err != nil {
			continue
		}
		fbName := strings.TrimSpace(string(b))
		if fbName == name {
			dev := filepath.Join("/dev", filepath.Base(dir))
			return dev, nil
		}
	}
	return "", ErrFrameBufferNotFound
}

// SetPixel on the LED matrix.
func SetPixel(fb string, x, y int, color Color) {
	f, err := os.Create(fb)
	if err != nil {
		log.Fatal(err)
	}

	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(color))

	_, err = f.WriteAt(b, int64((y*8+x)*2))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Find the frame buffer for the Sense HAT
	fb, err := frameBuffer("RPi-Sense FB")
	if err != nil {
		log.Fatal(err)
	}

	SetPixel(fb, 0, 0, NewColor(255, 0, 0))
	SetPixel(fb, 1, 1, NewColor(0, 255, 0))
	SetPixel(fb, 2, 2, NewColor(0, 0, 255))
}
