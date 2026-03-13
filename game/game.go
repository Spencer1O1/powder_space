package game

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update(dt float64) {
	// simulation/game update will go here
}
