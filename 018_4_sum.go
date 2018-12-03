func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    if len(nums)<4 {
    	return [][]int{}
    }
    var res [][]int
    sz := len(nums)
    fmt.Println(nums)
    for i:=0;i<sz-3; {
    	target1 := target-nums[i]
    	for j:=i+1;j<sz-2;{
    		target2 := target1 - nums[j]
            for k,m:= j+1,sz-1;m>k;{
    			if(nums[k]+nums[m]<target2) {
    				k++
                } else {
                    if(nums[k]+nums[m]==target2 && (m==sz-1 || nums[m]!=nums[m+1])){
    			        res=append(res,[]int{nums[i],nums[j],nums[k],nums[m]})
                    // important [k,m must cannot be repeated
                    }
                    m--
                }
    		}
    		for j++;j<sz && nums[j]==nums[j-1];j++{
    		}
    	}
    	for i++;i<sz && nums[i-1]==nums[i];i++{
    	}
    }
    return res
}