package main

import "fmt"

// numbers: arbitrary array of numbers
// targetSum: an arbitrary value
// return: can any combintation of numbers within the array of numbers be used to sum the targetSum
// you may use the numbers in the array as many times as needed
// assume numbers in the array are non negative

// O -> targetSum ^ len(numbers)
func canSum(targetSum int, numbers []int) bool {
	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for i := 0; i < len(numbers); i++ {
		res := canSum(targetSum-numbers[i], numbers)
		if res {
			return true
		}
	}

	return false
}

func canSum2(targetSum int, numbers []int, memo map[int]bool) bool {
	if val, ok := memo[targetSum]; ok {
		return val
	}

	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for _, num := range numbers {
		rem := targetSum - num
		res := canSum2(rem, numbers, memo)
		if res {
			memo[targetSum] = true
			return true
		}
	}

	memo[targetSum] = false
	return false
}

func main() {
	array1 := []int{5, 3, 4, 7}
	array2 := []int{2, 4}
	array3 := []int{7, 14}

	fmt.Println(canSum2(7, array1, map[int]bool{}))
	fmt.Println(canSum2(7, array2, map[int]bool{}))
	fmt.Println(canSum2(300, array3, map[int]bool{}))
}
