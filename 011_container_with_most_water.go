func maxArea(height []int) int {
    if len(height)==0 {
        return 0
    }
    i,j := 0,len(height)-1
    res := min(height[i],height[j])*(j-i)
    for j>i {
        // 循环内结构此种优化很棒，只要<原先min_num,即应迭代到下一项
        min_num := min(height[i],height[j])
        for i<j && height[i]<=min_num {
            i++
        }
        for i<j && height[j]<=min_num {
            j--
        }
        res = max(res,min(height[i],height[j])*(j-i))
    }
    return res
}



func min(a int,b int) int {
    if a<=b {
        return a
    }
    return b
}

func max(a int,b int) int {
    if a>=b {
        return a
    }
    return b
}