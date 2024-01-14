package main

import (
	"fmt"
	"strings"
)

func canConstruct(target string, wordBank []string) bool {
	return solve(target, wordBank, map[string]bool{})
}

func solve(target string, wordBank []string, memo map[string]bool) bool {
	if val, ok := memo[target]; ok {
		return val
	}
	if target == "" {
		return true
	}

	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			suffix := target[len(word):]
			if solve(suffix, wordBank, memo) {
				memo[target] = true
				return true
			}
		}
	}

	memo[target] = false
	return false
}

func main() {
	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	fmt.Println(canConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeee", []string{"e", "ee", "eee", "eeee"}))
}
