package engine

import (
	"github.com/Spencer1O1/powder_space/v2/game"
	gfxcolor "github.com/Spencer1O1/powder_space/v2/gfx/color"
	"github.com/Spencer1O1/powder_space/v2/inputx"
	"github.com/Spencer1O1/powder_space/v2/mathx"
	rr "github.com/Spencer1O1/powder_space/v2/renderer/raylib"
	"github.com/Spencer1O1/powder_space/v2/scene"
)

type App struct {
	window   *rr.Window
	renderer *rr.Renderer
	input    InputSource
	game     *game.Game

	timeScale        float32
	fixedAccumulator float32
	fixedTick        uint32

	inputState inputx.State

	prevPointerPos mathx.Vec2
	currPointerPos mathx.Vec2
}

func NewApp(window *rr.Window, renderer *rr.Renderer, input InputSource, game *game.Game) *App {
	return &App{
		window:           window,
		renderer:         renderer,
		input:            input,
		game:             game,
		timeScale:        0.25,
		fixedAccumulator: 0,
		fixedTick:        0,
	}
}

func (a *App) Run() error {
	for !a.window.ShouldClose() {
		frameDt := rr.GetFrameTime()
		if frameDt > maxAccumulatedFrameDt {
			frameDt = maxAccumulatedFrameDt
		}

		a.pollInput()
		a.game.Update(frameDt)

		a.fixedAccumulator += frameDt * a.timeScale

		stepsToRun := 0
		tempAccumulator := a.fixedAccumulator
		for tempAccumulator >= fixedDt && stepsToRun < maxFixedUpdatesPerFrame {
			tempAccumulator -= fixedDt
			stepsToRun++
		}

		for step := 0; step < stepsToRun; step++ {
			alpha := float32(1.0)
			if stepsToRun > 0 {
				alpha = float32(step+1) / float32(stepsToRun)
			}

			a.fixedUpdate(fixedDt, alpha)

			a.fixedAccumulator -= fixedDt
			a.fixedTick++
		}

		a.window.Begin()
		a.window.Clear(gfxcolor.Black)
		a.render()
		a.window.End()
	}

	return nil
}

func (a *App) fixedUpdate(dt, alpha float32) {
	if a.inputState.Continuous.Pointer.PrimaryDown {
		if a.fixedTick%spawnFixedTickInterval == 0 {
			pos := a.prevPointerPos.Lerp(a.currPointerPos, alpha)
			a.game.SpawnPowder(pos)
		}
	}
	a.game.FixedUpdate(dt)
}

func (a *App) pollInput() {
	a.inputState = a.input.Poll()

	a.prevPointerPos = a.currPointerPos
	a.currPointerPos = a.inputState.Continuous.Pointer.Position

	if a.inputState.Discrete.Pointer.SecondaryPressed {
		a.game.SetAnchor(a.currPointerPos)
	}
	if a.inputState.Discrete.Pointer.SecondaryReleased {
		a.game.ResetAnchor()
	}

	if a.inputState.Discrete.Action.SetSpeed1 {
		a.timeScale = 0.25
	}
	if a.inputState.Discrete.Action.SetSpeed2 {
		a.timeScale = 0.50
	}
	if a.inputState.Discrete.Action.SetSpeed3 {
		a.timeScale = 1.00
	}
	if a.inputState.Discrete.Action.SetSpeed4 {
		a.timeScale = 2.00
	}
}

func (a *App) render() {
	scene.DrawHud(
		a.renderer,
		a.game,
		scene.HudParams{
			TimeScale: a.timeScale,
		},
	)
	scene.DrawGame(
		a.renderer,
		a.game,
		scene.GameViewParams{
			PointerPos: a.inputState.Continuous.Pointer.Position,
		},
	)
}
