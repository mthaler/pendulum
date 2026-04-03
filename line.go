package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func CreateLineplotPlot(points plotter.XYs, title, labelx, labely string, file string) {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = labelx
	p.Y.Label.Text = labely
	p.X.Min, p.X.Max, p.Y.Min, p.Y.Max = -0, 5, -0.5, 0.5

	err := plotutil.AddLines(p, "equation of motion", points)
	if err != nil {
		log.Fatalf("could not create lineplot: %+v", err)
	}

	err = p.Save(20*vg.Centimeter, 10*vg.Centimeter, file)
	if err != nil {
		log.Fatalf("could not save lineplot: %+v", err)
	}
}
