package game

import (
	"github.com/nsf/termbox-go"
	tb "github.com/nsf/termbox-go"
)

var runes = []rune{
	'█', '◆', '✦', '★', '✱', '✚', '·',
}

type view struct {
	height     int
	width      int
	aliveColor tb.Attribute
	cellRune   rune
}

const backgroundColor = tb.ColorBlack

func newView(colorName string) *view {
	err := tb.Init()
	if err != nil {
		panic(err)
	}

	view := view{}
	view.aliveColor = parseColor(colorName)

	//maybe make this a parameter
	view.cellRune = runes[0]

	view.width, view.height = tb.Size()

	return &view
}

func (view *view) close() {
	tb.Close()
}

func (view *view) render(model *model) {

	tb.Clear(backgroundColor, backgroundColor)

	for _, row := range model.grid {
		for _, point := range row {
			cell := model.cells[point]
			var color tb.Attribute
			if cell.alive {
				color = view.aliveColor
			} else {
				color = backgroundColor
			}
			tb.SetCell(cell.point.Col, cell.point.Row, view.cellRune, color, backgroundColor)
		}
	}

	tb.Flush()
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
