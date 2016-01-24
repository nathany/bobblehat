// Package screen provides access to the Sense HAT's 8x8 LED matrix.
package screen

import (
	"encoding/binary"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gophergala2016/bobblehat/sense/screen/color"
)

// FrameBuffer for 8x8 LED Matrix (row, column)
type FrameBuffer [8][8]color.Color

// Draw a buffer to the LED matrix screen.
func Draw(fb *FrameBuffer) error {
	return draw(displayDevice, fb)
}

// Clear the screen (off).
func Clear() error {
	return draw(displayDevice, new(FrameBuffer))
}

// the LED matrix screen
var displayDevice string

func init() {
	var err error
	displayDevice, err = getDevice("RPi-Sense FB")
	if err != nil {
		panic(err)
	}
}

func draw(backBuffer string, fb *FrameBuffer) error {
	f, err := os.Create(backBuffer)
	if err != nil {
		return err
	}
	defer f.Close()

	return binary.Write(f, binary.LittleEndian, fb)
}

// getDevice finds the named frame buffer.
func getDevice(name string) (string, error) {
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
	return "", errFrameBufferNotFound
}

// errors
var (
	errFrameBufferNotFound = errors.New("frame buffer device not found")
)
