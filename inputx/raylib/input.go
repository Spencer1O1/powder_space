package raylib

import (
	"github.com/Spencer1O1/powder_space/v2/inputx"
	"github.com/Spencer1O1/powder_space/v2/mathx"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Input struct{}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) PollMouse() inputx.MouseState {
	mouse := rl.GetMousePosition()

	return inputx.MouseState{
		Position:      mathx.V(float64(mouse.X), float64(mouse.Y)),
		LeftDown:      rl.IsMouseButtonDown(rl.MouseLeftButton),
		LeftPressed:   rl.IsMouseButtonPressed(rl.MouseLeftButton),
		LeftReleased:  rl.IsMouseButtonReleased(rl.MouseLeftButton),
		RightDown:     rl.IsMouseButtonDown(rl.MouseRightButton),
		RightPressed:  rl.IsMouseButtonPressed(rl.MouseRightButton),
		RightReleased: rl.IsMouseButtonReleased(rl.MouseRightButton),
	}
}
