package core

// treat -1 as outside the grid?
type Wall struct {
	Cell1 int
	Cell2 int
}

// creates a wall where the smallest index is set as the first.
// This is done so that walls are comparable
func NewWall(cell1, cell2 int) Wall {
	if cell1 < cell2 {
		return Wall{Cell1: cell1, Cell2: cell2}
	} else {
		return Wall{Cell1: cell2, Cell2: cell1}
	}
}
