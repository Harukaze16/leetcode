func divide(dividend int, divisor int) int {
    minus := false
    if dividend^divisor < 0 {
        minus = true
    }
    dividend,divisor = abs(dividend),abs(divisor)
    res := 0
    
    cur_shifts,cur_num := uint(0),divisor
    for dividend>=(cur_num<<1) {
        cur_num = cur_num<<1
        cur_shifts++
    }
    for dividend>=divisor {
        if dividend>=cur_num {
            dividend-=cur_num
            res+=(1<<cur_shifts)
        }
        cur_shifts--
        cur_num =cur_num>>1
    }
    if minus {
        res=-res
    } else if res>math.MaxInt32 {
        return math.MaxInt32
    }
    return res
}

func abs(num int) int {
    if(num<0) {
        return -num
    } else {
        return num
    }
}