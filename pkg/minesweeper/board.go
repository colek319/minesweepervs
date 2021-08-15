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
			cells[i][j] = newCell()
		}
	}
	return board{
		cells:  cells,
		width:  width,
		height: height,
	}
}

func (b board) uncover(pos Position) {
	cell := b.getCell(pos)
	cell.covered = false
}

func (b board) flag(pos Position) {
	cell := b.getCell(pos)
	cell.flagged = true
}

func (b board) getCell(pos Position) cell {
	return b.cells[pos.row][pos.col]
}

type cell struct {
	covered, flagged bool
}

func newCell() cell {
	return cell{
		covered: true,
		flagged: false,
	}
}
