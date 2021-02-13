package minesweeper

type Cell struct {
	bomb  bool
	label uint
}

type Minesweeper struct {
	board                [][]Cell
	width, height, score uint
}

func New(width uint, height uint) Minesweeper {
	ms := new(Minesweeper)
	ms.score, ms.width, ms.height = 0, width, height
	ms.board = make([][]Cell, width)
	for i := range ms.board {
		board[i] = make([]Cell, height)
	}
}
