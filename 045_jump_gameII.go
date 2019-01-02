func jump(nums []int) int {
    for i,maxR,level,maxRCurr := 0,0,0,0;level<len(nums)-1;{
 //       fmt.Print(maxRCurr)
        level++
        for ;i<=maxR;i++ {
            maxRCurr = max(maxRCurr,nums[i]+i)
            if maxRCurr>=len(nums)-1 {
                return level
            }
        }
//        level++
        maxR = maxRCurr
    }
    return 0
}

func max(a int,b int) int{
    if a>=b {
        return a
    } else {
        return b
    }
}