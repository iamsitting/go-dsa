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

func solve(targetSum int, numbers []int) ([]int, error) {
	ok, res := howSum2(targetSum, numbers, []int{}, map[int]bool{})
	if ok {
		return res, nil
	} else {
		return nil, fmt.Errorf("the sum can't be made with these numbers")
	}
}

func howSum3(targetSum int, numbers []int, memo map[int][]int) ([]int, error) {
	if val, ok := memo[targetSum]; ok {
		if val != nil {
			return val, nil
		} else {
			return nil, fmt.Errorf("no more")
		}
	}
	if targetSum == 0 {
		return []int{}, nil
	}
	if targetSum < 0 {
		return nil, fmt.Errorf("no more")
	}

	for _, num := range numbers {
		rem := targetSum - num
		res, e := howSum3(rem, numbers, memo)
		if e != nil {
			memo[targetSum] = append(res, num)
			return memo[targetSum], nil
		}
	}

	memo[targetSum] = nil

	return nil, fmt.Errorf("no more")
}

func main() {
	array1 := []int{5, 3, 4, 7}
	array2 := []int{2, 4}
	array3 := []int{7, 14}

	res, e := solve(7, array1)
	fmt.Println(res, e)
	res, e = solve(7, array2)
	fmt.Println(res, e)
	res, e = solve(300, array3)
	fmt.Println(res, e)
}
