package main

import "fmt"

func canJump(nums []int) bool {
	memo := map[int]bool{}

	// Base Case: If we are at the last position, then we know we can reach the end
	memo[len(nums)-1] = true

	// Start from the next to last position
	for index := len(nums) - 2; index >= 0; index-- {
		// The farthest we can go from this position = the position we are on plus the value at that position
		maxLeap := index + nums[index]

		// For each possible step from this index, see if we can get to the end via one of those steps
		// If we can, we can also get to the end via this step
		// If we can't, then the value for memo[index] remains false (the default)
		for step := index + 1; step <= maxLeap; step++ {
			if memo[step] {
				memo[index] = true
				break
			}
		}
	}

	// If memo[0] is true, we can reach the end from the first position
	return memo[0]
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums) == true)
	nums = []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums) == false)
}
