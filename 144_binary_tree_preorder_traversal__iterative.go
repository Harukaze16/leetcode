func preorderTraversal(root *TreeNode) []int {
    var res []int
    var stack []*TreeNode
    stack=append(stack,root)
    for len(stack)>0 {
        cur := stack[len(stack)-1]
        stack=stack[:len(stack)-1]
        
        if cur == nil {
            continue
        }
        res=append(res,cur.Val)
        stack=append(stack,cur.Right)
        stack=append(stack,cur.Left)
    }
    
    return res
}