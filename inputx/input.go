package inputx

type ContinuousState struct {
	Pointer ContinuousPointerState
	Action  ContinuousActionState
}

type DiscreteState struct {
	Pointer DiscretePointerState
	Action  DiscreteActionState
}

type State struct {
	Continuous ContinuousState
	Discrete   DiscreteState
}
