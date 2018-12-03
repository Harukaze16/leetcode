func removeDuplicates(nums []int) int {
    i,sz := 0,len(nums)
    for j := 1;j<sz;j++ {
    	if(nums[j]!=nums[i]) {
    		i++
    		nums[i]=nums[j]
    	}
    }
    return i+1;
}