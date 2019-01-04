func maxSubArray(nums []int) int {
    max,cur:= nums[0],0
    for i:=0;i<len(nums);i++ { 
        if cur<0 {
            cur = 0
        }
        cur += nums[i]
        if cur > max {
            max = cur
        }
    }
    return max
}

