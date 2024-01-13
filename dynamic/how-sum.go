package main

import "fmt"

// targetSum: arbitrary value
// numbers: arbitrary array of values
// array containing any combinatio off elements that add up to targetSum
// if none exist, return error
// if multiple, return any combination
//

func howSum(targetSum int, numbers []int, addends []int) (bool, []int) {
	if targetSum == 0 {
		return true, addends
	}

	if targetSum < 0 {
		return false, addends
	}

	for _, num := range numbers {
		res, newAddends := howSum(targetSum-num, numbers, addends)
		if res {
			return true, append(newAddends, num)
		}
	}

	return false, addends
}

func howSum2(targetSum int, numbers []int, addends []int, memo map[int]bool) (bool, []int) {
	if val, ok := memo[targetSum]; ok {
		return val, addends
	}

	if targetSum == 0 {
		return true, addends
	}

	if targetSum < 0 {
		return false, addends
	}

	for _, num := range numbers {
		res, newAddends := howSum2(targetSum-num, numbers, addends, memo)
		if res {
			memo[targetSum] = true
			return true, append(newAddends, num)
		}
	}

	memo[targetSum] = false
	return false, addends
}

func main() {
	array1 := []int{5, 3, 4, 7}
	array2 := []int{2, 4}
	array3 := []int{7, 14}

	_, res := howSum2(7, array1, []int{}, map[int]bool{})
	fmt.Println(res)
	_, res = howSum2(7, array2, []int{}, map[int]bool{})
	fmt.Println(res)
	_, res = howSum2(300, array3, []int{}, map[int]bool{})
	fmt.Println(res)

	// fmt.Println(canSum(300, array3, map[int]bool{}))
}
