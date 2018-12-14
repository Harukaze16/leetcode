func search(nums []int, target int) int { 
	beg,end := 0,len(nums)
	for beg<end {
        mid := (beg+end)/2
		if nums[mid]==target {
			return mid
		} else if (nums[beg]<nums[mid]) {
			if nums[mid] > target && nums[beg]<=target {
				end = mid
			} else {
				beg = mid+1
			}
		} else {
			if nums[mid] < target && nums[end-1]>=target {
				beg = mid+1
			} else {
				end = mid
			}
		}
	}
    return -1
}

