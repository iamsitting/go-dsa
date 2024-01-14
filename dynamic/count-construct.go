package main

import (
	"fmt"
	"strings"
)

// the function shoudl return the number of ways that the "target" can
// be constructed by concatenating elements of the wordBank array
// the words in the wordBank may be reused
func countConstruct(target string, wordBank []string) int {
	return solve(target, wordBank, map[string]int{})
}

func solve(target string, wordBank []string, memo map[string]int) int {
	if val, ok := memo[target]; ok {
		return val
	}
	if target == "" {
		return 1
	}

	count := 0
	for _, word := range wordBank {
		if strings.Index(target, word) == 0 { // target, tar -> 0
			suffix := target[len(word):] // target, tar -> get
			count += solve(suffix, wordBank, memo)
		}
	}

	memo[target] = count
	return count
}

func main() {
	fmt.Println(countConstruct("eeee", []string{"e", "ee"}))
}
