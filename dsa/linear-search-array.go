package main

import (
	"fmt"
)

func linearSearch(array []int, val int) bool {
	n := len(array)
	for i := 0; i < n; i++ {
		if val == array[i] {
			return true
		}
	}
	return false
}

func main() {
	arr1 := []int{1, 3, 4, 45, 65, 44, 123, 6}
	found := linearSearch(arr1, 45)
	fmt.Println(found)
}
