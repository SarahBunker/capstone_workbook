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
	println("simple merge")
	s1 := []int{1, 3, 5}
	l1 := createList(s1)
	n1 := []int{2, 4, 6}
	l2 := createList(n1)
	l3 := mergeTwoLists(l1, l2)
	printList(l3)

	println("first: 1 1 2 3 4 4")
	s1 = []int{1, 2, 4}
	l1 = createList(s1)
	n1 = []int{1, 3, 4}
	l2 = createList(n1)
	l3 = mergeTwoLists(l1, l2)
	printList(l3)

	println("empty")
	l1 = &ListNode{}
	l2 = &ListNode{}
	l3 = mergeTwoLists(l1, l2)
	printList(l3)

	println("first list empty")
	l1 = &ListNode{}
	n1 = []int{1, 3, 4}
	l2 = createList(n1)
	l3 = mergeTwoLists(l1, l2)
	printList(l3)

	println("second list empty")
	s1 = []int{1, 2, 4}
	l1 = createList(s1)
	l2 = &ListNode{}
	l3 = mergeTwoLists(l1, l2)
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
	if list1 == nil {
		c.Next = list2
		return dummy.Next
	}

	if list2 == nil {
		c.Next = list1
		return dummy.Next
	}

	for list1 != nil || list2 != nil {
		if list1.Val <= list2.Val {
			c.Next = list1
			list1 = list1.Next
			c = c.Next
		}

		if list1 == nil {
			c.Next = list2
			return dummy.Next
		}

		if list2.Val < list1.Val {

			c.Next = list2
			list2 = list2.Next
			c = c.Next
		}

		if list2 == nil {
			c.Next = list1
			return dummy.Next
		}

	}
	return dummy.Next
}
