package inputx

import "github.com/Spencer1O1/powder_space/v2/mathx"

type ContinuousPointerState struct {
	Position      mathx.Vec2
	PrimaryDown   bool
	SecondaryDown bool
}

type DiscretePointerState struct {
	PrimaryPressed    bool
	PrimaryReleased   bool
	SecondaryPressed  bool
	SecondaryReleased bool
}
