package main

import "fmt"

func binarySearch(array []int, val int, minIndex int, maxIndex int) bool {
	for minIndex < maxIndex {

		midIndex := minIndex + ((maxIndex - minIndex) / 2)
		midPoint := array[midIndex]

		if val == midPoint {
			return true
		} else if val > midPoint {
			minIndex = midIndex + 1
		} else {
			maxIndex = midIndex
		}

	}

	return false
}

func main() {
	array1 := []int{1, 3, 4, 6, 23, 25, 27, 31, 32, 43, 45, 54, 65}

	found := binarySearch(array1, 23, 0, len(array1)-1)

	fmt.Println("found", found)

	found = binarySearch(array1, 32, 0, len(array1)-1)
	fmt.Println("found", found)
}
