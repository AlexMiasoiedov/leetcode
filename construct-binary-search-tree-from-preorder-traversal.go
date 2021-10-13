/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    if len(preorder) < 1 { return nil }

    var node TreeNode = TreeNode{Val: preorder[0]}

    var startLeft, startRight, end int
    startLeft = 0
    end = len(preorder)

    for startRight = startLeft; startRight < end; startRight++ {
        if preorder[startLeft] < preorder[startRight] { break }
    }

    node.Left = bstFromPreorder(preorder[startLeft+1:startRight])
    node.Right = bstFromPreorder(preorder[startRight:])

    return &node
}
