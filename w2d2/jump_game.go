package main

import "fmt"

// import "fmt"

func main() {
	nums := []int{2, 0, 0, 4}
	res := canJump(nums)
	fmt.Println(res)

}

func canJump(nums []int) bool {
	memo := map[int]int{}
	bol := jump(nums, 0, memo) == 0
	fmt.Println(memo)
	return bol
}

func jump(nums []int, x int, memo map[int]int) (res int) {
	lastx := len(nums) - 1
	val := nums[x]

	if val+x >= lastx {
		return 0
	}

	if val == 0 {
		return 1
	}

	a, ok := memo[x]

	if ok {
		return a
	}

	res = 1

	for i := 1; i <= val && i+x <= lastx; i++ {
		res *= jump(nums, x+i, memo)
	}

	memo[x] = res
	return
}

/*

iterative call

for the current index
if current index is last index
  return true
if  current index is 0
	return false
if current index included in memoize
	return value from memo
if curr index not in memoize
	add return of current index to memoize
	call function with each index reachable by current value
	return || of each recursive call function


start at index 0
memoize all the indices you have visited
go to all the indices

base case getting to last index

start at last index
find
*/
