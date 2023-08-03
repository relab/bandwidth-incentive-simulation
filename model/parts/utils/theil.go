package utils

import "math"

func Theil(x []int) float64 {
	if len(x) == 0 {
		return 0.0
	}

	total := 0.0
	avg := Mean(x)
	zero_count := 0
	println("##############")
	for _, xi := range x {
		print(xi, " ")
		if xi == 0 {
			xi = 1;
			zero_count ++;
		}
		dev := float64(xi) / avg
		total += dev * math.Log(dev)
	}
	println()

	println("zero count", zero_count, len(x))
	println("avg", avg)

	return total / float64(len(x))
}
