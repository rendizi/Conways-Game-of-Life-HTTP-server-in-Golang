package life

import (
	"errors"
	"math/rand"
	"time"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	if height < 0 || width < 0 {
		return nil, errors.New("Height and width can't be negative")
	}
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}
func (w *World) Next(x, y int) bool {
	n := w.neighbors(x, y)       // get the number of living neighbors
	alive := w.Cells[y][x]       // current state of the cell
	if n < 4 && n > 1 && alive { // if there are two or three neighbors and the cell is alive
		return true // then the next state is alive
	}
	if n == 3 && !alive { // if the cell is dead but has three neighbors
		return true // the cell comes to life
	}

	return false // in any other cases - the cell is dead
}
func (w *World) neighbors(x, y int) int {
	k := 0
	neighborOffsets := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, offset := range neighborOffsets {
		nx, ny := x+offset[0], y+offset[1]
		if nx >= 0 && ny >= 0 && nx < w.Height && ny < w.Width && w.Cells[nx][ny] {
			k++
		}
	}
	return k
}

func NextState(oldWorld, newWorld *World) {
	// let's go through all the cells to understand what state they are in
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			// for each cell we get a new state
			newWorld.Cells[i][j] = oldWorld.Next(j, i)
		}
	}
}
func (w *World) RandInit(percentage int) {
	//number of alive cells
	numAlive := percentage * w.Height * w.Width / 100
	// Fill the first cells alive
	w.fillAlive(numAlive)
	// Get random numbers
	r := rand.New(rand.NewSource(time.Now().Unix()))

	// Randomly swap places
	for i := 0; i < w.Height*w.Width; i++ {
		randRowLeft := r.Intn(w.Width)
		randColLeft := r.Intn(w.Height)
		randRowRight := r.Intn(w.Width)
		randColRight := r.Intn(w.Height)

		w.Cells[randRowLeft][randColLeft] = w.Cells[randRowRight][randColRight]
	}
}

func (w *World) fillAlive(num int) {
	aliveCount := 0
	for j, row := range w.Cells {
		for k := range row {
			w.Cells[j][k] = true
			aliveCount++
			if aliveCount == num {

				return
			}
		}
	}
}
