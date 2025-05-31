package main

import (
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
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
		return 1
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
func (g *Grid) findStart() (index int, err error) {
	for i, state := range g.cellState {
		//checks if the cell is in the mask
		if state == 0 {
			return i, nil
		}
	}
	return -1, errors.New("no mask found")
}

func (g *Grid) generateMaze() {

}
