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
	s1 := []int{1, 2, 3, 4, 5}
	l1 := createList(s1)
	n1 := []int{1, 2, 3, 4, 5}
	l2 := createList(n1)
	l3 := mergeTwoLists(l1, l2)
	printList(l3)

	// s = []int{1, 2, 3, 4, 5}
	// l = createList(s)
	// n = 1
	// fmt.Println(mergeTwoLists(l, n))

	// s = []int{1, 2}
	// l = createList(s)
	// n = 1
	// fmt.Println(mergeTwoLists(l, n))

	// s = []int{1}
	// l = createList(s)
	// n = 1
	// fmt.Println(mergeTwoLists(l, n))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	c := dummy
	for c != nil {
		if list1.Val <= list2.Val {
			c.Next = list1
			list1 = list1.Next
			c = c.Next
		}

		if list2.Val < list1.Val {
			c.Next = list2
			list2 = list2.Next
			c = c.Next
		}

	}
	return dummy.Next
}
