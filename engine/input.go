package engine

import "github.com/Spencer1O1/powder_space/v2/inputx"

type InputSource interface {
	Poll() inputx.State
}
