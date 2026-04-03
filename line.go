package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func CreateLineplotPlot(points plotter.XYs, file string) {
	p := plot.New()
	p.Y.Min, p.X.Min, p.Y.Max, p.X.Max = 0, 0, 10, 10

	err := plotutil.AddLines(p, "line1", points)
	if err != nil {
		log.Fatalf("could not create lineplot: %+v", err)
	}

	err = p.Save(20*vg.Centimeter, 10*vg.Centimeter, file)
	if err != nil {
		log.Fatalf("could not save lineplot: %+v", err)
	}
}
