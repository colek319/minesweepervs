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
	if !ms.mineField.hasMine(move.pos) {
		return
	}

	ms.mineField.removeMine(move.pos)

	for i := 0; i < ms.width; i++ {
		for j := 0; j < ms.height; j++ {
			pos := move.pos
			if (i != pos.row || j != pos.col) && !ms.mineField.hasMine(pos) {
				ms.mineField.removeMine(move.pos)
				return
			}
		}
	}
}

func (ms Minesweeper) Move(move Move) error {
	switch move.op {
	case UncoverOp:
		ms.uncover(move)
	case FlagOp:
		ms.flag(move.pos)
	default:
		errorMessage := fmt.Sprintf("invalid operation: %d", int(move.op))
		return errors.New(errorMessage)
	}
	return nil
}

func (ms *Minesweeper) uncover(move Move) (err error) {
	if ms.mineField.hasMine(move.pos) {
		ms.gameOver = true
		return
	}
	numBombs := ms.mineField.minesAroundLocation(move.pos)
	if numBombs == 0 {
		ms.floodUncover(move.pos)
	}
	ms.gameBoard.uncover(move.pos)
	return
}

func (ms *Minesweeper) flag(pos Position) {
	ms.gameBoard.flag(pos)
}

func (ms *Minesweeper) floodUncover(pos Position) {
	seenLocations := map[string]bool{}
	queue := make(chan Position, ms.width*ms.height) // size of board
	// initialize queue
	queue <- pos
	var currPos Position
	for len(queue) != 0 {
		currPos = <-queue
		if currPos.row < 0 || currPos.col < 0 || currPos.row >= ms.height || currPos.col >= ms.width {
			continue
		} else if _, ok := seenLocations[currPos.stringify()]; !ok {
			ms.gameBoard.uncover(currPos)
			seenLocations[currPos.stringify()] = true
			row, col := currPos.row, currPos.col
			queue <- Position{row + 1, col}
			queue <- Position{row - 1, col}
			queue <- Position{row, col - 1}
			queue <- Position{row, col + 1}
		}
	}
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
	pos Position
	op  Operation
}

type Position struct {
	row, col int
}

func positionFromString(pos string) Position {
	var row, col int
	fmt.Sscanf(pos, "(%d,%d)", &row, &col)
	return Position{row, col}
}

func (p Position) stringify() string {
	return fmt.Sprintf("(%d,%d)", p.row, p.col)
}

type Operation int

const (
	UncoverOp Operation = iota
	FlagOp
)

func (ms Minesweeper) NewMove(op Operation, row, col int) (Move, error) {
	if row < 0 || row >= ms.height || col < 0 || col >= ms.width {
		errorMessage := fmt.Sprintf("invalid move: row %d or col %d out of range", row, col)
		return Move{}, errors.New(errorMessage)
	}
	return Move{
		pos: Position{row, col},
		op:  op,
	}, nil
}
