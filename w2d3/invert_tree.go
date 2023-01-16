package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func invertTree(root *TreeNode) *TreeNode {
//   if root.Left == nil && root.Right == nil {
// 		return root
// 	}
// 	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
// 	return root
// }

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func main() {
	fmt.Println("Hello")
}

/*
base case
return ptr to root if Left and Right are nil

recursive
	Left, Right = inversionRight, inversionLeft
*/
