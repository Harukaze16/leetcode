func trap(height []int) int {
    var res int
    l,r,maxl,maxr:=0,len(height)-1,0,0
    for l<=r {
        if maxl<=maxr {
            if height[l]>=maxl {
                maxl=height[l]
            } else {
                res += (maxl-height[l])
            }
            l++
        } else {
            if height[r]>=maxr {
                maxr = height[r]
            } else {
                res += (maxr-height[r])
            }
            r--
        }
    }
    return res
}