package problems

import "fmt"

func maxDepth(root *TreeNode) int {
	return calculateDepth(root, 0)
}

func calculateDepth(node *TreeNode, currentDepth int) int {
	if node == nil {
		return currentDepth
	}
	currentDepth++
	return Max(calculateDepth(node.Left, currentDepth), calculateDepth(node.Right, currentDepth))
}
func SolveMaxDepth() {
	fmt.Println("104. Maximum Depth of Binary Tree")
	fmt.Println("https://leetcode.com/problems/maximum-depth-of-binary-tree/")
	fmt.Println("---------------------------")

}
