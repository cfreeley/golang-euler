package main

func P3() int {
	return HighestPrimeFactor(600851475143)
}

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

func IsPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}