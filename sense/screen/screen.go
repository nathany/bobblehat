// Package screen provides access to the Sense HAT's 8x8 LED matrix.
package screen

import (
	"encoding/binary"
	"errors"
	"image"
	"image/color"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	rgb565color "github.com/nathany/bobblehat/sense/screen/color"
	"github.com/nathany/bobblehat/sense/screen/texture"
)

// FrameBuffer is an 8x8 texture that can draw to the LED Matrix.
type FrameBuffer struct {
	*texture.Texture
}

// NewFrameBuffer creates a back buffer for the screen.
func NewFrameBuffer(options ...func(*FrameBuffer)) *FrameBuffer {
	fb := &FrameBuffer{
		Texture: texture.New(8, 8),
	}

	for _, o := range options {
		o(fb)
	}

	return fb
}

// ColorModel of the frame buffer
func (fb *FrameBuffer) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds of the frame buffer (it is 8x8)
func (fb *FrameBuffer) Bounds() image.Rectangle {
	return image.Rect(0, 0, 8, 8)
}

// At returns the color of the LED at x,y
func (fb *FrameBuffer) At(x, y int) color.Color {
	if x < 0 || y < 0 || x > 7 || y > 7 {
		return color.RGBA{}
	}

	sc := uint16(fb.GetPixel(x, y))

	r := (sc & 0xF800) >> 8
	g := (sc & 0x07E0) >> 3
	b := (sc & 0x001F) << 3

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 0xff,
	}
}

// Set the color of the LED at x,y
func (fb *FrameBuffer) Set(x, y int, c color.Color) {
	if x < 0 || y < 0 || x > 7 || y > 7 {
		return
	}

	r, g, b, _ := c.RGBA()

	if fb.Texture == nil {
		fb.Texture = texture.New(8, 8)
	}

	fb.Texture.SetPixel(x, y, rgb565color.New(uint8(r>>8), uint8(g>>8), uint8(b>>8)))
}

// SetImage sets the frame buffer to the provided image
func (fb *FrameBuffer) SetImage(m image.Image) {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			r, g, b, _ := m.At(x, y).RGBA()
			fb.Texture.SetPixel(x, y, rgb565color.New(uint8(r>>8), uint8(g>>8), uint8(b>>8)))
		}
	}
}

// DrawImage draws an image to the LED matrix screen.
func DrawImage(m image.Image) error {
	return Draw(NewFrameBuffer(withImage(m)))
}

func withImage(m image.Image) func(*FrameBuffer) {
	return func(fb *FrameBuffer) {
		fb.SetImage(m)
	}
}

// Draw a buffer to the LED matrix screen.
func Draw(fb *FrameBuffer) error {
	return draw(displayDevice, fb)
}

var blankFrameBuffer = NewFrameBuffer()

// Clear the screen (off).
func Clear() error {
	return draw(displayDevice, blankFrameBuffer)
}

// the LED matrix screen
var displayDevice string

func init() {
	var err error

	displayDevice, err = getDevice("RPi-Sense FB")
	if err != nil {
		os.Stdout.WriteString(err.Error() + "\n")
	}
}

func draw(backBuffer string, fb *FrameBuffer) error {
	f, err := os.Create(backBuffer)
	if err != nil {
		return err
	}
	defer f.Close()

	return binary.Write(f, binary.LittleEndian, fb.Texture.Pixels)
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
