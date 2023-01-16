package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
	climbStairs(14)
}

func climbStairs(n int) int {
	memo := map[int]int{}
	return helper(n, memo)
}

func helper(n int, cache map[int]int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	res, ok := cache[n]

	if ok {
		return res
	}

	cache[n] = helper(n-1, cache) + helper(n-2, cache)

	return cache[n]
}

/*
https://leetcode.com/problems/climbing-stairs/submissions/

base case
	n=1, 1 uniq path
	n=2, 2 (1,1) (2,1)

f(n-1) + f(n-2)
n=3, (1,1,1) (2,1) (1,2)
n=4 (1,1,1,1) (2,1,1) (1,2,1) (2,2) (1,1,2)

To get to step n
we can jump 2 steps from all the paths to get to step n - 2
we can jump 1 step from all the paths to get to step n - 1

This is the fibonacci problem because we are talking about the ways, not the number of steps
*/
