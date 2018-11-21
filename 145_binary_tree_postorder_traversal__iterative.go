func postorderTraversal(root *TreeNode) []int {
    var res []int
   
    //判断上个节点是否是当前孩子的右节点以此确定当前节点是否右孩子节点已遍历过
    var stack []*TreeNode
    cur:=root
    var last_node *TreeNode
    for len(stack)>0 || cur!=nil {
        for cur!=nil {
            stack=append(stack,cur)
            cur=cur.Left
        }
        cur = stack[len(stack)-1]
        if cur.Right==last_node || cur.Right==nil {
            res=append(res,cur.Val)
            // 出栈
            stack=stack[:len(stack)-1]
            last_node = cur
            cur=nil
        } else { //处理右子树
            cur=cur.Right
        }
    }
    return res
}