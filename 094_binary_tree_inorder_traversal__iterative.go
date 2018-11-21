func inorderTraversal(root *TreeNode) []int {
    var res []int
    var stack [] *TreeNode 
    var cur = root
    for cur!=nil || len(stack)>0 {
        
        //此循环只会增加stack内容，而不会减少stack元素数量    
        for cur!=nil {
            stack=append(stack,cur)
            cur=cur.Left
        }
        //栈顶元素出栈，输出
        cur = stack[len(stack)-1]
        stack=stack[:len(stack)-1]
        res = append(res,cur.Val)
        cur = cur.Right
    }
    return res
}