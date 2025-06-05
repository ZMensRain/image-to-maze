package main

// treat -1 as outside the grid?
type Wall struct {
	cell1 int
	cell2 int
}

// creates a wall where the smallest index is set as the first.
// This is done so that walls are comparable
func newWall(cell1, cell2 int) Wall {
	if cell1 < cell2 {
		return Wall{cell1: cell1, cell2: cell2}
	} else {
		return Wall{cell1: cell2, cell2: cell1}
	}

}

type Grid struct {
	// 0:base state can be moved to
	// 1:visited state
	// 2:outside of mask
	cellState map[int]int
	cellWalls map[Wall]bool
	width     uint32
	height    uint32
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
	g := Grid{cellState: make(map[int]int, total), cellWalls: make(map[Wall]bool), width: width, height: height}
	return &g
}
func (g *Grid) indexToXY(index int) (x, y int) {
	x = index % int(g.width)
	y = (index - x) / int(g.width)
	return x, y
}

// returns the state of an unvisited area or -1 if none is found
func (g *Grid) findUnvisited() int {
	for _, j := range g.cellState {
		if j != 1 {
			return j
		}
	}
	return -1
}
