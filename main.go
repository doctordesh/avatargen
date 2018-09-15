package main

import (
	"flag"
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"time"
)

const usage = "Usage: avatargen [ <flags-and-options> ] <output-directory>"

func main() {

	var size, scale, steps, count int
	var foregroundColor, backgroundColor string
	var alpha float64
	flag.IntVar(&size, "size", 32, "The size of the icon (not in pixels, see 'scale').")
	flag.IntVar(&scale, "scale", 5, "The scale of the output image. An icon with 'size' 32 and 'scale' 1 will measure 32x32 pixels. An icon with 'size' 32 and 'scale' 2 will measure 64x64 pixels.")
	flag.IntVar(&steps, "steps", 100, "Number of steps to take when drawing the image.")
	flag.IntVar(&count, "count", 100, "Number of images to generate.")
	flag.Float64Var(&alpha, "alpha", 0.9, "Alpha for the right side of the icon.")
	flag.StringVar(&foregroundColor, "color-foreground", "FFFFFF", "Foreground color, format <rrggbb> in hex.")
	flag.StringVar(&backgroundColor, "color-background", "000000", "Background color, format <rrggbb> in hex.")

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Error: no output directory specified")
		fmt.Println(usage)
		return
	}

	path := flag.Args()[0]

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("Error: Directory '%s' does not exist\n", path)
		fmt.Println(usage)
		return
	}

	rand.Seed(time.Now().UnixNano())

	fgColor, err := HexStringToRGB(foregroundColor)
	if err != nil {
		panic(err)
	}

	bgColor, err := HexStringToRGB(backgroundColor)
	if err != nil {
		panic(err)
	}

	compColorR := uint8(alpha*float64(fgColor.R) + (1-alpha)*float64(bgColor.R))
	compColorG := uint8(alpha*float64(fgColor.G) + (1-alpha)*float64(bgColor.G))
	compColorB := uint8(alpha*float64(fgColor.B) + (1-alpha)*float64(bgColor.B))

	ig := ImageGenerator{
		ForegroundColor:              fgColor,
		BackgroundColor:              bgColor,
		ForegroundComplementaryColor: color.RGBA{compColorR, compColorG, compColorB, 255},
		Size:  size,
		Scale: scale,
		Steps: steps,
	}

	ig.Generate(count, path)
}
