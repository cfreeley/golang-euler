package main 

/* Problem:
A number chain is created by continuously adding the square of the digits in a 
number to form a new number until it has been seen before.

For example,

44 → 32 → 13 → 10 → 1 → 1
85 → 89 → 145 → 42 → 20 → 4 → 16 → 37 → 58 → 89

Therefore any chain that arrives at 1 or 89 will become stuck in an endless loop. 
What is most amazing is that EVERY starting number will eventually arrive at 1 or 89.

How many starting numbers below ten million will arrive at 89?
*/

func P92() int {
	arr, count := make([]int, 10000000), 0
	for i := 1; i < 10000000; i++ {
		if SquareDigitChain(i, arr) == 89 {
			count++
		}
	}
	return count
}

// SquareDigitChain: Relies on memoization to turn this into an O(n) solution (I think)
func SquareDigitChain(n int, arr []int) (endpoint int) {
	if n == 1 || n == 89 {
		return n
	} 
	if arr[n] != 0 {
		return arr[n]
	}
	sum := SquareDigitSum(n)
	endpoint = SquareDigitChain(sum, arr)
	arr[n] = endpoint
	return
}

// SquareDigitSum calculates the sum of the square digits of a given number
func SquareDigitSum(n int) (sum int) {
	for _, i := range NumberSlice(n) { // NumberSlice is a call back to solution 004.go
		sum += i * i
	}
	return;
}