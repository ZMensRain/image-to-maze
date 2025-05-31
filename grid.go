package main

// treat -1 as outside the grid?
type Wall struct {
	cell1 int
	cell2 int
}

func newWall(cell1, cell2 int) Wall {
	return Wall{cell1: cell1, cell2: cell2}
}

type Grid struct {
	// 0:base state can be moved to
	// 1:visited state
	// 2:outside of mask
	cellState map[int]int
	cellWalls map[Wall]bool
}

func (grid *Grid) length() int {
	return len(grid.cellState)
}

func (grid *Grid) updateState(cell, state int) {
	grid.cellState[cell] = state
}

func (grid *Grid) removeWall(wall Wall) {
	delete(grid.cellWalls, wall)
}
func (grid *Grid) addWall(wall Wall) {
	grid.cellWalls[wall] = true
}
func newGrid(width, height uint32) *Grid {
	total := width * height
	g := Grid{cellState: make(map[int]int, total), cellWalls: make(map[Wall]bool)}
	return &g
}
