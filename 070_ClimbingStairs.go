func climbStairs(n int) int {
    last,cur := 1,1
    for i:=2;i<=n;i++ {
        temp := last
        last = cur
        cur = last+temp
    }
    return cur
}