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
	s := []int{1, 2, 3, 4, 5}
	l := createList(s)
	n := 2
	fmt.Println(removeNthFromEnd(l, n))
	printList(l)

	s = []int{1, 2, 3, 4, 5}
	l = createList(s)
	n = 1
	fmt.Println(removeNthFromEnd(l, n))
	printList(l)

	s = []int{1, 2}
	l = createList(s)
	n = 1
	fmt.Println(removeNthFromEnd(l, n))
	printList(l)

	s = []int{1, 2}
	l = createList(s)
	n = 2
	fmt.Println(removeNthFromEnd(l, n))
	printList(l)

	s = []int{1}
	l = createList(s)
	n = 1
	fmt.Println(removeNthFromEnd(l, n))
	printList(l)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	hsh := map[int]*ListNode{}
	count := 0
	c := head
	for i := 0; c != nil; i++ {
		count++
		hsh[i] = c
		c = c.Next
	}

	beforeNode := hsh[count-n-1]
	afterNode := hsh[count-n+1]
	fmt.Println(beforeNode, afterNode)
	if beforeNode == nil && afterNode == nil {
		return nil
	}

	if beforeNode == nil {
		return afterNode
	}

	if afterNode == nil {
		beforeNode.Next = nil
		return head
	}

	beforeNode.Next = afterNode
	return head
}
