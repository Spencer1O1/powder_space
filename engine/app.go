package engine

import (
	"github.com/Spencer1O1/powder_space/v2/content"
	"github.com/Spencer1O1/powder_space/v2/game"
	gfxcolor "github.com/Spencer1O1/powder_space/v2/gfx/color"
	"github.com/Spencer1O1/powder_space/v2/mathx"
	rr "github.com/Spencer1O1/powder_space/v2/renderer/raylib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type App struct {
	window      *rr.Window
	renderer    *rr.Renderer
	game        *game.Game
	accumulator float64

	mousePos mathx.Vec2
}

func NewApp(window *rr.Window, renderer *rr.Renderer, game *game.Game) *App {
	return &App{
		window:      window,
		renderer:    renderer,
		game:        game,
		accumulator: 0,
	}
}

func (a *App) Run() error {
	for !a.window.ShouldClose() {
		frameDt := float64(rr.GetFrameTime())
		if frameDt > maxFrameDt {
			frameDt = maxFrameDt
		}

		a.handleInput()

		a.accumulator += frameDt

		steps := 0
		for a.accumulator >= fixedPhysicsDt && steps < maxPhysicsStepsPerFrame {
			a.game.Update(fixedPhysicsDt)
			a.accumulator -= fixedPhysicsDt
			steps++
		}

		a.window.Begin()
		a.window.Clear(gfxcolor.Black)

		a.render()

		a.window.End()
	}

	return nil
}

func (a *App) handleInput() {
	mouse := rl.GetMousePosition()
	a.mousePos.X = float64(mouse.X)
	a.mousePos.Y = float64(mouse.Y)

	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		a.game.SetAnchor(a.mousePos)
	}

	if rl.IsMouseButtonReleased(rl.MouseButtonRight) {
		a.game.ResetAnchor()
	}

	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		a.game.SpawnPowder(a.mousePos)
	}
}

func (a *App) render() {
	a.renderer.DrawText(content.TitleString, 20, 20, 32, gfxcolor.White)
	a.renderer.DrawGame(a.game, a.mousePos)
}
