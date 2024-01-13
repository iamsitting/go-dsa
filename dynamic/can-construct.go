package main

import (
	"fmt"
	"strings"
)

/**
* the function should return a boolean indidcating whether or not the target can be constructed by contatenating elements of the word bank array
* canConstruct("abcdef", ["ab", "abc", "cd", "def", "abcd"]) -> true
* canContruct("skateboard", ["bo", "rd", "ate", "t", "ska", "sk", "boar"]) -> false
 */
func canConstruct(target string, wordBank []string) bool {
	return algo(target, wordBank)
}

func algo(target string, wordBank []string) bool {
	if isBaseCase(target) {
		return true
	}

	for _, word := range wordBank {
		newTarget := strings.Replace(target, word, "*", 1)
		if newTarget == target {
			continue
		}
		res := algo(newTarget, wordBank)
		if res {
			return true
		}
	}

	return false
}

/**
* if string only contains *
 */
func isBaseCase(s string) bool {
	for _, letter := range s {
		if letter != '*' {
			return false
		}
	}
	return true
}

func main() {
	res := canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"})
	fmt.Println(res)

	res = canConstruct("abcgdef", []string{"ab", "abc", "cd", "def", "abcd"})
	fmt.Println(res)
}
