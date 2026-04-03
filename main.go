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
	CreateLineplotPlot(points, "t - x", "t", "x", "eom.png")

	points2 := plotter.XYs{}
	for i := 0; i <= 500; i++ {
		points2 = append(points, plotter.XY{
			X: float64(i) / 100.0,
			Y: x0 * math.Cos(math.Sqrt(g/l)*float64(i)/100.0),
		})
	}
	CreateLineplotPlot(points2, "t - x", "t", "x", "eom.png")
}
