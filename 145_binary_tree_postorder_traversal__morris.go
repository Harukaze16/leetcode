func postorderTraversal(root *TreeNode) []int {
    var res []int
    head := TreeNode{0,root,nil}
    cur := &head
    for cur != nil {
        if cur.Left == nil {
            cur = cur.Right
        } else {
            leftRightMost := cur.Left
            for leftRightMost.Right!=nil && leftRightMost.Right!=cur {
                leftRightMost=leftRightMost.Right
            }
            if leftRightMost.Right==nil {
                leftRightMost.Right=cur
                cur = cur.Left
            } else {
                leftRightMost.Right = nil
                temp := reverse(cur.Left);
                for newLeft := temp;newLeft!=nil;newLeft = newLeft.Right {
                    res = append(res,newLeft.Val)
                }
                reverse(temp)
                cur = cur.Right
            }
        }
    }
    return res
}
func reverse(cur *TreeNode) *TreeNode {
    var  last_node *TreeNode = nil
    for cur!=nil {
        next := cur.Right
        cur.Right = last_node
        last_node = cur
        cur = next
    }
    return last_node
}