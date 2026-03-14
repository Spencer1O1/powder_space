package engine

import (
	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/game"
	gfxcolor "github.com/Spencer1O1/powder_space/v2/gfx/color"
	"github.com/Spencer1O1/powder_space/v2/inputx"
	rr "github.com/Spencer1O1/powder_space/v2/renderer/raylib"
)

type App struct {
	window      *rr.Window
	renderer    *rr.Renderer
	input       InputSource
	game        *game.Game
	accumulator float64

	mouseState inputx.MouseState

	physicsTick uint64
}

func NewApp(window *rr.Window, renderer *rr.Renderer, input InputSource, game *game.Game) *App {
	return &App{
		window:      window,
		renderer:    renderer,
		input:       input,
		game:        game,
		accumulator: 0,
		physicsTick: 0,
	}
}

func (a *App) Run() error {
	for !a.window.ShouldClose() {
		frameDt := float64(rr.GetFrameTime())
		if frameDt > maxAccumulatedFrameDt {
			frameDt = maxAccumulatedFrameDt
		}

		a.pollInput()

		a.game.Update(frameDt)

		a.accumulator += frameDt

		steps := 0
		for a.accumulator >= fixedDt && steps < maxFixedUpdatesPerFrame {
			a.fixedUpdate(fixedDt)
			a.accumulator -= fixedDt
			steps++
		}

		a.window.Begin()
		a.window.Clear(gfxcolor.Black)

		a.render()

		a.window.End()
	}

	return nil
}

func (a *App) fixedUpdate(dt float64) {
	a.physicsTick++

	if a.mouseState.LeftDown {
		if a.physicsTick%spawnFixedTickInterval == 0 {
			a.game.SpawnPowder(a.mouseState.Position)
		}
	}

	a.game.FixedUpdate(dt)
}

func (a *App) pollInput() {
	a.mouseState = a.input.PollMouse()

	if a.mouseState.RightPressed {
		a.game.SetAnchor(a.mouseState.Position)
	}
	if a.mouseState.RightReleased {
		a.game.ResetAnchor()
	}
}

func (a *App) render() {
	a.renderer.DrawText(content.TitleString, 20, 20, 32, gfxcolor.White)
	a.renderer.DrawGame(a.game, a.mouseState.Position)
}
