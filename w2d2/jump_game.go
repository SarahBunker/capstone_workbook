package main

// import "fmt"

func main() {

}

func canJump(nums []int) bool {

	return false
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
