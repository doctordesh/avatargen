package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

type ImageGenerator struct {
	BackgroundColor              color.Color
	ForegroundColor              color.Color
	ForegroundComplementaryColor color.Color
	Size                         int
	Scale                        int
	Steps                        int
}

func (ig *ImageGenerator) Generate(count int, destination string) {
	for i := 0; i < count; i++ {
		img_data := ig.GenerateImage()
		img := ig.DrawImage(img_data)
		err := ig.Output(img, fmt.Sprintf("%s/avatar-%d.png", destination, i+1))
		if err != nil {
			panic(err)
		}
	}
}

func (ig *ImageGenerator) GenerateImage() [][]uint8 {
	data := make([][]uint8, ig.Size)
	for i := range data {
		data[i] = make([]uint8, ig.Size)
	}

	min_x := 1
	max_x := ig.Size/2 - 1
	min_y := 1
	max_y := ig.Size - 2

	pos := Point{X: max_x, Y: RandInRange(min_y, max_y)}
	data[pos.X][pos.Y] = 1

	for i := 0; i < ig.Steps; i++ {
		for {
			potential_direction := RandomDirection()
			potential_position, err := MovePointInDirection(pos, potential_direction, min_x, min_y, max_x, max_y)
			if err != nil {
				continue
			}

			pos = potential_position
			break
		}

		data[pos.X][pos.Y] = 1
	}

	// Mirrors the left side to the right side
	for i := 0; i < len(data)/2; i++ {
		data[len(data)-i-1] = data[i]
	}

	return data
}

func (i *ImageGenerator) DrawImage(data [][]uint8) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, len(data)*i.Scale, len(data)*i.Scale))

	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			v := data[x][y]
			for sx := 0; sx < i.Scale; sx++ {
				for sy := 0; sy < i.Scale; sy++ {
					if v > 0 {
						if x < len(data)/2 {
							m.Set((x*i.Scale)+sx, (y*i.Scale)+sy, i.ForegroundColor)
						} else {
							m.Set((x*i.Scale)+sx, (y*i.Scale)+sy, i.ForegroundComplementaryColor)
						}
					} else {
						m.Set((x*i.Scale)+sx, (y*i.Scale)+sy, i.BackgroundColor)
					}
				}
			}
		}
	}

	return m
}

func (i *ImageGenerator) Output(img image.Image, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	err = png.Encode(f, img)
	if err != nil {
		return err
	}

	return nil
}

func RandInRange(from, to int) int {
	return rand.Intn(to-from) + from
}
