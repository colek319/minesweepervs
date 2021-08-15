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
		ms.uncover(move)
	case FlagOp:
		ms.flag(move.loc)
	default:
		errorMessage := fmt.Sprintf("invalid operation: %d", int(move.op))
		return errors.New(errorMessage)
	}
	return nil
}

func (ms *Minesweeper) uncover(move Move) (err error) {
	if ms.mineField.hasMine(move.loc) {
		ms.gameOver = true
		return
	}
	var lab label
	numBombs := ms.mineField.minesAroundLocation(move.loc)
	switch numBombs {
	case 0:
		lab = Empty
		ms.floodUncover(move.loc)
	case 1:
		lab = OneBomb
	case 2:
		lab = TwoBomb
	case 3:
		lab = ThreeBomb
	case 4:
		lab = FourBomb
	case 5:
		lab = FiveBomb
	case 6:
		lab = SixBomb
	case 7:
		lab = SevenBomb
	case 8:
		lab = EightBomb
	default:
		ms.gameBoard.setCell(move.loc, Empty)
		return errors.New("error in uncover: unknown bomb value")
	}
	ms.gameBoard.setCell()
	return
}

func (ms *Minesweeper) floodUncover(loc location) {
	seenLocations := map[location]bool{}
	queue := make(chan location, ms.width*ms.height) // size of board
	queue <- loc
	var currLoc location
	for len(queue) != 0 {
		currLoc = <-queue
		row, col := currLoc.RowCol()
		if row < 0 || col < 0 || row >= ms.height || col >= ms.width {
			continue
		} else if _, ok := seenLocations[currLoc]; !ok {
			cell := ms.gameBoard.getCell(currLoc)
			if cell.covered {
				cell.covered = false
			}

			seenLocations[currLoc] = true
			queue <- currLoc.Up()
			queue <- currLoc.Down()
			queue <- currLoc.Left()
			queue <- currLoc.Right()
		}
	}
}

func (ms *Minesweeper) flag(loc location) {
	ms.gameBoard.getCell()
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

func (l location) Up() location {
	row, col := l.RowCol()
	return Location(row+1, col)
}

func (l location) Down() location {
	row, col := l.RowCol()
	return Location(row-1, col)
}

func (l location) Left() location {
	row, col := l.RowCol()
	return Location(row, col-1)
}

func (l location) Right() location {
	row, col := l.RowCol()
	return Location(row, col+1)
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
