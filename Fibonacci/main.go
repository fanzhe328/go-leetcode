package main

import "fmt"

func main() {
	fmt.Println(Fibonacci(1))
	fmt.Println(Fibonacci(2))
	fmt.Println(Fibonacci(3))
	fmt.Println(Fibonacci(10))
	fmt.Println(FibonacciRecursion(10))
}

func Fibonacci(N int) int {
	if N <= 1 {
		return N
	}
	first := 0
	second := 1
	third := 0
	for i := 2; i <= N; i++ {
		third = first + second
		first = second
		second = third
	}
	return second
}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}
