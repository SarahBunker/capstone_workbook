package main

import "fmt"

func main() {
	fmt.Println("Hello")
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
		{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0}}
	res := minimumTotal(triangle)
	fmt.Println(res)

}

func minimumTotal(triangle [][]int) int {
	memo := map[[2]int]int{}
	return minSum(triangle, 0, 0, memo)
}

func minSum(triangle [][]int, row, col int, memo map[[2]int]int) int {
	rowMax := len(triangle) - 1
	if row == rowMax {
		return triangle[row][col]
	}

	directBottom := minSum(triangle, row+1, col, memo)
	oneLeftBottom := minSum(triangle, row+1, col+1, memo)
	res := min(directBottom, oneLeftBottom) + triangle[row][col]
	memo[[2]int{row, col}] = res
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/*
Can access indices on the next row with the same or index plus one
 if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.

base case
	reaching the last row with total

memo < store the best sum at each index of the last row

base
	sum at bottom row is value at current index

recursion
	return min()

*/

//-------------------------------------------------------------------------------------------------

/*

Recursive Solution

func minimumTotal(triangle [][]int) int {
	memo := map[[2]int]int{}
	return minSum(triangle, 0, 0, memo)
}

func minSum(triangle [][]int, row, col int, memo map[[2]int]int) int {
	rowMax := len(triangle) - 1
	if row == rowMax {
		return triangle[row][col]
	}

	directBottom := minSum(triangle, row+1, col, memo)
	oneLeftBottom := minSum(triangle, row+1, col+1, memo)
	res := min(directBottom, oneLeftBottom) + triangle[row][col]
	memo[[2]int{row, col}] = res
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

*/
