package game

import "github.com/davidwmartines/life/internal/grid"

type cell struct {
	point *grid.Point
	alive bool
}
