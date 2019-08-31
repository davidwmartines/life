package game

import (
	"time"

	"github.com/davidwmartines/life/internal/grid"
	tb "github.com/nsf/termbox-go"
)

const cellRune = ' '

const speed = 100 * time.Millisecond

const size = 100
const centerFactor = 4

const aliveColor = tb.ColorGreen
const backgroundColor = tb.ColorBlack

var cells map[*grid.Point]*cell
var gr grid.Grid

// Start starts a game.
func Start(seed [][]int, generations int) {

	err := tb.Init()
	if err != nil {
		panic(err)
	}
	defer tb.Close()

	initGrid()

	applySeed(seed)

	render()

	for i := 0; i < generations; i++ {
		tick()
		render()
		time.Sleep(speed)
	}
}

func initGrid() {
	cells = make(map[*grid.Point]*cell)
	gr = grid.Create(size, size)
	for _, row := range gr {
		for _, point := range row {
			cell := cell{point, false, false}
			cells[point] = &cell
		}
	}
}

func applySeed(seed [][]int) {
	for _, pair := range seed {
		seedAlive(pair[0]+(size/centerFactor), pair[1]+(size/centerFactor))
	}
}

func seedAlive(row, col int) {
	cell := cells[gr[row][col]]
	cell.alive = true
	cell.nextState = true
}

func render() {

	tb.Clear(backgroundColor, backgroundColor)

	for _, row := range gr {
		for _, point := range row {
			cell := cells[point]
			var color tb.Attribute
			if cell.alive {
				color = aliveColor
			} else {
				color = backgroundColor
			}
			tb.SetCell(cell.point.Col, cell.point.Row, cellRune, color, color)
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
				cell.nextState = false
			}
		} else {
			if aliveNeighbors == 3 {
				cell.nextState = true
			}
		}
	}

	for _, cell := range cells {
		cell.alive = cell.nextState
	}
}
