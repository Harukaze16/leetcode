func reverse(x int) int {
    if x==0 {
        return 0
    }
    var res int64
    sign := 1
    if(x<0) {
        x=-x
        sign = -1
    }
    for x%10==0 {
        x/=10
    }
    for x>0 {
        res =res*10+(int64)(x%10)
        x/=10
    }
    // this condition perfect
    if (res>>31) !=0 {
        return 0
    }
    true_res := res*int64(sign)
    return int(true_res)
}