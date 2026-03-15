package content

import "github.com/Spencer1O1/powder_space/v2/mathx/geo"

type MaterialID int
type MaterialMap map[MaterialID]MaterialDef
type CompositionMap map[MaterialID]float32

const (
	MaterialDust MaterialID = iota
)

type MaterialDef struct {
	Name    string
	Density float32
}

var Materials = MaterialMap{
	MaterialDust: {
		Name:    "Dust",
		Density: 2.387,
	},
}

func (c CompositionMap) GetSphericalDerivedValues() (
	mass,
	radius,
	volume,
	density float32,
) {
	totalMass := float32(0.0)
	totalVolume := float32(0.0)

	for materialID, materialMass := range c {
		if materialMass <= 0 {
			continue
		}

		mat := Materials[materialID]
		density := mat.Density

		totalMass += materialMass
		totalVolume += materialMass / density
	}

	if totalVolume <= 0 {
		return 0.0, 0.0, 0.0, 1.0
	}

	d := totalMass / totalVolume
	r := geo.SphericalRadiusFromVolume(totalVolume)

	return totalMass, r, totalVolume, d
}
