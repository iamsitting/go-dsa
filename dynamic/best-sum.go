package main

import "fmt"

/*
* best sum takes in a targetSum and an array of numbers
* should return an array with the shortest combination of numbers
* if there is a tie, return any combination
*
* ex: bestSum(8, [2, 3, 5]) -> [3, 5]
* */

func algo(target int, numbers []int, memo map[int][]int) ([]int, bool) {
	if val, ok := memo[target]; ok {
		if val == nil {
			return val, false
		} else {
			return val, true
		}
	}
	if target == 0 {
		return []int{}, true
	}
	if target < 0 {
		return nil, false
	}

	var shortestCombo []int = nil

	for _, num := range numbers {
		rem := target - num
		remCombo, ok := algo(rem, numbers, memo)
		if ok {
			combo := append(remCombo, num)
			if shortestCombo == nil || len(combo) < len(shortestCombo) {
				shortestCombo = combo
			}
		}
	}

	memo[target] = shortestCombo
	if shortestCombo == nil {
		return shortestCombo, false
	} else {
		return shortestCombo, true
	}
}

func bestSum(target int, numbers []int) []int {
	val, ok := algo(target, numbers, map[int][]int{})
	if ok {
		return val
	} else {
		return nil
	}
}

func main() {
	res := bestSum(7, []int{5, 3, 4, 7})
	fmt.Println(res)
	res = bestSum(8, []int{2, 3, 5})
	fmt.Println(res)
	res = bestSum(8, []int{1, 4, 5})
	fmt.Println(res)
	res = bestSum(100, []int{1, 2, 5, 25})
	fmt.Println(res)
}
