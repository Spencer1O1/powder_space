package engine

import (
	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/game"
	gfxcolor "github.com/Spencer1O1/powder_space/v2/gfx/color"
	rr "github.com/Spencer1O1/powder_space/v2/renderer/raylib"
)

type App struct {
	window   *rr.Window
	renderer *rr.Renderer
	game     *game.Game
}

func NewApp(window *rr.Window, renderer *rr.Renderer, game *game.Game) *App {
	return &App{
		window:   window,
		renderer: renderer,
		game:     game,
	}
}

func (a *App) Run() error {
	for !a.window.ShouldClose() {
		a.handleInput()

		dt := rr.GetFrameTime()
		a.game.Update(dt)

		a.window.Begin()
		a.window.Clear(gfxcolor.Black)

		a.render()

		a.window.End()
	}

	return nil
}

func (a *App) handleInput() {
	// later: mouse tools, camera movement, hotkeys
}

func (a *App) render() {
	a.renderer.DrawText(content.TitleString, 20, 20, 32, gfxcolor.White)
}
