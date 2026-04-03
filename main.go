package main

import (
	"math"

	"gonum.org/v1/plot/plotter"
)

const (
	x0 = 0.3
	g  = 9.81
	l  = 4.0
)

func main() {
	points := plotter.XYs{}
	for i := 0; i <= 500; i++ {
		points = append(points, plotter.XY{
			X: float64(i) / 100.0,
			Y: x0 * math.Cos(math.Sqrt(g/l)*float64(i)/100.0),
		})
	}
	b := bounds{
		xmin: 0,
		xmax: 5,
		ymin: -0.5,
		ymax: 0.5,
	}
	l := labels{
		x: "t",
		y: "x",
	}
	CreateLineplotPlot(points, "t - x", l, b, "eom.png")
}
