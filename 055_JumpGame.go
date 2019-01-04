func canJump(nums []int) bool {
    sz := len(nums)
    if sz ==0 {
        return true
    }
    start,end:=0,0
    for start<=end {
        last_end := end
        for i:=start;i<=last_end;i++ {
            if nums[i]+i>=sz-1 {
                return true
            } else if nums[i]+i>end {
                end = nums[i]+i
            }
        }
        start = last_end+1
    }
    return false
}