package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func decodeImage(filename string) (image.Image, string, error) {
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

func visualize(mask [][]int8) {
	for y := range len(mask) {
		var line = mask[y]
		fmt.Println(line)
	}
}

func main() {
	const pixel_size = 8
	img, _, err := decodeImage("input.png")
	if err != nil {
		return
	}
	bounds := img.Bounds()
	e := [][]int8{}
	for y := 0; y < bounds.Dx(); y += pixel_size {
		line := []int8{}
		for x := 0; x < bounds.Dx(); x += pixel_size {

			r, g, b, _ := img.At(x, y).RGBA()
			if r == 0 && g == 0 && b == 0 {
				line = append(line, 0)
			} else {
				line = append(line, 1)
			}
		}
		e = append(e, line)
	}
	visualize(e)
}
