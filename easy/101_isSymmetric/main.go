package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func compareList(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left != nil && right == nil) ||
		(left == nil && right != nil) ||
		left.Val != right.Val {
		return false
	}

	return compareList(left.Left, right.Right) &&
		compareList(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return compareList(root.Left, root.Right) &&
		compareList(root.Right, root.Left)
}

func main() {
	fmt.Print(isSymmetric(&TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}},
	}))
	fmt.Print(isSymmetric(&TreeNode{Val: 2,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
	}))
}
