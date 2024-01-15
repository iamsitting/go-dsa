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

func fib3(n int) int {
	table := make([]int, n+1)
	table[0] = 1
	for i, entry := range table {
		if i+1 <= n {
			table[i+1] += entry
		}
		if i+2 <= n {
			table[i+2] += entry
		}
	}

	return table[n-1]
}

func main() {
	fmt.Println(fib3(6))
	fmt.Println(fib3(50))
	fmt.Println(fib2(6, map[int]int{}))
	fmt.Println(fib2(7, map[int]int{}))
	fmt.Println(fib2(8, map[int]int{}))
	fmt.Println(fib2(50, map[int]int{}))
}
