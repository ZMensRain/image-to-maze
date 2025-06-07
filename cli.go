package main

import (
	"flag"
	"fmt"
	"image/color"
	"os"
)

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}

func handleArgs() {

	in := flag.String("input", "./input.png", "path to the input mask must be a png")
	out := flag.String("output", "./output.png", "path to the output image must be a png")
	backgroundHex := flag.String("background-color", "#ffffff", "maze -background-color=\"#ffffff\"")
	foregroundHex := flag.String("foreground-color", "#000000", "maze -background-color=\"#000000\"")

	flag.Parse()

	parsedBackground, backErr := ParseHexColor(*backgroundHex)
	parsedForeground, foreErr := ParseHexColor(*foregroundHex)

	inputPath = *in
	outputPath = *out

	if backErr != nil {
		fmt.Println("Invalid Background color")
		os.Exit(1)
	}
	if foreErr != nil {
		fmt.Println("Invalid foreground color")
		os.Exit(2)
	}

	backgroundColor = parsedBackground
	foregroundColor = parsedForeground
}
