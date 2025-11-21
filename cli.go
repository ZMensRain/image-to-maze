package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ZMensRain/image-to-maze/utils"
)

func handleArgs() {

	in := flag.String("input", "./input.png", "path to the input mask must be a png")
	out := flag.String("output", "./output.png", "path to the output image must be a png")
	backgroundHex := flag.String("background-color", "#ffffff", "maze -background-color=\"#ffffff\"")
	foregroundHex := flag.String("foreground-color", "#000000", "maze -background-color=\"#000000\"")

	flag.Parse()

	parsedBackground, backErr := utils.ParseHexColor(*backgroundHex)
	parsedForeground, foreErr := utils.ParseHexColor(*foregroundHex)

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
