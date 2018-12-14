func searchRange(nums []int, target int) []int {
    beg,end := 0,len(nums)

    // find range begin

    mid,bege,ende := 0,0,0
    for beg<end {
    	mid = (beg+end)/2
    	if nums[mid]<target {
    		beg = mid+1
    	} else if nums[mid]>target {
    		end=mid
    	} else {
    		if mid==beg || nums[mid-1]!=target {
    			break 
    		} else {
    			end = mid
    		}
    	}
    }
    bege = mid
    
    if beg == end {
    	return []int{-1,-1}
    }
    //find range end
    beg,end = 0,len(nums)
    for beg<end {
    	mid=(beg+end)/2
		if nums[mid]<target {
    		beg = mid+1
    	} else if nums[mid]>target {
    		end = mid
    	} else {
    		if mid==end-1 || nums[mid+1]!=target {
    			break 
    		} else {
    			beg = mid+1
    		}
    	}
    }
    ende = mid
    return []int{bege,ende}
}