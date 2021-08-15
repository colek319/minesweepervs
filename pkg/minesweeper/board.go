package minesweeper

type board struct {
	cells         [][]cell
	width, height int
}

func newBoard(width, height int) board {
	cells := make([][]cell, height)
	for i := range cells {
		cells[i] = make([]cell, width)
		for j := range cells[i] {
			cells[i][j] = newCell(Empty)
		}
	}
	return board{
		cells:  cells,
		width:  width,
		height: height,
	}
}

func (b *board) setCell(loc location, l label) {
	row, col := loc.RowCol()
	b.cells[row][col] = cell{
		label: l,
	}
}

func (b board) getCell(loc location) cell {
	row, col := loc.RowCol()
	return b.cells[row][col]
}

type cell struct {
	label            label
	covered, flagged bool
}

func newCell(label label) cell {
	return cell{
		label:   label,
		covered: true,
		flagged: false,
	}
}

type label int

const (
	Empty label = iota
	OneBomb
	TwoBomb
	ThreeBomb
	FourBomb
	FiveBomb
	SixBomb
	SevenBomb
	EightBomb
	BombLabel
	FlagLabel
)
