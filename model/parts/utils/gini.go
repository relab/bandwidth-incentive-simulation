package utils

import "math"

func Gini(x []int) float64 {
	total := 0.0
	for i, xi := range x[:len(x)-1] {
		for _, xj := range x[i+1:] {
			total += math.Abs(float64(xi) - float64(xj))
		}
	}
	avg := Mean(x)
	denom := (math.Pow(float64(len(x)), 2) * avg)
	return total / denom
}

func Mean(x []int) float64 {
	total := 0.0
	for _, xi := range x {
		if xi > 0 {
			total += float64(xi)
		}
	}
	return total / float64(len(x))
}
