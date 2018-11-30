
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "math"

func longestConsecutive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := helper(root, 1) + helper(root, -1) + 1
	return int(math.Max(float64(res), math.Max(float64(longestConsecutive(root.Left)), float64(longestConsecutive(root.Right)))))
}

func helper(node *TreeNode, diff int) int {
	if node == nil {
		return 0
	}
	var left, right int
	if node.Left != nil && (node.Val-node.Left.Val) == diff {
		left = 1 + helper(node.Left, diff)
	}
	if node.Right != nil && (node.Val-node.Right.Val) == diff {
		right = 1 + helper(node.Right, diff)
	}
	return int(math.Max(float64(left), float64(right)))
}
