package screen

import (
	"image/color"
	"testing"
)

func TestFrameBufferAsImage(t *testing.T) {
	fb := NewFrameBuffer()

	// Black pixel
	if got, want := fb.At(0, 0), (color.RGBA{0, 0, 0, 255}); got != want {
		t.Fatalf("fb.At(0,0) = %+v, want %+v", got, want)
	}

	// Pixel changed to all red
	fb.Set(0, 0, color.RGBA{255, 0, 0, 255})

	// Conversion back from RGB565 is lossy
	if got, want := fb.At(0, 0), (color.RGBA{248, 0, 0, 255}); got != want {
		t.Fatalf("fb.At(0,0) = %+v, want %+v", got, want)
	}

	// Pixel changed to a red color that can be converted back and forth between RGB565 and RGB888
	fb.Set(0, 0, color.RGBA{248, 0, 0, 255})

	if got, want := fb.At(0, 0), (color.RGBA{248, 0, 0, 255}); got != want {
		t.Fatalf("fb.At(0,0) = %+v, want %+v", got, want)
	}
}
