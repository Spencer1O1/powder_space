package main

import (
	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/gfx"
	rr "github.com/Spencer1O1/powder_space/v2/renderer/raylib"
)

func main() {
	window := rr.NewWindow(1280, 720, "PowderSpace")
	defer window.Close()

	renderer := rr.NewRenderer()

	for !window.ShouldClose() {
		window.Begin()
		window.Clear()

		renderer.DrawText(content.TitleString, 20, 20, 32, gfx.White)

		window.End()
	}
}
