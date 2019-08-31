package main

import (
	"github.com/davidwmartines/life/internal/game"
)

func main() {

	beehive := [][]int{
		{2, 3},
		{2, 4},
		{3, 2},
		{3, 5},
		{4, 3},
		{4, 4},
	}

	game.Start(beehive)

	// tb.Clear(tb.ColorDefault, tb.ColorDefault)

	// tb.SetCell(1, 1, ' ', tb.ColorCyan, tb.ColorCyan)
	// tb.Flush()

	// for {

	// }
}
