package main

func SumMultiples(n int) int {
	out := 0
	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			out += i
		}
	}
	return out
}
