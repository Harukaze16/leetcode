/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    _,res := depthBalanced(root)
    return res
}


func depthBalanced(root *TreeNode) (int,bool) {
    if root==nil {
        return 0,true
    }
    leftDep,leftBalance := depthBalanced(root.Left)
    rightDep,rightBalance := depthBalanced(root.Right)
    return max(leftDep,rightDep)+1,leftBalance&&rightBalance && diff(leftDep,rightDep)<=1
}
func diff(a,b int) int {
    if a>=b {
        return a-b
    } else {
        return b-a
    }
}
func max(a,b int) int {
	if a>= b {
		return a
	} else {
		return b 
	}
}