package main

import "fmt"

func main() {
	img, _, err := DecodeImage("./input.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	grid := GridFromImage(img)
	grid.generateMaze(2)
	grid.generateMaze(0)

	grid.renderWalls()
}
