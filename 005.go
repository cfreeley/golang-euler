package main

/* Problem:
2520 is the smallest number that can be divided by each of the numbers 
from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by 
all of the numbers from 1 to 20?
*/

func P5() int {
	return SmallestEvenlyDivisibleBy(20)
}

// Finds the smallest number which is evenly divisible by all ints up to n
func SmallestEvenlyDivisibleBy(n int) int {
	i := n 
	for {
		if allEvenlyDisibleBy(i, n) {
			return i
		}
		i += n // since answer is a multiple of n, we can increment by n
	}
}

// Checks if the candidate can be evenly divided by all numbers up to goal
func allEvenlyDisibleBy(candidate, goal int) bool {
	for i := goal; i > 1; i-- {
		if candidate % i > 0 {
			return false
		}
	}
	return true
}