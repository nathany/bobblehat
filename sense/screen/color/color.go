// Package color defines a 16-bit RGB565 color.
package color

// Color of a pixel (16 bit RGB565)
type Color uint16

// Predefined colors
var (
	Red   = New(255, 0, 0)
	Green = New(0, 255, 0)
	Blue  = New(0, 0, 255)
	Black = New(0, 0, 0) // Off
	White = New(255, 255, 255)
)

// New color from RGB components (0-255).
func New(red, green, blue uint8) Color {
	r := red >> 3
	g := green >> 2
	b := blue >> 3
	return (Color(r) << 11) + (Color(g) << 5) + Color(b)
}
