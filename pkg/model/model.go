package model

import (
	"github.com/davidwmartines/life/internal/grid"
)

//Seed seed pattern.
type Seed [][]int

//Model life model.
type Model struct {
	Height int
	Width  int
	Cells  map[*grid.Point]*Cell
	Grid   grid.Grid
}

//Cell cell in a model.
type Cell struct {
	Point     *grid.Point
	Alive     bool
	nextState bool
}

const centerFactor = 3

// Create creates a new Model.
func Create(height, width int) *Model {
	model := Model{}
	model.Height = height
	model.Width = width
	model.Cells = make(map[*grid.Point]*Cell)
	model.Grid = grid.Create(height, width)
	for _, row := range model.Grid {
		for _, point := range row {
			cell := Cell{point, false, false}
			model.Cells[point] = &cell
		}
	}
	return &model
}

// ApplySeed applies a seed.
func (model *Model) ApplySeed(seed Seed) {
	for _, pair := range seed {
		model.seedAlive(pair[0]+(model.Height/centerFactor), pair[1]+(model.Width/centerFactor))
	}
}

func (model *Model) seedAlive(row, col int) {
	cell := model.Cells[model.Grid[row][col]]
	cell.Alive = true
	cell.nextState = true
}

// Tick advances the model one generation.
func (model *Model) Tick() {
	for _, cell := range model.Cells {
		neighbors := model.Grid.Neighbors(cell.Point)
		aliveNeighbors := 0
		for _, n := range neighbors {
			if model.Cells[n].Alive {
				aliveNeighbors++
			}
		}
		if cell.Alive {
			if aliveNeighbors < 2 || aliveNeighbors > 3 {
				cell.nextState = false
			}
		} else {
			if aliveNeighbors == 3 {
				cell.nextState = true
			}
		}
	}

	for _, cell := range model.Cells {
		cell.Alive = cell.nextState
	}
}
