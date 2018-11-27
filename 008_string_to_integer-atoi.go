func myAtoi(str string) int {
     src_str := []byte(str)
    
    const int_max int = int(int32((^uint32(0))>>1))
    const int_min int = ^(int(int32((^uint32(0))>>1)))
    //filter the blank space
    index,res,sign,length := 0,int64(0),1,len(src_str)
    for index < length && src_str[index]==' ' {
    	index++
    }
    if index <length && (src_str[index]=='-' || src_str[index]=='+'){
        if src_str[index]=='-' {
    	    sign = -1
        }
    	index++
    }
    for index <length && is_digit(src_str[index]) && (res>>31)==int64(0) { //notice it's must be == 
        res = res*10 + int64(src_str[index]-'0')
    	index++
    }
    if (res>>31) != int64(0) {
    	if sign == 1 {
    		return int_max
    	} else {
    		return int_min
    	}
    }
    return int(res*int64(sign))
}

func is_digit(cur byte) bool {
    if cur >= '0' && cur <='9'{
        return true
    }
    return false
}