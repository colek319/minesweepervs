package minesweeper

import (
	"math/rand"
	"time"
)

type minefield struct {
	width, height, mineCount int
	mines                    map[location]interface{}
}

func newMinefield(width, height, mineCount int) minefield {
	mineLocs := make([]bool, width*height)
	for i := 0; i < int(mineCount); i++ {
		mineLocs[i] = true
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(mineLocs), func(i, j int) { mineLocs[i], mineLocs[j] = mineLocs[j], mineLocs[i] })
	mines := map[location]interface{}{}
	for i, b := range mineLocs {
		if b {
			row, col := i/width, i%width
			mines[Location(row, col)] = true
		}
	}
	return minefield{
		width:     width,
		height:    height,
		mineCount: mineCount,
		mines:     mines,
	}
}

func (m *minefield) addMine(loc location) {
	m.mines[loc] = true
	m.mineCount++
}

func (m *minefield) removeMine(loc location) {
	m.mines[loc] = false
	m.mineCount--
}

func (m minefield) hasMine(loc location) bool {
	return m.hasMine(loc)
}

func (m minefield) minesAroundLocation(loc location) int {
	row, col := loc.RowCol()
	locationsToCheck := [8]location{
		Location(row-1, col-1),
		Location(row-1, col),
		Location(row-1, col+1),
		Location(row, col-1),
		Location(row, col+1),
		Location(row+1, col-1),
		Location(row+1, col),
		Location(row-1, col+1),
	}
	count := 0
	for _, location := range locationsToCheck {
		if _, ok := m.mines[location]; ok {
			count += 1
		}
	}
	return count
}
