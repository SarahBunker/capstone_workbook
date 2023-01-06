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
	println("simple reverse")
	s1 := []int{1, 3, 5}
	l1 := createList(s1)
	l2 := reverseList(l1)
	printList(l2)

	println("short reverse")
	s1 = []int{1, 3}
	l1 = createList(s1)
	l2 = reverseList(l1)
	printList(l2)

	println("shortest reverse")
	// s1 = []int{1, 3}
	l1 = &ListNode{}
	l2 = reverseList(l1)
	printList(l2)

	println("empty")
	// s1 = []int{1, 3}
	l1 = nil
	l2 = reverseList(l1)
	printList(l2)

}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	var p *ListNode = nil
	c := head
	for c != nil {
		n := c.Next
		c.Next = p
		p = c
		c = n
	}
	return p
}
