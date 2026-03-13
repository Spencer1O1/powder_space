package raylib

import (
	gfxcolor "github.com/Spencer1O1/powder_space/v2/gfx/color"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Window struct {
	width  int32
	height int32
	title  string
}

func NewWindow(w, h int32, title string) *Window {
	rl.InitWindow(1280, 720, "PowderSpace")

	rl.ShowCursor()
	rl.EnableCursor()
	rl.SetMouseCursor(rl.MouseCursorDefault)

	rl.SetTargetFPS(60)

	return &Window{
		width:  w,
		height: h,
		title:  title,
	}
}

func (w *Window) Close() {
	rl.CloseWindow()
}

func (w *Window) ShouldClose() bool {
	return rl.WindowShouldClose()
}

func (w *Window) Begin() {
	rl.BeginDrawing()
}

func (w *Window) End() {
	rl.EndDrawing()
}

func (w *Window) Clear(c gfxcolor.Color) {
	rl.ClearBackground(toRLColor(c))
}

func (w *Window) Width() int32 {
	return w.width
}

func (w *Window) Height() int32 {
	return w.height
}
