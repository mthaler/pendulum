package main

import "math"

func x(t float64) float64 {
	omega := math.Sqrt(g / l)
	return x0 * math.Cos(omega*t)
}

func phi(t float64) float64 {
	omega := math.Sqrt(g / l)
	return phimax * math.Sin(omega*t)
}
