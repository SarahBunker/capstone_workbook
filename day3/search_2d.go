package main

func main() {

}

/*
binary search

look for the row that has the element
  the target will be greater then the first element
  and less then the last element in the array

initialize start, middle, and end
set j to 0
current element is element at middle and j indexes
if current element is the target return true
if target is greater then current element
  if row could contains element
    break
  if middle equals end
    return false
  select the right half: start to middle +1
if target is less then current element
  if middle equals 0
    return false
  select the left half: end to middle - 1


Once a possible subarray is found set i to that index
do a binary search with the elements in the subarray

*/
