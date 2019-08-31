package game

import (
	"github.com/davidwmartines/life/internal/grid"
	"github.com/davidwmartines/life/internal/seeds"
)

type model struct {
	height int
	width  int
	cells  map[*grid.Point]*cell
	grid   grid.Grid
}

type cell struct {
	point     *grid.Point
	alive     bool
	nextState bool
}

const centerFactor = 3

func newModel(height, width int) *model {
	model := model{}
	model.height = height
	model.width = width
	model.cells = make(map[*grid.Point]*cell)
	model.grid = grid.Create(height, width)
	for _, row := range model.grid {
		for _, point := range row {
			cell := cell{point, false, false}
			model.cells[point] = &cell
		}
	}
	return &model
}

func (model *model) applySeed(seed seeds.Seed) {
	for _, pair := range seed {
		model.seedAlive(pair[0]+(model.height/centerFactor), pair[1]+(model.width/centerFactor))
	}
}

func (model *model) seedAlive(row, col int) {
	cell := model.cells[model.grid[row][col]]
	cell.alive = true
	cell.nextState = true
}

func (model *model) tick() {
	for _, cell := range model.cells {
		neighbors := model.grid.Neighbors(cell.point)
		aliveNeighbors := 0
		for _, n := range neighbors {
			if model.cells[n].alive {
				aliveNeighbors++
			}
		}
		if cell.alive {
			if aliveNeighbors < 2 || aliveNeighbors > 3 {
				cell.nextState = false
			}
		} else {
			if aliveNeighbors == 3 {
				cell.nextState = true
			}
		}
	}

	for _, cell := range model.cells {
		cell.alive = cell.nextState
	}
}
