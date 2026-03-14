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

func (i *Input) pollContinuous() inputx.ContinuousState {
	mouse := rl.GetMousePosition()

	continuousPointerState := inputx.ContinuousPointerState{
		Position:      mathx.V(float64(mouse.X), float64(mouse.Y)),
		PrimaryDown:   rl.IsMouseButtonDown(rl.MouseLeftButton),
		SecondaryDown: rl.IsMouseButtonDown(rl.MouseRightButton),
	}

	continuousActionState := inputx.ContinuousActionState{}

	return inputx.ContinuousState{
		Pointer: continuousPointerState,
		Action:  continuousActionState,
	}
}

func (i *Input) pollDiscrete() inputx.DiscreteState {
	discretePointerState := inputx.DiscretePointerState{
		PrimaryPressed:    rl.IsMouseButtonPressed(rl.MouseLeftButton),
		PrimaryReleased:   rl.IsMouseButtonReleased(rl.MouseLeftButton),
		SecondaryPressed:  rl.IsMouseButtonPressed(rl.MouseRightButton),
		SecondaryReleased: rl.IsMouseButtonReleased(rl.MouseRightButton),
	}

	discreteActionState := inputx.DiscreteActionState{
		SetSpeed1: rl.IsKeyPressed(rl.KeyOne),
		SetSpeed2: rl.IsKeyPressed(rl.KeyTwo),
		SetSpeed3: rl.IsKeyPressed(rl.KeyThree),
		SetSpeed4: rl.IsKeyPressed(rl.KeyFour),
	}

	return inputx.DiscreteState{
		Pointer: discretePointerState,
		Action:  discreteActionState,
	}
}

func (i *Input) Poll() inputx.State {
	return inputx.State{
		Continuous: i.pollContinuous(),
		Discrete:   i.pollDiscrete(),
	}
}
