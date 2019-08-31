package game

import (
	"time"

	"github.com/davidwmartines/life/internal/grid"
	"github.com/davidwmartines/life/internal/seeds"
	tb "github.com/nsf/termbox-go"
)

var runes = []rune{
	'█', '◆', '✦', '★', '✱', '✚', '·',
}

const backgroundColor = tb.ColorBlack
const centerFactor = 3

var width, height int

var cells map[*grid.Point]*cell
var gr grid.Grid

// Start starts a game.
func Start(seed seeds.Seed, generations int, speed int, color tb.Attribute) {

	err := tb.Init()
	if err != nil {
		panic(err)
	}
	defer tb.Close()

	width, height = tb.Size()

	initGrid()

	applySeed(seed)

	//maybe make this a flag
	cellRune := runes[0]

	render(color, cellRune)

	for i := 0; i < generations; i++ {
		tick()
		render(color, cellRune)
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}

func initGrid() {

	cells = make(map[*grid.Point]*cell)
	gr = grid.Create(height, width)
	for _, row := range gr {
		for _, point := range row {
			cell := cell{point, false, false}
			cells[point] = &cell
		}
	}
}

func applySeed(seed seeds.Seed) {
	for _, pair := range seed {
		seedAlive(pair[0]+(height/centerFactor), pair[1]+(width/centerFactor))
	}
}

func seedAlive(row, col int) {
	cell := cells[gr[row][col]]
	cell.alive = true
	cell.nextState = true
}

func render(aliveColor tb.Attribute, cellRune rune) {

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
			tb.SetCell(cell.point.Col, cell.point.Row, cellRune, color, backgroundColor)
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
