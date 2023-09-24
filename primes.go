package main

import (
	"math"
)

func FindPrimes(n int) []int {
	var out []int = make([]int, 0, int(math.Sqrt(float64(n))))
	if n < 0 {
		return out
	}
	for i := 2; i < n; i++ {
		isP := true
		for _, p := range out {
			if p <= int(math.Sqrt(float64(i))) && i%p == 0 {
				isP = false
			}
		}
		if isP {
			out = append(out, i)
		}
	}
	return out
}
