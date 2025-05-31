package main

import (
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type Node struct {
	outsideMask bool
	visited     bool
	openWalls   [4]bool
}
type Plane struct {
	width  int
	height int
	nodes  map[int]Node
}

func EmptyPlane(vertices, width, height int) *Plane {
	p := Plane{nodes: make(map[int]Node, vertices), width: width, height: height}
	return &p
}

func PlaneFromImage(img image.Image) *Plane {
	size := img.Bounds()
	p := EmptyPlane(size.Dx()*(size.Dy()), size.Dx(), size.Dy())

	for y := 0; y < size.Dy(); y++ {
		for x := 0; x < size.Dx(); x++ {
			index := y*size.Dx() + x
			r, g, b, _ := img.At(x, y).RGBA()
			if r == 255 && g == 255 && b == 255 {
				p.markOutsideMask(index)
			}
		}
	}
	return p
}

func (plane *Plane) markVisited(index int) {
	e := plane.nodes[index]
	e.visited = true
	plane.nodes[index] = e
}
func (plane *Plane) markOutsideMask(index int) {
	e := plane.nodes[index]
	e.outsideMask = true
	plane.nodes[index] = e
}

func (plane *Plane) findAStart() (int, error) {
	for index, node := range plane.nodes {
		if !node.outsideMask {
			return index, nil
		}
	}
	return 0, errors.New("no start")
}

func (plane *Plane) neighbors(index int) []int {
	x, y := plane.getXY(index)
	neighbors := []int{}
	if x > 0 && !plane.nodes[index-1].visited {
		neighbors = append(neighbors, index-1)
	}
	if x < plane.width && !plane.nodes[index+1].visited {
		neighbors = append(neighbors, index+1)
	}
	if y > 0 && !plane.nodes[index+plane.width].visited {
		neighbors = append(neighbors, index-plane.width)
	}
	if y < plane.height && !plane.nodes[index-plane.width].visited {
		neighbors = append(neighbors, index+plane.width)
	}
	return neighbors
}

func (plane *Plane) getXY(index int) (x, y int) {
	x = index % plane.width
	y = (index - index%plane.width) / plane.width
	return x, y
}

func (plane *Plane) generate() {
	start, err := plane.findAStart()
	if err != nil {
		fmt.Println(err)
		return
	}
	plane.iterateGeneration(start)
}

func (plane *Plane) iterateGeneration(from int) {
	plane.markVisited(from)

	for a := plane.neighbors(from); len(a) != 0; a = plane.neighbors(from) {
		// index := rand.Intn(len(a))

	}
}

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

func open(filepath string) (*Plane, error) {
	img, _, err := DecodeImage(filepath)

	if err != nil {
		return EmptyPlane(0, 0, 0), err
	}

	return PlaneFromImage(img), nil
}

func main() {
	plane, err := open("./input.png")
	if err != nil {
		return
	}
	plane.generate()
}
