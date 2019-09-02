package seeds

import "github.com/davidwmartines/life/pkg/model"

// New creates a known Seed by name.
func New(seedName string) (seed model.Seed) {
	switch seedName {
	case "acorn":
		seed = Acorn
	case "blinker":
		seed = Blinker
	case "diehard":
		seed = DieHard
	case "glider":
		seed = Glider
	case "rpentomino":
		seed = RPentomino
	default:
		panic("unknown seed")
	}
	return
}

/* Still lifes */

var Beehive = model.Seed{
	{2, 3},
	{2, 4},
	{3, 2},
	{3, 5},
	{4, 3},
	{4, 4},
}

var Block = model.Seed{
	{2, 3},
	{2, 4},
	{3, 3},
	{3, 4},
}

/* Oscliators */

var Blinker = model.Seed{
	{4, 4},
	{4, 5},
	{4, 6},
}

/* Spaceships */

var Glider = model.Seed{
	{2, 5},
	{3, 6},
	{4, 4},
	{4, 5},
	{4, 6},
}

/* Methuselahs */

var RPentomino = model.Seed{
	{2, 3},
	{2, 4},
	{3, 2},
	{3, 3},
	{4, 3},
}

var Acorn = model.Seed{
	{2, 3},
	{3, 5},
	{4, 2},
	{4, 3},
	{4, 6},
	{4, 7},
	{4, 8},
}

var DieHard = model.Seed{
	{2, 8},
	{3, 2},
	{3, 3},
	{4, 3},
	{4, 7},
	{4, 8},
	{4, 9},
}
