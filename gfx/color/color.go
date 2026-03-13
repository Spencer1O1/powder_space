package gfxcolor

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func RGBA(r, g, b, a uint8) Color {
	return Color{r, g, b, a}
}
