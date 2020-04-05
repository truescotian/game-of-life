package main

import (
	"fmt"
	"time"
)

// represets a world
type world struct {
	grid [][]*cell
}

type moveOption string

const (
	UP        moveOption = "UP"
	DOWN      moveOption = "DOWN"
	LEFT      moveOption = "LEFT"
	RIGHT     moveOption = "RIGHT"
	ARRAYSIZE int        = 40
)

// two worlds / states of the game.
var (
	A *world
	B *world
)

// creates a new world with the initialized grid
func newWorld() (w *world) {
	w = &world{
		grid: make([][]*cell, ARRAYSIZE),
	}
	for i := range w.grid {
		w.grid[i] = make([]*cell, ARRAYSIZE)
	}
	return w
}

// copy A into B. We need this to manipulate the next state
// while not losing our current game state.
func copyWorlds() {
	B = &world{
		grid: make([][]*cell, ARRAYSIZE),
	}

	for i := 0; i < ARRAYSIZE; i++ {
		B.grid[i] = make([]*cell, len(A.grid[i]))
		for j := 0; j < ARRAYSIZE; j++ {
			if A.grid[i][j] != nil {
				B.grid[i][j] = &cell{
					alive: A.grid[i][j].alive,
					w:     A.grid[i][j].w,
					x:     A.grid[i][j].x,
					y:     A.grid[i][j].y,
				}
			}
		}
	}
}

// sets up a world with cells
func (w *world) setup() {
	for i := 0; i < ARRAYSIZE; i++ {
		for j := 0; j < ARRAYSIZE; j++ {
			w.grid[i][j] = &cell{
				alive: false,
				w:     w,
				x:     i,
				y:     j,
			}
		}
	}

	w.populate()
}

// populates a world with cells. I just use a line.
func (w *world) populate() {
	xPos := 10
	yPos := 10

	for yPos < 20 {
		c := newCell(w)
		c.alive = true
		c.setPosition(xPos, yPos)
		w.setCell(xPos, yPos, c)
		yPos++
	}

}

// displays the world in console
func (w *world) display() {
	time.Sleep(100 * time.Millisecond)
	clearScreen()
	for x := 0; x < ARRAYSIZE; x++ {
		for y := 0; y < ARRAYSIZE; y++ {
			if w.grid[x][y] == nil {
				fmt.Print(" - ")
			} else if w.grid[x][y].alive {
				fmt.Print(" H ")
			} else if !w.grid[x][y].alive {
				fmt.Print(" - ")
			}
		}
		fmt.Println()
	}
}

// sets the cell in a world
func (w *world) setCell(x, y int, c *cell) {
	w.grid[x][y] = c
}

// iterates through all cells to determine their move for the day.
func (w *world) iterate() {

	copyWorlds() // copies A into B

	for x := 0; x < ARRAYSIZE; x++ {
		for y := 0; y < ARRAYSIZE; y++ {
			if B.grid[x][y] != nil {
				B.grid[x][y].live()
			}
		}
	}

}

func main() {
	clearScreen()
	A = newWorld()
	A.setup()
	A.display()
	day := 0
	for day < 1000 {
		fmt.Println("Day: ", day)
		A.iterate()
		A.display()
		day++
	}

	return
}
