package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

const sX = 300
const sY = 400

const margin = 30.0

const topOffset = margin
const leftOffset = margin

const radius = sX / 20.

func addFingering(str float64, fret float64, context gg.Context, stringCount float64, fretCount float64) {
	str = stringCount - str
	context.SetLineWidth(2)

	if fret == -1 {
		// Draw an X
		context.MoveTo(
			leftOffset+margin+(str)*(sX-(2*margin+leftOffset))/(stringCount-1)+radius,
			topOffset+margin-2.5*radius,
		)
		context.LineTo(
			leftOffset+margin+((str)*(sX-(2*margin+leftOffset))/(stringCount-1))-radius,
			topOffset+margin-0.75*radius,
		)
		context.Stroke()
		context.MoveTo(
			leftOffset+margin+((str)*(sX-(2*margin+leftOffset))/(stringCount-1))-radius,
			topOffset+margin-2.5*radius,
		)
		context.LineTo(
			leftOffset+margin+(str)*(sX-(2*margin+leftOffset))/(stringCount-1)+radius,
			topOffset+margin-0.75*radius,
		)
		context.Stroke()
	} else {
		context.DrawEllipse(
			leftOffset+margin+(str*(sX-(2*margin+leftOffset))/(stringCount-1)),
			topOffset+margin+((fret-0.5)*(sY-(2*margin))/fretCount),
			radius, radius,
		)
		if fret == 0 {
			// Draw the hollow open-string "O"
			context.Stroke()
		} else {
			// Draw the solid fingering
			context.Fill()
		}
	}
}

func main() {
	fingering := flag.String("f", "X554X5", "Fingering (from highest)")
	outPath := flag.String("o", "", "Path at which to save diagram")
	fretCount := flag.Float64("s", 6.0, "Number of frets to include")
	flag.Parse()
	if *outPath == "" {
		fmt.Println("Need a -o path.")
		os.Exit(1)
	}
	drawDiagram(*fingering, *fretCount, *outPath)
}

func drawDiagram(fingering string, fretCount float64, path string) string {

	minFret := 999
	for i := 0; i < len(fingering); i++ {
		if fingering[i] != byte("X"[0]) {
			val, _ := strconv.Atoi(string(fingering[i]))
			minFret = int(math.Min(float64(minFret), float64(val)))
		}
	}

	// Initialize the graphic context on an RGBA image
	stringCount := float64(len(fingering))
	gc := gg.NewContext(sX, sY)

	gc.SetRGBA(0, 0, 0, 1)

	// TODO: Smartify this based upon fretCount and number of visible frets
	fretOffset := 0
	if minFret > 5 {
		fretOffset = minFret - 2
		font, err := truetype.Parse(goregular.TTF)
		if err != nil {
			log.Fatal(err)
		}
		face := truetype.NewFace(font, &truetype.Options{Size: radius * 2})
		gc.SetFontFace(face)
		gc.DrawString(strconv.Itoa(int(minFret-2)), leftOffset/2, margin+topOffset+radius)
	}

	// Nut
	gc.SetLineWidth(8)
	gc.MoveTo(leftOffset+margin, topOffset+margin)
	gc.LineTo(sX-margin, topOffset+margin)
	gc.Stroke()

	// Strings
	gc.SetLineWidth(2)
	for i := margin; i <= sX-(margin+leftOffset); i += (sX - (2*margin + leftOffset)) / (stringCount - 1) {
		gc.MoveTo(leftOffset+i, topOffset+margin)
		gc.LineTo(leftOffset+i, topOffset+sY-margin)
		gc.Stroke()
	}

	// Frets
	gc.SetLineWidth(1)
	for i := margin; i <= sY-margin; i += (sY - 2*margin) / fretCount {
		gc.MoveTo(leftOffset+margin, topOffset+i)
		gc.LineTo(sX-margin, topOffset+i)
		gc.Stroke()
	}

	// Draw chord fingers
	for i := 0; i < len(fingering); i++ {
		if fingering[i] == byte("X"[0]) {
			addFingering(float64(i+1), -1., *gc, stringCount, fretCount)
		} else {
			val, _ := strconv.Atoi(string(fingering[i]))
			addFingering(float64(i+1), float64(val-fretOffset), *gc, stringCount, fretCount)
		}
	}

	gc.SavePNG(path)
	return path
}
