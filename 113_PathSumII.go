func pathSum(root *TreeNode, sum int) [][]int {
    var res [][]int
    recurSum(root,sum,[]int{},&res)
    return res
}

func recurSum(root *TreeNode,sum int,cur []int, res *([][]int)) {
	if root !=nil {
		cur = append(cur,root.Val)
		if root.Left == nil && root.Right == nil {
			if sum == root.Val {
				temp := make([]int,len(cur))
				copy(temp,cur)
				(*res) = append(*res,temp)
			}
		} else {
			recurSum(root.Left,sum-root.Val,cur,res)
			recurSum(root.Right,sum-root.Val,cur,res)
		}
	}
}