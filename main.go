package main

import (
	"fmt"
	"image/color"
)

var backgroundColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
var foregroundColor = color.RGBA{R: 0, G: 0, B: 0, A: 0}
var inputPath = "./input.png"
var outputPath = "./output.png"

func main() {
	handleArgs()
	img, _, err := DecodeImage(inputPath)
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
	grid.renderWalls(outputPath, backgroundColor, foregroundColor)
}
