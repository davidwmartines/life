package game

import (
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

func parseColor(colorName string) (color tb.Attribute) {
	switch colorName {
	case "green":
		color = tb.ColorGreen
	case "blue":
		color = tb.ColorBlue
	case "red":
		color = tb.ColorRed
	case "cyan":
		color = tb.ColorCyan
	case "magenta":
		color = tb.ColorMagenta
	case "yellow":
		color = tb.ColorYellow
	case "white":
		color = tb.ColorWhite
	case "black":
		color = tb.ColorBlack
	default:
		panic("invalid color")
	}
	return
}
