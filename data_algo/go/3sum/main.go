package main

import (
	"fmt"
	"sort"
	// "sort"
)

func main() {
	o1 := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(o1)
	// o2 := twosum(-1, []int{-1, 0, 1, 2})
	// fmt.Println(o2)
	// o2 = twosum(-1, []int{-1, -1, 0, 1, 2, 2})
	// fmt.Println(o2)
}

func threeSum(nums []int) (result [][]int) {
	// i, j, k := 0, 1,
	sort.Ints(nums)
	for i := 0; i <= len(nums)-3; i++ {
		target := nums[i]
		pair := twosum(target, nums[i:])
		if len(pair) > 0 {
			pair = append(pair, nums[i])
			result = append(result, pair)
		}
	}

	return result
}

func twosum(target int, nums []int) (result []int) {
	i := 0
	k := len(nums) - 1
	tally := map[int]int{}

	for x := 0; x < len(nums)-1; x++ {
		tally[nums[x]]++
	}

	pairSum := nums[i] + nums[k]

	for  pairSum + target != 0 {
		if i == k {
			return
		}

		if pairSum 
	}

	for ; i < len(nums)-1; i++ {
		if nums[i]+nums[k]+target == 0 {
			break
		}
		
	}

	// countLower :=  tally[nums[i]]
	// for  countLower > 0 {
	// 	countUpper := nums[k]
	// 	for countUpper > 0 {
	// 		result = append(result, []int{nums[i], nums[k]})
	// 		countUpper --
	// 	}
	// 	countLower --
	// }

	return []int{nums[i], nums[k]}
}
