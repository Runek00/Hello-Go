package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	fmt.Println("Hello, World!")
	wg.Add(1)
	go func() {
		x := SumMultiples(10)
		fmt.Println("SumMultiples(10): ", x)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		x := FindPrimes(20)
		fmt.Println("FindPrimes(20): ", x)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		x := IsPalindrome("racecar")
		fmt.Println("IsPalindrome(\"racecar\"): ", x)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		x := IsPalindrome("hello")
		fmt.Println("IsPalindrome(\"hello\"): ", x)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		x := Fibonacci(7)
		fmt.Println("Fibonacci(7): ", x)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		x := ReverseLinkedList(example(3)).toStr()
		fmt.Println("ReverseLinkedList(example(3)): ", x)
		wg.Done()
	}()
	wg.Wait()
}
