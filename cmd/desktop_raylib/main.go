package main

import (
	"log"

	"github.com/Spencer1O1/powder_space/v2/engine"
	"github.com/Spencer1O1/powder_space/v2/game"
	rr "github.com/Spencer1O1/powder_space/v2/renderer/raylib"
)

func main() {
	window := rr.NewWindow(1280, 720, "PowderSpace")
	defer window.Close()

	renderer := rr.NewRenderer()
	g := game.NewGame()

	app := engine.NewApp(window, renderer, g)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
