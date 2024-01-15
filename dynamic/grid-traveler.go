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

func gridTraverler3(x int, y int) int {
	table := make([][]int, y+1)
	for rowI := range table {
		table[rowI] = make([]int, x+1)
	}
	table[0][0] = 0
	table[1][1] = 1
	fmt.Println("table:", table)
	for j := 0; j <= y; j++ {
		for i := 0; i <= x; i++ {
			if i+1 <= x {
				table[j][i+1] += table[j][i]
			}

			if j+1 <= y {
				table[j+1][i] += table[j][i]
			}
		}
	}

	return table[y][x]
}

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
	fmt.Println(gridTraverler3(2, 3))
	fmt.Println(gridTraveler2(2, 3, map[string]int{}))
	fmt.Println(gridTraveler2(3, 2, map[string]int{}))
	fmt.Println(gridTraveler2(1, 1, map[string]int{}))
	fmt.Println(gridTraveler2(4, 4, map[string]int{}))

	fmt.Println("done")
}
