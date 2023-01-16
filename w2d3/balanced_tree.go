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

func isBalanced(root *TreeNode) bool {
	_, balanced := balanceDepth(root)
	return balanced
}

func balanceDepth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	left, balanced := balanceDepth(root.Left)
	if !balanced {
		return 0, false
	}
	right, balanced := balanceDepth(root.Right)
	if !balanced {
		return 0, false
	}
	if left-right > 1 || right-left > 1 {
		return 0, false
	}
	return max(left, right) + 1, true
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
