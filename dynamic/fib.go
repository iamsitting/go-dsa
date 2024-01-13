package main

import "fmt"

// 1, 1, 2, 3, 5, 8, 13, 21, 34
// input: nth fib number
// output: fib number
// O -> 2^n
func fib(n int) int {
	if n <= 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

// O -> n
func fib2(n int, memo map[int]int) int {
	if n <= 2 {
		return 1
	}

	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = fib2(n-1, memo) + fib2(n-2, memo)
	return memo[n]
}

func main() {
	fmt.Println(fib2(6, map[int]int{}))
	fmt.Println(fib2(7, map[int]int{}))
	fmt.Println(fib2(8, map[int]int{}))
	fmt.Println(fib2(50, map[int]int{}))
}
