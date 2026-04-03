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
	for i := 0; i <= 5; i++ {
		points = append(points, plotter.XY{
			X: float64(i),
			Y: x0 * x(float64(i)) * math.Cos(math.Sqrt(g/l)*float64(i)),
		})
	}
	CreateLineplotPlot(points, "eom.png")
}
