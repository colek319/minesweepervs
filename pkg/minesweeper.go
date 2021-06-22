package minesweeper

import (
	"errors"
	"fmt"
)

type Minesweeper struct {
	gameBoard                board
	mineField                minefield
	width, height, mineCount int
	gameOver                 bool
}

func New(width, height, mineCount int) Minesweeper {
	return Minesweeper{
		width:     width,
		height:    height,
		mineCount: mineCount,
		gameBoard: newBoard(width, height),
		mineField: newMinefield(width, height, mineCount),
	}
}

func (ms *Minesweeper) FirstMove(move Move) {
	// TODO: check row, col bounds
	if !ms.mineField.hasMine(move.loc) {
		return
	}

	ms.mineField.removeMine(move.loc)

	for i := 0; i < ms.width; i++ {
		for j := 0; j < ms.height; j++ {
			row, col := move.loc.RowCol()
			if (i != row || j != col) && !ms.mineField.hasMine(move.loc) {
				ms.mineField.removeMine(move.loc)
				return
			}
		}
	}
}

func (ms Minesweeper) Move(move Move) error {
	switch move.op {
	case UncoverOp:
		ms.uncover(move.loc)
	case FlagOp:
		ms.flag(move.loc)
	default:
		errorMessage := fmt.Sprintf("invalid operation: %d", int(move.op))
		return errors.New(errorMessage)
	}
	return nil
}

func (ms *Minesweeper) uncover(loc location) {
	if ms.mineField.hasMine(loc) {
		ms.gameOver = true
		return
	}
	numBombs := ms.mineField.minesAroundLocation(loc)
	ms.board.SetLabel(numBombs)
}

func (ms *Minesweeper) floodFill(loc location) {

}

func (ms *Minesweeper) flag(loc location) {

}

func (ms Minesweeper) CalculateScore() int {
	return 0
}

func (ms Minesweeper) get3BV() int {
	score3BV := 0
	_ = score3BV
	return score3BV
}

type Move struct {
	loc location
	op  Operation
}

type Operation int

const (
	UncoverOp Operation = iota
	FlagOp
)

func (ms Minesweeper) NewMove(op Operation, row, col int) (Move, error) {
	if row < 0 || row >= ms.height || col < 0 || col >= ms.width {
		errorMessage := fmt.Sprint("invalid move: row %d or col %d out of range", row, col)
		return Move{}, errors.New(errorMessage)
	}
	return Move{
		loc: Location(row, col),
		op:  op,
	}, nil
}

type location string

func Location(row, col int) location {
	return location(fmt.Sprintf("(%d,%d)", row, col))
}

func (l location) String() string {
	return string(l)
}

func (l location) RowCol() (row, col int) {
	fmt.Sscanf(l.String(), "(%d,%d)", &row, &col)
	return
}

func (l location) Row() (row int) {
	row, _ = l.RowCol()
	return
}

func (l location) Cow() (col int) {
	_, col = l.RowCol()
	return
}
