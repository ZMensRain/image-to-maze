package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"math/rand"
	"os"
)

// Opens image and makes it usable
func DecodeImage(filename string) (image.Image, string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return nil, "", err
	}

	return img, format, nil
}

func pixelToState(r, g, b, _a uint32) int {
	// fmt.Println(r, g, b, _a)
	if r == 65535 && g == 65535 && b == 65535 {
		return 2
	}
	return 0
}

func GridFromImage(img image.Image) *Grid {
	size := img.Bounds()
	grid := newGrid(uint32(size.Dx()), uint32(size.Dy()))
	for y := 0; y < size.Dy(); y++ {
		for x := 0; x < size.Dx(); x++ {
			index := x + (y * size.Dx())
			// sets state
			r, g, b, a := img.At(x, y).RGBA()
			grid.updateState(index, pixelToState(r, g, b, a))

			// sets walls
			if y < size.Dy()-1 {
				//May have issue here where dy is 1 bigger than it should be
				grid.addWall(newWall(index, index+size.Dy()))
			}

			if x < size.Dx()-1 {
				//May have issue here where dy is 1 bigger than it should be
				grid.addWall(newWall(index, index+1))
			}
		}
	}

	return grid
}

// returns -1 if no start can be found
func (g *Grid) findStart(state int) (index int, err error) {
	for i, cellState := range g.cellState {
		//checks if the cell is in the mask
		if cellState == state {
			return i, nil
		}
	}
	return -1, errors.New("no mask found")
}

func (g *Grid) generateMaze(state int) {
	start, err := g.findStart(state)
	if err != nil {
		return
	}
	g.iterateGeneration(start, state)
}

func (g *Grid) getNeighbors(from int, state int) []int {
	x := from % int(g.width)
	y := (from - x) / int(g.width)

	neighbor := []int{}

	// left neighbor
	var index = from - 1
	if x > 0 && g.cellState[index] == state {
		neighbor = append(neighbor, index)
	}

	// top neighbor
	index = x + (y-1)*int(g.width)
	if y > 0 && g.cellState[index] == state {
		neighbor = append(neighbor, index)
	}

	// right neighbor
	index = from + 1
	if x < int(g.width)-1 && g.cellState[index] == state {
		neighbor = append(neighbor, index)
	}

	// bottom neighbor
	index = x + (y+1)*int(g.width)
	if y < int(g.height)-1 && g.cellState[index] == state {
		neighbor = append(neighbor, index)
	}

	return neighbor
}

func (g *Grid) iterateGeneration(from, state int) {
	re := func() []int { return g.getNeighbors(from, state) }
	g.updateState(from, 1)
	for neighbors := re(); len(neighbors) != 0; neighbors = re() {
		i := rand.Intn(len(neighbors))
		g.removeWall(newWall(from, neighbors[i]))
		g.iterateGeneration(neighbors[i], state)
	}
}

func (g *Grid) renderWalls(path string, background color.RGBA, foreground color.RGBA) {
	// creates the image and sets up some useful variables
	width := int(g.width)*2 + 1
	height := int(g.height)*2 + 1
	area := image.Rect(0, 0, width, height)
	img := image.NewRGBA(area)

	// draws a border around the image
	draw.Draw(img, img.Bounds(), image.NewUniform(foreground), image.Pt(0, 0), draw.Src)
	draw.Draw(img, image.Rect(1, 1, width-1, height-1), image.NewUniform(background), image.Pt(1, 1), draw.Src)

	for i, exists := range g.cellWalls {
		x1, y1 := g.indexToXY(i.cell1)
		x2, y2 := g.indexToXY(i.cell2)
		if !exists {
			continue
		}
		transX1 := 2*x1 + 1
		transY1 := 2*y1 + 1
		transX2 := 2*x2 + 1
		transY2 := 2*y2 + 1
		if transY1 == transY2 {
			img.SetRGBA((transX1+transX2)/2, transY1, foreground)
			img.SetRGBA(((transX1 + transX2) / 2), transY1+1, foreground)
			img.SetRGBA(((transX1 + transX2) / 2), transY1-1, foreground)
		} else if transX1 == transX2 {
			img.SetRGBA(transX1, (transY1+transY2)/2, foreground)
			img.SetRGBA(transX1+1, ((transY1 + transY2) / 2), foreground)
			img.SetRGBA(transX1-1, ((transY1 + transY2) / 2), foreground)
		} else {
			fmt.Println("invalid wall", i)
			continue
		}

	}

	//Writes the image to output.png
	encoder, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	png.Encode(encoder, img)
}
