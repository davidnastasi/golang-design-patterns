package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

type OutputStrategy interface {
	Draw() error
}

type ConsoleSquare struct {}

type ImageSquare struct {
	DestinationFilePath string
}

func (c *ConsoleSquare) Draw() error {
	println("Square")
	return nil
}

func (t *ImageSquare) Draw() error {
	width := 800
	height := 600

	origin := image.Point{0, 0}

	bgRectangle := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})

	bgColor := image.Uniform{C: color.RGBA{R: 70, G: 70, B: 70, A:0}}
	draw.Draw(bgRectangle, bgRectangle.Bounds(), &bgColor, origin, draw.Src)

	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{C: color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareImg := image.NewRGBA(square)

	draw.Draw(bgRectangle, squareImg.Bounds(), &squareColor, origin, draw.Src)

	w, err := os.Create(t.DestinationFilePath)
	if err != nil {
		return fmt.Errorf("error opening image")
	}
	defer w.Close()

	quality := &jpeg.Options{Quality: 75}
	if err = jpeg.Encode(w, bgRectangle, quality); err != nil {
		return fmt.Errorf("error writing image to disk")
	}

	return nil

}

func main() {
	var output = flag.String("output", "console", "The output to use between 'console' and 'image' file")
	flag.Parse()

	var activeStrategy OutputStrategy

	switch *output {
	case "console":
		activeStrategy = &ConsoleSquare{}
	case "image":
		activeStrategy = &ImageSquare{"/tmp/image.jpg"}
	default:
		activeStrategy = &ConsoleSquare{}
	}

	err := activeStrategy.Draw()
	if err != nil {
		log.Fatal(err)
	}
}