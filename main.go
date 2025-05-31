package main

import "fmt"

func main() {
	img, _, err := DecodeImage("./test.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	grid := GridFromImage(img)
	grid.generateMaze()

}
