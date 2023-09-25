package main

func Fibonacci(n int) []int {
	out := []int{0, 1}
	for i := 2; i < n; i++ {
		out = append(out, out[i-1]+out[i-2])
	}
	return out
}
