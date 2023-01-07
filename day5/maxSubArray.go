// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
)

/*
Recursive divide-and-conquer solution to get Right Sum and Left Sum

maxSubArray
	maxLeftAndRight = helper(arr, 0, len(arr)-1)
	crossSum = cross(arr)
	return max(maxLeftAndRight, crossSum)

helper(arr []int, start, end int) int {
if start == end
	return the value at the array
else
	mid = start + end /2
	leftSum = helper(arr, start, mid)
	rightSum = helper(arr, mid + 1, end)
	return max(leftSum, rightSum)
}

Cross Sum solution
	take midpoint

	leftTotal, initalize to val at mid
		iterate leftwards from midpoint
			currentSum adds that element
			if currentSum > leftTotal
				update LeftTotal = currentSum
			until end (check index against length, stop when index is 0

	rightTotal, initalize to val at mid+1
		iterate rightwards from midpoint
      currentSum adds that element
      if currentSum > righTotal
        update LeftTotal = currentSum
			until end (check index against length, stop when index is length minus one

Return greatest of RS, LS, CS.


*/

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4} // 6 {4 -1 2 1}
	fmt.Println(maxSubArray(nums) == 6)
	// fmt.Println(cross(nums))
	nums = []int{1} // 1     {1}
	fmt.Println(maxSubArray(nums) == 1)
	nums = []int{5, 4, -1, 7, 8} // 23 full array
	fmt.Println(maxSubArray(nums) == 23)
	nums = []int{5, 4, -1, -7, -8} // 9 left sum  {5 4}
	fmt.Println(maxSubArray(nums) == 9)
	nums = []int{-5, -4, -1, 7, 8} // 15 {7, 8}
	fmt.Println(maxSubArray(nums) == 15)
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// maxLeftAndRight := helper(nums, 0, len(nums) -1)
	// crossSum := cross(nums)
	right := len(nums) - 1
	return helper(nums, 0, right)
}

func helper(nums []int, start, end int) int {
	// fmt.Println(nums, start, end)
	if start == end {
		// fmt.Println("nums start: "nums[start])
		return nums[start]
	}
	mid := start + (end-start)/2
	leftSum := helper(nums, start, mid)
	// fmt.Println("left sum:" , leftSum)
	rightSum := helper(nums, mid+1, end)
	crossMax := cross(nums, start, end)
	// fmt.Println("right sum: ", rightSum)
	// fmt.Println("max sum: ", max(leftSum, rightSum))
	return max(leftSum, rightSum, crossMax)
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func cross(nums []int, left, right int) (res int) {
	if len(nums) == 1 {
		return nums[0]
	}
	mid := left + (right-left)/2
	lt := nums[mid]
	cs := lt
	for i := mid - 1; i >= left; i-- {
		cs += nums[i]
		if cs > lt {
			lt = cs
		}
	}
	rt := nums[mid+1]
	cs = rt
	for i := mid + 2; i <= right; i++ {
		cs += nums[i]
		if cs > rt {
			rt = cs
		}
	}
	return rt + lt
}
