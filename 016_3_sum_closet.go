func threeSumClosest(nums []int, target int) int {
    if len(nums)<3 {
    	return 0
    }
    res := nums[0]+nums[1]+nums[2]
    sort.Ints(nums)

    //find the closet num
    for i,sz:=0,len(nums);i<sz;{
        for j,k :=i+1,sz-1;j<k;{
    		cur := nums[i]+nums[j]+nums[k]
    		if(abs(cur-target)<abs(res-target)) {
    			res = cur
    		}
            if(cur<target){
                j++
            } else {
                k--;
            }
    	}
    	for i++;i<sz && nums[i]==nums[i-1];i++{
    	}
    }
    return res;
}

func abs(a int) int {
	if(a<0) {
		return -a;
	}
	return a;
}