package main

import "fmt"

func main() {
	s := []int{4, 5, 6, 1, 2, 3}
	fmt.Println(findPeakElement(s) == 2)
	s = []int{1, 2, 3, 4, 5}
	fmt.Println(findPeakElement(s) == 4)
	s = []int{3, 4}
	fmt.Println(findPeakElement(s) == 1)
	s = []int{4, 3}
	fmt.Println(findPeakElement(s) == 0)
	s = []int{3, 2, 1}
	fmt.Println(findPeakElement(s) == 0)

}

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) / 2)
		if mid == 0 {
			// start index
			if nums[mid+1] > nums[mid] {
				// two element array where mid points smaller val
				return mid + 1
			}
			return mid
		} else if mid == len(nums)-1 {
			// end index
			if nums[mid-1] < nums[mid] {
				return mid
			}
		} else if nums[mid+1] < nums[mid] && nums[mid-1] < nums[mid] {
			// On the peak
			return mid
		} else if nums[mid-1] < nums[mid] {
			// There is a peak to the right
			left = mid + 1
		} else {
			// There is a peak to the left
			right = mid - 1
		}
	}

	return 0
}

/*
peak can be
the first element
  if index points to first element
    return current element if next element is less then current element
last element
  if index points to last element
    return current element if previous element is less then current element
or in the middle
  return current element if element is greater the both previous and next element
... binary search

Returning index, not element
*/
