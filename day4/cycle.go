package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func createList(nums []int) (head *ListNode) {
	head = &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		next := &ListNode{Val: nums[i]}
		current.Next = next
		current = next
	}
	return head
}

func main() {
	s := []int{1, 2, 3, 3, 4, 4, 5}
	l := createList(s)
	fmt.Println(hasCycle(l))

	s = []int{1, 1, 1, 1, 1, 2, 3, 3, 4, 4, 5}
	l = createList(s)
	fmt.Println(hasCycle(l))

	s = []int{1, 1, 1, 1, 1, 2, 3, 3, 4, 4, 5, 5, 5}
	l = createList(s)
	fmt.Println(hasCycle(l))
}

// func hasCycle(head *ListNode) bool {
// 	hsh := map[*ListNode]bool{}
// 	c := head
// 	for c != nil {
// 		_, ok := hsh[c]
// 		if ok {
// 			return true
// 		}
// 		hsh[c] = true
// 		c = c.Next
// 	}
// 	return false
// }

func hasCycle(head *ListNode) bool {
	s := head
	f := head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}
	return false
}

/*

Using boiler plate code for creating lists I was able to create some test cases
None of them are cycled lists though.

Iterate through list and if current node is already in list of nodes then return true
otherwise store node in hash with boolean true
increment current to next node.

*/
