package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println("Hello")
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	left, right := rangeSumBST(root.Left, low, high), rangeSumBST(root.Right, low, high)
	if root.Val <= high && root.Val >= low {
		return root.Val + left + right
	}
	if root.Val > high {
		return left
	}
	if root.Val < low {
		return right
	}
	return 0
}

/*
base
root is nil
return 0

if cur val is within range
return cur Val + rangeSumBST(root.Left) + rangeSumBST(root.Left)
if cur val is greater then range
return rangeSumBST(root.Left)
if cur val is less then range
return rangeSumBST(root.Right)
*/
