package game

import (
	"time"

	"github.com/davidwmartines/life/internal/grid"
	tb "github.com/nsf/termbox-go"
)

const speed = 300 * time.Millisecond

var cells map[*grid.Point]*cell
var gr grid.Grid

// Start starts a game.
func Start(seed [][]int) {

	err := tb.Init()
	if err != nil {
		panic(err)
	}
	defer tb.Close()

	initGrid()

	applySeed(seed)

	render()
	for {
		tick()
		render()
	}
}

func initGrid() {
	cells = make(map[*grid.Point]*cell)
	gr = grid.Create(100, 100)
	for _, row := range gr {
		for _, point := range row {
			cell := cell{point, false}
			cells[point] = &cell
		}
	}
}

func applySeed(seed [][]int) {

	for _, pair := range seed {
		seedAlive(pair[0], pair[1])
	}
	// seedAlive(2, 3)
	// seedAlive(2, 4)
	// seedAlive(3, 2)
	// seedAlive(3, 5)
	// seedAlive(4, 3)
	// seedAlive(4, 4)

}

func seedAlive(row, col int) {
	cells[gr[row][col]].alive = true
}

func render() {

	tb.Clear(tb.ColorDefault, tb.ColorDefault)

	for _, row := range gr {
		for _, point := range row {
			cell := cells[point]
			var color tb.Attribute
			if cell.alive {
				color = tb.ColorGreen
			} else {
				color = tb.ColorDefault
			}
			tb.SetCell(cell.point.Col, cell.point.Row, ' ', color, color)
		}
	}

	tb.Flush()
}

func tick() {
	for _, cell := range cells {
		neighbors := gr.Neighbors(cell.point)
		aliveNeighbors := 0
		for _, n := range neighbors {
			if cells[n].alive {
				aliveNeighbors++
			}
		}
		if cell.alive {
			if aliveNeighbors < 2 || aliveNeighbors > 3 {
				cell.alive = false
			}
		} else {
			if aliveNeighbors == 3 {
				cell.alive = true
			}
		}
	}
	time.Sleep(speed)
}
