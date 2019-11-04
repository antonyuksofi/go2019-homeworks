package image

import (
	"fmt"
	"homeworks/third/data"
)

const DefaultWidth = 38
const DefaultHeight = 18

type Image struct {
	width  int
	height int
	points []ColoredPoint
}

func NewImage() *Image {
	image := Image{DefaultWidth, DefaultHeight, make([]ColoredPoint, DefaultWidth*DefaultHeight)}
	return &image
}

type ColoredPoint struct {
	// Point
	color Color
}

type Color uint8

func (image Image) Load(rawData data.RawData) {
	if len(image.points) != len(rawData) {
		panic("PANIC: the sizes of image and raw data are different")
	}

	for i, v := range rawData {
		image.points[i].color = Color(v)
	}
}

type Printer struct {
	Driver
}

func NewPrinter() *Printer {
	colorsToSymbols := map[Color]string{
		0: "-",
		4: ":",
		9: "G",
	}

	printer := Printer{Driver{colorsToSymbols}}

	return &printer
}

func (pr Printer) Print(img Image) {
	for i := 0; i < DefaultWidth; i++ {
		pr.Driver.PrintColor(img.points[i].color)
	}
	fmt.Println()
}

func (pr Printer) Println(img Image) {
	for i := 0; i < len(img.points); i++ {
		pr.Driver.PrintColor(img.points[i].color)
		if i%img.width == img.width-1 {
			fmt.Println()
		}
	}
	fmt.Println()
}

type Driver struct {
	colorsMap map[Color]string
}

func (dr Driver) PrintColor(color Color) {
	currentSymbolToPrint, _ := dr.colorsMap[color]

	fmt.Print(currentSymbolToPrint)
}
