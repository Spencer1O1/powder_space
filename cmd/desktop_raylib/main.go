package main

import (
	"log"

	"github.com/Spencer1O1/powder_space/v2/engine"
	"github.com/Spencer1O1/powder_space/v2/game"
	ri "github.com/Spencer1O1/powder_space/v2/inputx/raylib"
	rr "github.com/Spencer1O1/powder_space/v2/renderer/raylib"
)

func main() {
	window := rr.NewWindow(1920, 1080, "PowderSpace")
	defer window.Close()

	renderer := rr.NewRenderer(1920, 1080)
	input := ri.NewInput()
	g := game.NewGame()

	app := engine.NewApp(window, renderer, input, g)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
