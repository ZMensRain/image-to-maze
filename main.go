package main

import "fmt"

func main() {

	img, _, err := DecodeImage("./input.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("creating the grid based on your image")
	grid := GridFromImage(img)
	fmt.Println("Generation Started")
	for i := grid.unvisited(); i != -1; i = grid.unvisited() {
		grid.generateMaze(i)
	}
	fmt.Println("Generation Finished")
	grid.renderWalls()
}
