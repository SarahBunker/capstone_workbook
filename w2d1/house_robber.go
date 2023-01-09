// https://leetcode.com/problems/house-robber/submissions/

package main

import "fmt"

func main() {
	a := []int{2, 7, 9}
	fmt.Println(rob(a))
	b := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(b))

}

// Going backwards:

// recursive case:
//   return MAX(matrix[n] + matrix[n-2], matrix[n-1])

func rob(nums []int) int {
	memo := map[int]int{}
	return helper(nums, len(nums)-1, memo)
}

//{2,7,9}, 2, {}

func helper(nums []int, indexOfN int, memo map[int]int) int {
	if indexOfN < 0 {
		return 0
	}

	if indexOfN == 0 {
		return nums[0]
	}

	if indexOfN == 1 {
		return max(nums[0], nums[1])
	}

	i, ok := memo[indexOfN]

	if ok {
		return i
	}

	memo[indexOfN] = max(nums[indexOfN]+helper(nums, indexOfN-2, memo), helper(nums, indexOfN-1, memo))

	return memo[indexOfN]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

/*
Can't rob two adjecent houses
base

base cases:
between 1 or 2 houses return the max

if length == 1, return that num
if length == 2, return whichever is greater
--if equal, have to look at subsequent nums



func sum(index) {
	currentBestSum
}


[1,5, , 3]
[2,10,9,3,1]
				 1
			 3
		10

Base Case
	- if n < 0
		return 0


cache[indexOfn] = greatest
return greatest

curr val index n
n + 2 >
n + 3 >

[4, 3, 6, 10, 1000]

[4]
[4,3]
[4,3,6]
[4, 3, 6, 10]
[4, 3, 6, 10, 1000]


Going backwards:


recursive case:
	return MAX(matrix[n] + matrix[n-2], matrix[n-1])



greatest sum is max of
vac + greatest path among ^


*/
