/*
GoLang solutions of Project Euler problems
(FYI:
This is a package comment. 
It only needs to be in one file in the package to show up in the godoc.)
*/
package main

/* Problem:
If we list all the natural numbers below 10 that are multiples of 3 or 5, 
we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

func P1() int {
	return ThreeOrFiveSum(1000)
}

/*
This algorithm comes from an old anecdote- look up Gauss' trick.
The sum of numbers from 1 to 100 can be computed in O(1):
	1 + 100 = 101
	2 + 99 = 101
	...
	50 + 51 = 101
Which means the sum is really 101 * 50 = 5050
Abstracted this is (n+1)(n/2)
*/
// Returns the sum of all numbers from 1 to n
func GaussSum (n int) int {
	return ((n+1) * n) / 2 
}

// Problem: Find the sum of all the multiples of 3 or 5 below 1000
// Which means the answer is the sums of the 3s and the 5s, MINUS the sums of the 15s
// (because 15 is a multiple of both, we would otherwise double count the 15s)
func ThreeOrFiveSum (cap int) int {
	cap-- // Because the cap is not included in the range
	threes := GaussSum(cap/3) * 3
	fives := GaussSum(cap / 5) * 5
	fifteens := GaussSum(cap / 15) * 15
	return threes + fives - fifteens
}