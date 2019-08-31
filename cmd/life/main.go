package main

import (
	"github.com/davidwmartines/life/internal/game"
	"github.com/davidwmartines/life/internal/seeds"
)

func main() {
	game.Start(seeds.DieHard, 130)
}
