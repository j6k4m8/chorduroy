package main

import (
	"flag"
	"strconv"

	"github.com/fogleman/gg"
)

const sX = 300
const sY = 400

const margin = 30.0
const topOffset = margin
const radius = sX / 20.

func addFingering(str float64, fret float64, context gg.Context, stringCount float64, fretCount float64) {
	str = stringCount - str
	context.SetLineWidth(2)

	if fret == -1 {
		context.MoveTo(
			margin+(str)*(sX-2*margin)/(stringCount-1)+radius,
			topOffset+margin-2.5*radius,
		)
		context.LineTo(
			margin+((str)*(sX-2*margin)/(stringCount-1))-radius,
			topOffset+margin-0.75*radius,
		)
		context.Stroke()
		context.MoveTo(
			margin+((str)*(sX-2*margin)/(stringCount-1))-radius,
			topOffset+margin-2.5*radius,
		)
		context.LineTo(
			margin+(str)*(sX-2*margin)/(stringCount-1)+radius,
			topOffset+margin-0.75*radius,
		)
		context.Stroke()
	} else {
		context.DrawEllipse(
			margin+(str*(sX-2*margin)/(stringCount-1)),
			topOffset+margin+((fret-0.5)*(sY-2*margin)/fretCount),
			radius, radius,
		)
		if fret == 0 {
			context.Stroke()
		} else {
			context.Fill()
		}
	}
}

func main() {
	fingering := flag.String("f", "X554X5", "Fingering (from highest)")
	outPath := flag.String("o", "diagram.png", "Path at which to save diagram")
	fretCount := flag.Float64("s", 6.0, "Number of frets to include")
	flag.Parse()
	drawDiagram(*fingering, *fretCount, *outPath)
}

func drawDiagram(fingering string, fretCount float64, path string) string {
	// Initialize the graphic context on an RGBA image
	stringCount := float64(len(fingering))
	gc := gg.NewContext(sX, sY)

	gc.SetRGBA(0, 0, 0, 1)

	gc.SetLineWidth(8)
	gc.MoveTo(margin, topOffset+margin)
	gc.LineTo(sX-margin, topOffset+margin)
	gc.Stroke()

	// Strings
	gc.SetLineWidth(2)
	for i := margin; i <= sX-margin; i += (sX - 2*margin) / (stringCount - 1) {
		gc.MoveTo(i, topOffset+margin)
		gc.LineTo(i, topOffset+sY-margin)
		gc.Stroke()
	}

	// Frets
	gc.SetLineWidth(1)
	for i := margin; i <= sY-margin; i += (sY - 2*margin) / fretCount {
		gc.MoveTo(margin, topOffset+i)
		gc.LineTo(sX-margin, topOffset+i)
		gc.Stroke()
	}

	// Draw chord fingers
	for i := 0; i < len(fingering); i++ {
		if fingering[i] == byte("X"[0]) {
			addFingering(float64(i+1), -1., *gc, stringCount, fretCount)
		} else {
			val, _ := strconv.Atoi(string(fingering[i]))
			addFingering(float64(i+1), float64(val), *gc, stringCount, fretCount)
		}
	}

	gc.SavePNG(path)
	return path
}
