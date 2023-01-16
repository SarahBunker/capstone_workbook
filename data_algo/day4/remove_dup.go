// To execute Go code, please declare a func main() in a package "main"
// https://leetcode.com/problems/remove-duplicates-from-sorted-list

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

/*
currNode = head
next = currNode.next

while next != nil
		while next.next == currNode
			reassign next to next.next
	Assign current.next to next
	current.next = next
	next = next.next
return head
*/

func deleteDuplicates(head *ListNode) *ListNode {
	currNode := head
	nextNode := currNode.Next
	for nextNode != nil {
		for currNode.Val == nextNode.Val {
			nextNode = nextNode.Next
			if nextNode.Next == nil {
				break
			}

		}
		currNode.Next = nextNode
		currNode = nextNode
		nextNode = nextNode.Next
	}
	currNode.Next = nil
	return head
}

// {1, 1, 1, 1, 1, 2,3,3,4,4,5}
// H
//                 N
// A
