package main

import (
	"fmt"
	"strings"
)

/**
* function shoudl return a 2D array containing all of the ways the target can be constructed
* by concatenating elements of the word array
* Each element of the 2D array should represent one cobinvation that constrcuts the target
 */
func allConstruct(target string, wordBank []string) [][]string {
	return solve(target, wordBank)
}

func solve(target string, wordBank []string) [][]string {
	if target == "" {
		return [][]string{{}}
	}

	result := [][]string{}

	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			suffix := target[len(word):]
			suffixCombo := solve(suffix, wordBank)
			for _, item := range suffixCombo {
				combination := append([]string{word}, item...)
				result = append(result, combination)
			}

		}
	}

	return result
}

func main() {
	fmt.Println(allConstruct("eeeee", []string{"e", "ee"}))
}
