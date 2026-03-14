package inputx

import "github.com/Spencer1O1/powder_space/v2/mathx"

type MouseState struct {
	Position      mathx.Vec2
	LeftDown      bool
	LeftPressed   bool
	LeftReleased  bool
	RightDown     bool
	RightPressed  bool
	RightReleased bool
}
