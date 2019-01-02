func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    } 
    return isSym(root.Left,root.Right)
}
func isSym(root1 *TreeNode,root2 *TreeNode) bool{ 
    if root1==nil && root2==nil {
        return true
    } else if root1==nil || root2==nil {
        return false
    } else {
        return root1.Val==root2.Val && isSym(root1.Left,root2.Right) && isSym(root1.Right,root2.Left)
    }
}