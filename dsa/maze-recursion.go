package main

// you need to figure out how to get solve a maze
// this is the input
// ########E##
// ######   ##
// ###    ####
// ###S#######
// let's assume we have a square (X = Y)

// Clearly a recurive problem
// Base Case:
// invalid positions - false
// if hits wall - false
// if i have visited before - false
// if at end - true
// otherwise - walk recursively
// I want the path

type Position struct {
	x, y int
}

var direction = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func walk(maze [][]byte, current Position, end Position, seen [][]bool, path []Position) bool {
	// invalid
	if current.x < 0 || current.y < 0 || current.y > len(maze)-1 || current.x > len(maze[0])-1 {
		return false
	}

	// hits wall
	if maze[current.y][current.x] == '#' {
		return false
	}

	// have been here before
	if seen[current.y][current.x] {
		return false
	}

	// found the end
	if current.y == end.y && current.x == end.x {
		return true
	}

	seen[current.y][current.x] = true
	path = append(path, current)

	// now we need to recurse in 4 directions: x+1, x-1, y+1, y-1
	for step := 0; step < len(direction); step++ {
		thatWay := direction[step]
		nextPosition := Position{x: current.x + thatWay[0], y: current.y + thatWay[1]}
		foundEnd := walk(maze, nextPosition, end, seen, path)
		if foundEnd {
			return true
		}
	}

	// current position found a dead end
	path = path[:len(path)-1]
	return false
}

func entryPoint(maze [][]byte, start Position, end Position) []Position {
	path := []Position{}
	seen := [][]bool{}

	for i := 0; i < len(maze); i++ {
		seen[i] = append(seen[i], false)
	}

	walk(maze, start, end, seen, path)
	return path
}
