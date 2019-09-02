package game

import (
	"time"

	"github.com/davidwmartines/life/pkg/model"
)

// Start starts a game.
func Start(seed model.Seed, generations int, speed int, color string) {

	view := newView(color)
	defer view.close()

	model := model.Create(view.height, view.width)
	model.ApplySeed(seed)

	view.render(model)

	for i := 0; i < generations; i++ {
		model.Tick()
		view.render(model)
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}
