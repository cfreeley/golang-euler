package main

import (
	"fmt"
)

func main() {

}

// Fib calculates the nth fibonacci number
// Contract: n >= 0
func Fib(n int) (answer int) {
	return fibAcc(n, 0, 1) 
}
// FibAcc is a helper accumulator for Fib
func fibAcc (n, prev, current int) int {
	fmt.Println(current)
	if n <= 0 {
		return current;
	}
	return fibAcc(n-1, current, current + prev)
}

