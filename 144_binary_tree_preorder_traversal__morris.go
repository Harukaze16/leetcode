func preorderTraversal(root *TreeNode) []int {
    var res []int
    cur := root
    for cur != nil {
        if(cur.Left==nil) {
            res = append(res,cur.Val)
            cur = cur.Right
        } else {
            leftRightMost := cur.Left
            for leftRightMost.Right != nil && leftRightMost.Right!=cur {
                leftRightMost = leftRightMost.Right
            }
            if leftRightMost.Right == nil {
                res = append(res,cur.Val)
                leftRightMost.Right = cur
                cur = cur.Left
            } else {
                leftRightMost.Right = nil
                cur = cur.Right
            }
        }           
    }
    return res
}