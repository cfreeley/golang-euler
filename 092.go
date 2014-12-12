package main 

import "fmt"

func main() {
	fmt.Println(P92())
}

func P92() int {
	arr, count := make([]int, 10000000), 0
	for i := 1; i < 10000000; i++ {
		// fmt.Println(i)
		if SquareDigitChain(i, arr) == 89 {
			count++
		}
	}
	return count
}

func SquareDigitChain(n int, arr []int) (endpoint int) {
	fmt.Println(n)
	// if n == 0 {
	// 	return 1
	// }
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

func SquareDigitSum(n int) (sum int) {
	for i := range NumberSlice(n) {
		sum += i * i
	}
	return;
}