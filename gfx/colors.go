package gfx

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func RGBA(r, g, b, a uint8) Color {
	return Color{r, g, b, a}
}

var (
	White = RGBA(255, 255, 255, 255)
	Black = RGBA(0, 0, 0, 255)
	Gray  = RGBA(128, 128, 128, 255)
	Red   = RGBA(255, 0, 0, 255)
	Green = RGBA(0, 255, 0, 255)
	Blue  = RGBA(0, 0, 255, 255)
)
