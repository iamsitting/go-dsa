// x,y grid
// begin top-left, end bottom-right
// can only move down or right
// how many possible paths from start to end
// gridTraveler(2, 3) -> 3

package main

import (
	"fmt"
	"strconv"
)

func gridTraveler(x int, y int) int {
	if x == 0 || y == 0 {
		return 0
	}

	if x == 1 && y == 1 {
		return 1
	}

	return gridTraveler(x-1, y) + gridTraveler(x, y-1)
}

func gridTraveler2(x int, y int, memo map[string]int) int {
	key := strconv.Itoa(x) + "," + strconv.Itoa(y)
	if val, ok := memo[key]; ok {
		return val
	}

	if x == 0 || y == 0 {
		return 0
	}

	if x == 1 && y == 1 {
		return 1
	}

	memo[key] = gridTraveler2(x-1, y, memo) + gridTraveler2(x, y-1, memo)
	return memo[key]
}

func main() {
	fmt.Println(gridTraveler2(2, 3, map[string]int{}))
	fmt.Println(gridTraveler2(3, 2, map[string]int{}))
	fmt.Println(gridTraveler2(1, 1, map[string]int{}))
	fmt.Println(gridTraveler2(4, 4, map[string]int{}))

	fmt.Println("done")
}
