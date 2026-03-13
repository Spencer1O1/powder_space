package content

type MaterialID int

const (
	MaterialDust MaterialID = iota
)

type MaterialDef struct {
	Name    string
	Density float64
}

var Materials = map[MaterialID]MaterialDef{
	MaterialDust: {
		Name:    "Dust",
		Density: 2.387,
	},
}
