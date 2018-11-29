func threeSum(nums []int) [][]int {
    var res [][]int
    sort.Ints(nums)
    for i,sz:=0,len(nums);i<sz;{
    	target := -nums[i]
        for j,k :=i+1,sz-1;j<k;{
    		if(nums[j]+nums[k]<target){
    			for j++;j<k&& nums[j-1]==nums[j];j++{
    			}
    		} else {
    			if(nums[j]+nums[k]==target) {
    				cur := []int{nums[i],nums[j],nums[k]}
    				res = append(res,cur)
    			} 
    			for k--;j<k && nums[k]==nums[k+1];k--{ 
                //  其实仍可简化代码，比如只需要在==target时，才迫切必须需要检验是否想等  				
    			}
    		}
    	}
        for i++;i<sz && nums[i]==nums[i-1];i++{
        }
    }
    return res
}