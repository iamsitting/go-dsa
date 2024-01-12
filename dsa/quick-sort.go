package main

import "fmt"

// quick sort picks and element (the pivot) using a partition algorithm
// the partition algorithm places the pivot in the center (weak sorting)

func quickSort(array []int, low int, high int) {
	if low <= high {
		pivotIndex := partition(array, low, high)
		quickSort(array, low, pivotIndex-1)
		quickSort(array, pivotIndex+1, high)
	}
}

func partition(array []int, low int, high int) int {
	pivot := array[high]

	index := low - 1 // this is the index where the pivot needs to be for the array to be weak sorted
	// it only increments when the element is less than then pivot
	// the value in the index will represent the right most elemen that is less than the pivot

	for i := low; i < high; i++ {
		if array[i] <= pivot {
			index++
			array[i], array[index] = array[index], array[i]
		}
	}

	index++ // one final increment to get the left most element that is greater than the pivot
	array[high], array[index] = array[index], pivot
	return index
}

func main() {
	array := []int{4, 5, 1, 500, 2431, 53, 432, 33, 55, 2, 24, 54, 42}

	fmt.Println("unsorted: ", array)

	quickSort(array, 0, len(array)-1)

	fmt.Println("sorted: ", array)
}
