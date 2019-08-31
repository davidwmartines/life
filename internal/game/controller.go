package game

import (
	"time"

	"github.com/davidwmartines/life/internal/seeds"
)

// Start starts a game.
func Start(seed seeds.Seed, generations int, speed int, color string) {

	view := newView(color)
	defer view.close()

	model := newModel(view.height, view.width)
	model.applySeed(seed)

	view.render(model)

	for i := 0; i < generations; i++ {
		model.tick()
		view.render(model)
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}
