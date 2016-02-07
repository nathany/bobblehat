// Package texture loads and creates abitrary sized RGB565 images.
package texture

import (
	"image"
	_ "image/png" // to load PNG files
	"os"

	"github.com/nathany/bobblehat/sense/screen/color"
)

// Texture is an RGB565 image of abitrary size.
type Texture struct {
	width, height int
	Pixels        []color.Color
}

// New texture allocates an empty texture.
func New(width, height int) *Texture {
	return &Texture{
		width:  width,
		height: height,
		Pixels: make([]color.Color, width*height),
	}
}

// SetPixel in the texture.
func (tx *Texture) SetPixel(x, y int, c color.Color) {
	tx.Pixels[y*tx.width+x] = c
}

// GetPixel from the texture.
func (tx *Texture) GetPixel(x, y int) color.Color {
	return tx.Pixels[y*tx.width+x]
}

// Width of the texture
func (tx *Texture) Width() int {
	return tx.width
}

// Height of the texture
func (tx *Texture) Height() int {
	return tx.height
}

// Blit copies a texture to another.
func Blit(dst *Texture, dstX, dstY int, src *Texture, srcX, srcY, width, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dst.SetPixel(dstX+x, dstY+y, src.GetPixel(srcX+x, srcY+y))
		}
	}
}

// Load a texture from a PNG.
func Load(name string) (*Texture, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	b := img.Bounds().Size()
	tx := New(b.X, b.Y)

	for y := 0; y < b.Y; y++ {
		for x := 0; x < b.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			tx.Pixels[y*tx.width+x] = color.New(uint8(r>>8), uint8(g>>8), uint8(b>>8))
		}
	}
	return tx, nil
}
