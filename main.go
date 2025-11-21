package main

import (
	"fmt"
	"image/color"

	"github.com/ZMensRain/image-to-maze/core"
	"github.com/ZMensRain/image-to-maze/utils"
)

var backgroundColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
var foregroundColor = color.RGBA{R: 0, G: 0, B: 0, A: 0}
var inputPath = "./input.png"
var outputPath = "./output.png"

func main() {
	handleArgs()
	img, _, err := utils.DecodeImage(inputPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("creating the grid based on your image")
	grid := core.GridFromImage(img)
	fmt.Println("Generation Started")
	for i := grid.FindUnvisited(); i != -1; i = grid.FindUnvisited() {
		grid.GenerateMaze(i)
	}
	fmt.Println("Generation Finished")
	utils.SaveMazeToFile(grid, outputPath, foregroundColor, backgroundColor)
}
