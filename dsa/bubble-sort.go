package main

import "fmt"

func bubbleSort(array []int) {
	length := len(array)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}

func main() {
	arr1 := []int{4, 5, 6, 1, 35, 14, 9, 2, 4}
	bubbleSort(arr1)
	fmt.Println(arr1)
}
