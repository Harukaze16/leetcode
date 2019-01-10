/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return recurValid(root,0,0,false,false)
}

func recurValid(root *TreeNode,min int,max int,minValid bool,maxValid bool) bool{
    if root==nil {
        return true
    } else if minValid && root.Val<=min || maxValid && root.Val>=max {
        return false
    } else {
        return recurValid(root.Left,min,root.Val,minValid,true) && recurValid(root.Right,root.Val,max,true,maxValid)
    }
}