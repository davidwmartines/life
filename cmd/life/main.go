package main

import (
	"flag"

	"github.com/nsf/termbox-go"

	"github.com/davidwmartines/life/internal/game"
	"github.com/davidwmartines/life/internal/seeds"
)

func main() {

	generations := flag.Int("gen", 100, "Number of generations.")
	seedName := flag.String("seed", "glider", "Name of a seed to use")
	speed := flag.Int("speed", 100, "The speed (higher=slower).")
	colorName := flag.String("color", "green", "The cell color.")
	flag.Parse()

	seed := seeds.New(*seedName)
	color := parseColor(*colorName)

	game.Start(seed, *generations, *speed, color)
}

func parseColor(colorName string) (color termbox.Attribute) {
	switch colorName {
	case "green":
		color = termbox.ColorGreen
	case "blue":
		color = termbox.ColorBlue
	case "red":
		color = termbox.ColorRed
	case "cyan":
		color = termbox.ColorCyan
	case "magenta":
		color = termbox.ColorMagenta
	case "yellow":
		color = termbox.ColorYellow
	case "white":
		color = termbox.ColorWhite
	case "black":
		color = termbox.ColorBlack
	default:
		panic("invalid color")
	}
	return
}
