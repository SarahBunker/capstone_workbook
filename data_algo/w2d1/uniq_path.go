package main

import "fmt"

func main() {
	fmt.Println("hello")
}

func uniqPaths() {

}

func helper(row, col int) int {
	if row == 0 || col == 0 {
		return 1
	}

	return helper(row, col-1) + helper(row-1, col)
}

/*
https://leetcode.com/problems/unique-paths/
*/
