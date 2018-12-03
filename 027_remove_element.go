func removeElement(nums []int, val int) int {
	i,sz := 0,len(nums)
	for j:= 0;j<sz;j++ {
		if(nums[j]!=val) {
			nums[i]=nums[j]
            i++
		}
	}
	return i;
}