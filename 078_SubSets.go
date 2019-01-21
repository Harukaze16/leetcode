func subsets(nums []int) [][]int {
    var res [][]int
    recur_subsets(nums,&res,[]int{},0,len(nums))
    return res
}

func recur_subsets(nums []int, res *[][] int,cur []int,index int,sz int) {
	temp := make([]int,len(cur))
	copy(temp,cur)
	*res = append(*res,temp)
	for i:=index;i<sz;i++ {
		cur = append(cur,nums[i])
		recur_subsets(nums,res,cur,i+1,sz)
		cur = cur[:len(cur)-1]
	}
}