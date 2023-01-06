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
	// printList(l)
	fmt.Println(deleteDuplicates(l))
	printList(l)

	s = []int{1, 1, 1, 1, 1, 2, 3, 3, 4, 4, 5}
	l = createList(s)
	fmt.Println(deleteDuplicates(l))
	printList(l)

	s = []int{1, 1, 1, 1, 1, 2, 3, 3, 4, 4, 5, 5, 5}
	l = createList(s)
	fmt.Println(deleteDuplicates(l))
	printList(l)
}
