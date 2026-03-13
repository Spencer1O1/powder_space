package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1280, 720, "PowderSpace")
	defer rl.CloseWindow()

	rl.ShowCursor()
	rl.EnableCursor()
	rl.SetMouseCursor(rl.MouseCursorDefault)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawText("PowderSpace | Powder Game but in SPACE", 20, 20, 32, rl.White)
		rl.EndDrawing()
	}
}	