package core

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
)

type Grid struct {
	Width  uint32
	Height uint32
	// if value is true wall is open else closed
	OpenWalls map[Wall]bool
	// 0:base state can be moved to
	// 1:visited state
	// 2:outside of mask
	CellState map[int]int
}

func NewGrid(width, height uint32) *Grid {
	total := width * height
	g := Grid{CellState: make(map[int]int, total), OpenWalls: make(map[Wall]bool), Width: width, Height: height}
	return &g
}

func (g *Grid) UpdateState(cell, state int) {
	g.CellState[cell] = state
}

func (g *Grid) RemoveWall(wall Wall) {
	g.OpenWalls[wall] = true
}

// returns the state of an unvisited area or -1 if none is found
func (g *Grid) FindUnvisited() int {
	for _, j := range g.CellState {
		if j != 1 {
			return j
		}
	}
	return -1
}

func (g *Grid) IndexToXY(index int) (x, y int) {
	x = index % int(g.Width)
	y = (index - x) / int(g.Width)
	return x, y
}

// -------------------------------------------------------
// returns -1 if no start can be found
func (g *Grid) findStart(state int) (index int, err error) {
	for i, cellState := range g.CellState {
		//checks if the cell is in the mask
		if cellState == state {
			return i, nil
		}
	}
	return -1, errors.New("no mask found")
}

func (g *Grid) GenerateMaze(state int) {
	start, err := g.findStart(state)
	if err != nil {
		return
	}
	g.iterateGeneration(start, state)
}

func (g *Grid) getNeighbors(from int, state int) []int {
	x := from % int(g.Width)
	y := (from - x) / int(g.Width)

	neighbor := []int{}

	// left neighbor
	var index = from - 1
	if x > 0 && g.CellState[index] == state {
		neighbor = append(neighbor, index)
	}

	// top neighbor
	index = x + (y-1)*int(g.Width)
	if y > 0 && g.CellState[index] == state {
		neighbor = append(neighbor, index)
	}

	// right neighbor
	index = from + 1
	if x < int(g.Width)-1 && g.CellState[index] == state {
		neighbor = append(neighbor, index)
	}

	// bottom neighbor
	index = x + (y+1)*int(g.Width)
	if y < int(g.Height)-1 && g.CellState[index] == state {
		neighbor = append(neighbor, index)
	}

	return neighbor
}

func (g *Grid) iterateGeneration(from, state int) {
	re := func() []int { return g.getNeighbors(from, state) }
	g.UpdateState(from, 1)
	for neighbors := re(); len(neighbors) != 0; neighbors = re() {
		i := rand.Intn(len(neighbors))
		g.RemoveWall(NewWall(from, neighbors[i]))
		g.iterateGeneration(neighbors[i], state)
	}
}

func (g *Grid) RenderWalls(background, foreground color.RGBA) *image.RGBA {
	// creates the image and sets up some useful variables
	width := int(g.Width)*2 + 1
	height := int(g.Height)*2 + 1
	area := image.Rect(0, 0, width, height)
	img := image.NewRGBA(area)

	// draws a border around the image
	draw.Draw(img, img.Bounds(), image.NewUniform(foreground), image.Pt(0, 0), draw.Src)
	for i, exists := range g.OpenWalls {
		x2, y2 := g.IndexToXY(i.Cell2)
		x1, y1 := g.IndexToXY(i.Cell1)
		// filters out closed walls
		if !exists {
			continue
		}
		transX1 := 2*x1 + 1
		transY1 := 2*y1 + 1
		transX2 := 2*x2 + 1
		transY2 := 2*y2 + 1
		if transY1 == transY2 {
			img.SetRGBA(transX1, transY1, background)
			img.SetRGBA(transX2, transY1, background)
			img.SetRGBA(((transX1 + transX2) / 2), transY1, background)

		} else if transX1 == transX2 {
			img.SetRGBA(transX1, transY1, background)
			img.SetRGBA(transX1, transY2, background)
			img.SetRGBA(transX1, ((transY1 + transY2) / 2), background)

		} else {
			fmt.Println("invalid wall", i)
			continue
		}

	}

	//Writes the image to output.png
	return img
}

func PixelToState(pixel color.Color) int {
	r, g, b, _ := pixel.RGBA()
	// fmt.Println(r, g, b, _a)
	if r >= 32_766 && g >= 32_766 && b >= 32_766 {
		return 2
	}
	return 0
}

func GridFromImage(img image.Image) *Grid {
	size := img.Bounds()
	grid := NewGrid(uint32(size.Dx()), uint32(size.Dy()))
	width, height := size.Dx(), size.Dy()

	for y := range height {
		for x := range width {
			index := x + (y * size.Dx())
			// sets state
			grid.UpdateState(index, PixelToState(img.At(x, y)))
		}
	}

	return grid
}
