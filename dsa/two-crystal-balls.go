package main

import (
	"fmt"
	"math"
)

// two crystal balls, 1 very tall building (n)
// find the lowest level at which the balls will break
// [false, false, ...., false, true, true, true.... true]

func findBreakPoint(breaks []bool) int {
	jumpAmount := int(math.Sqrt(float64(len(breaks))))
	fmt.Println("jump amount: $s", jumpAmount)

	i := jumpAmount

	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	fmt.Println("First ball broke at %s", i)

	i = i - jumpAmount

	fmt.Println("Second ball will start at %s and then check one by one", i)

	for j := 0; j <= jumpAmount && i < len(breaks); {
		if breaks[i] {
			return i
		}
		i++
		j++
	}

	return -1
}

func main() {
	array := []bool{false, false, false, false, false, false, false, false, false, true, true, true, true}
	result := findBreakPoint(array)
	fmt.Println(result)
}
