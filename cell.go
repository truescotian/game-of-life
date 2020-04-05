package main

// represents a cell
type cell struct {
	alive bool
	w     *world
	x     int
	y     int
}

// creates a new cell to be born into a world
func newCell(w *world) (c *cell) {
	return &cell{
		alive: false,
		w:     w,
	}
}

// sets the position of a cell
func (c *cell) setPosition(x, y int) {
	c.x = x
	c.y = y
}

// determines what move the cell will take today
func (c *cell) live() {
	numAliveNeighbors := 0
	var foundCell *cell

	// top left
	if c.x-1 >= 0 && c.y-1 >= 0 {
		foundCell = B.grid[c.x-1][c.y-1]
		if foundCell != nil && foundCell.alive {
			numAliveNeighbors++
		}
	}
	// above
	if c.x-1 >= 0 {
		foundCell = B.grid[c.x-1][c.y]
		if foundCell != nil && foundCell.alive {
			numAliveNeighbors++
		}
	}
	// top right
	if c.x-1 >= 0 && c.y+1 < ARRAYSIZE {
		foundCell = B.grid[c.x-1][c.y+1]
		if foundCell != nil && foundCell.alive {
			numAliveNeighbors++
		}
	}
	// right
	if c.y+1 < ARRAYSIZE {
		foundCell = B.grid[c.x][c.y+1]
		if foundCell != nil && foundCell.alive {
			numAliveNeighbors++
		}
	}
	// bottom right
	if c.x+1 < ARRAYSIZE && c.y+1 < ARRAYSIZE {
		foundCell = B.grid[c.x+1][c.y+1]
		if foundCell != nil && foundCell.alive {
			numAliveNeighbors++
		}
	}
	// bottom
	if c.x+1 < ARRAYSIZE {
		foundCell = B.grid[c.x+1][c.y]
		if foundCell != nil && foundCell.alive {
			numAliveNeighbors++
		}
	}
	// bottom left
	if c.x+1 < ARRAYSIZE && c.y-1 >= 0 {
		foundCell = B.grid[c.x+1][c.y-1]
		if foundCell != nil && foundCell.alive {
			numAliveNeighbors++
		}
	}
	// left
	if c.y-1 >= 0 {
		foundCell = B.grid[c.x][c.y-1]
		if foundCell != nil && foundCell.alive {
			numAliveNeighbors++
		}
	}
	if c.alive {
		if numAliveNeighbors < 2 || numAliveNeighbors > 3 {
			A.grid[c.x][c.y].alive = false
		}
	} else if !c.alive && numAliveNeighbors == 3 {
		A.grid[c.x][c.y].alive = true
	}
}
