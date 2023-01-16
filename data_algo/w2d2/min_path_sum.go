package main

import "fmt"

func main() {
	fmt.Println("Hello")
	arr := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	y := minPathSum(arr)
	fmt.Println(y)
}

func minPathSum(grid [][]int) int {
	memo := map[[2]int]int{}
	min := helper(grid, 0, 0, memo)
	// fmt.Println(memo)
	return min
}

func minSum(grid, row, col, memo) int {
	maxVal := 201
	maxRow := len(grid) - 1
	maxCol := len(grid[0]) - 1
	right, ok := memo[row][col + 1]
	if !ok {
		right = maxVal
	}
	down, ok := memo[row + 1][col]
	if !ok {
		down = maxVal
	}

	return findMin(right, down) + grid[]
}

func findMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/*
Recursive approach:
memo a nested array to store the min path sum from each location

start in bottom right
min path is current value

next level
min path is current value plus the smaller value between down and right
store min path in memo

stop when you reach bottom left
	m and n are greatest values
*/

//-------------------------------------------------------------------------------------------------

/*
Recursive Solution

func minPathSum(grid [][]int) int {
	memo := map[[2]int]int{}
	min := helper(grid, 0, 0, memo)
	// fmt.Println(memo)
	return min
}

func helper(grid [][]int, row, col int, memo map[[2]int]int) (min int) {
	maxRow := len(grid) - 1
	maxCol := len(grid[0]) - 1
	if row == maxRow && col == maxCol {
		// memo[[2]int{row, col}] = grid[row][col]
		return grid[row][col]
	}

	if row == maxRow {
		min = helper(grid, row, col+1, memo) + grid[row][col]
		// memo[[2]int{row, col}] = min
		return
	}

	if col == maxCol {
		min = helper(grid, row+1, col, memo) + grid[row][col]
		// memo[[2]int{row, col}] = min
		return
	}

	i, ok := memo[[2]int{row, col}]

	if ok {
		return i
	}
	min = findMin(helper(grid, row, col+1, memo), helper(grid, row+1, col, memo)) + grid[row][col]
	memo[[2]int{row, col}] = min
	return
}

func findMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
*/
