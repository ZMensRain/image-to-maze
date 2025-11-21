package utils

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"

	"github.com/ZMensRain/image-to-maze/core"
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

func SaveMazeToFile(g *core.Grid, path string, foreground, background color.RGBA) error {
	//Writes the image to output.png
	encoder, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	img := g.RenderWalls(background, foreground)

	return png.Encode(encoder, img)
}
