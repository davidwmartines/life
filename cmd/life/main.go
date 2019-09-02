package main

import (
	"flag"

	"github.com/davidwmartines/life/internal/game"
	"github.com/davidwmartines/life/pkg/seeds"
)

func main() {

	generations := flag.Int("gen", 50, "Number of generations.")
	seedName := flag.String("seed", "glider", "Name of a seed to use.")
	speed := flag.Int("speed", 75, "The speed (higher=slower).")
	color := flag.String("color", "green", "The cell color.")
	flag.Parse()

	seed := seeds.New(*seedName)

	game.Start(seed, *generations, *speed, *color)
}
