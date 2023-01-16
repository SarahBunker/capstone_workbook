package main

import "fmt"

func main() {
	s := []int{4, 5, 6, 1, 2, 3}
	fmt.Println(findMin(s))
	s = []int{2, 3, 4, 5, 6}
	fmt.Println(findMin(s))
	s = []int{3, 4}
	fmt.Println(findMin(s))
	s = []int{4, 3}
	fmt.Println(findMin(s))
	s = []int{5, 6, 4}
	fmt.Println(findMin(s))

}

func findMin(nums []int) int {
	max := len(nums) - 1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) / 2)
		if mid == 0 {
			if nums[mid] < nums[mid+1] {
				return nums[mid]
			}
			return nums[mid+1]
		} else if nums[mid] < nums[mid-1] {
			return nums[mid]
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nums[max]
}

/*
if array length is one return value at index zero

if middle index is at initial index
	if next value is less then current value
	YES:
		return next value
	NO:
		return current value

if middle points to value where previous value is greater
	YES: return current value






*/
