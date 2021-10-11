/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    var left, right, longest int

    left = dive(root.Left, &longest)
    right = dive(root.Right, &longest)

    if longest > (left + right) {
        return longest
    }
    
    return left + right
}

func dive(node *TreeNode, longest *int) int {
    if node == nil { return 0 }
    var left, right int = dive(node.Left, longest), dive(node.Right, longest)

    if *longest < (left + right) {
        *longest = left + right
    }

    if left > right {
        return left + 1
    }

    return right + 1
}
