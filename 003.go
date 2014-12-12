package main

/* Problem:
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

func P3() int {
	return HighestPrimeFactor(600851475143)
}

// Performs a prime factorization of start, keeping track of the highest prime
func HighestPrimeFactor(start int) int {
	max := 0
	for i := 2; i <= start; i++ {
		if start % i == 0 {
			if IsPrime(i) {
				max = i
			}
			start /= i
			i = 2
		}
	}
	return max
}

// Iterates up to the prime looking for divisors, returns false if it finds any
func IsPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}