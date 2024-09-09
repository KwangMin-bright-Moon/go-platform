package utils

import "math"

func Round1(value, unit float32) float32 {
	v := float64(value)
	u := float64(unit)
	r := math.Round(v/u) * u
	// -0 -> 0
	if math.Abs(r) < u {
		r = 0.0
	}
	return float32(r)
}
