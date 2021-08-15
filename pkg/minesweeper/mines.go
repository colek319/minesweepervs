package minesweeper

import (
	"math/rand"
	"time"
)

type minefield struct {
	width, height, mineCount int
	mines                    map[string]bool
}

func newMinefield(width, height, mineCount int) minefield {
	mineLocations := make([]bool, width*height)
	for i := 0; i < int(mineCount); i++ {
		mineLocations[i] = true
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(mineLocations), func(i, j int) { mineLocations[i], mineLocations[j] = mineLocations[j], mineLocations[i] })
	mines := map[string]bool{}
	for i, b := range mineLocations {
		if b {
			pos := Position{i / width, i % width}
			mines[pos.stringify()] = true
		}
	}
	return minefield{
		width:     width,
		height:    height,
		mineCount: mineCount,
		mines:     mines,
	}
}

func (m *minefield) addMine(pos Position) {
	m.mines[pos.stringify()] = true
	m.mineCount++
}

func (m *minefield) removeMine(pos Position) {
	m.mines[pos.stringify()] = false
	m.mineCount--
}

func (m minefield) hasMine(pos Position) bool {
	return m.mines[pos.stringify()]
}

func (m minefield) minesAroundLocation(pos Position) int {
	row, col := pos.row, pos.col
	positionsToCheck := [8]Position{
		{row - 1, col - 1},
		{row - 1, col},
		{row - 1, col + 1},
		{row, col - 1},
		{row, col + 1},
		{row + 1, col - 1},
		{row + 1, col},
		{row - 1, col + 1},
	}
	count := 0
	for _, pos := range positionsToCheck {
		if hasMine, ok := m.mines[pos.stringify()]; ok && hasMine {
			count += 1
		}
	}
	return count
}
