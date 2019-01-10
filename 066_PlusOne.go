func plusOne(digits []int) []int {
    res := make([]int,len(digits)+1)
    res[0]= 1
    copy(res[1:],digits)
    for i:=len(digits);i>=1;i--{
        res[i]=(res[i]+1)%10
        if res[i]!=0 {
            return res[1:]
        }
    }
    return res
}