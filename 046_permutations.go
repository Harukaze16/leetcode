func permute(nums []int) [][]int {
    var res [][]int
    recur_permute(nums,&res,len(nums),0)
    return res
}
func recur_permute(nums []int,res *[][]int,len int,cur int){
	if cur==len {
		temp := make([]int,len)
		copy(temp,nums)
		*res = append(*res,temp)
		return 
	}
    for i:=cur;i<len;i++ {
		swap(&nums[i],&nums[cur])
		recur_permute(nums,res,len,cur+1)
		swap(&nums[i],&nums[cur])
	}
}

func swap(a *int,b * int) {
	temp := *a
	*a = *b
	*b = temp
}