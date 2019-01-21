func search(nums []int, target int) bool {
    beg,end := 0,len(nums)
    for beg<end {
        mid := (beg+end)/2
    	if nums[mid] == target {
    		return true
    	} else if nums[mid]==nums[beg] && nums[mid]==nums[end-1] {
    		beg,end = beg+1,end-1
    	} else if nums[mid]>=nums[beg] { //前半段和后半段必有一个是有序的
    		if nums[beg]<=target && nums[mid]>target {
    			end =mid
    		} else {
    			beg = mid+1 
    		}
    	} else {
    		if nums[mid]<target && nums[end-1]>=target { //notice >=
    			beg = mid+1
    		} else {
    			end = mid;
    		}
    	}
    }
    return false
}