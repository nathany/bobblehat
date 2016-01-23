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
	r := uint16(red>>3) & 0x1F
	g := uint16(green>>2) & 0x3F
	b := uint16(blue>>3) & 0x1F
	return Color((r << 11) + (g << 5) + b)
}
