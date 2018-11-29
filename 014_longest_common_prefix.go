func longestCommonPrefix(strs []string) string {
    if len(strs)<=0 {
    	return ""
    }
    cur_len := 0
    max_len := int((^uint(0)>>1))
    for _,cur_str := range strs {
        if (len(cur_str)<max_len) {
            max_len = len(cur_str) 
        }
    }
    sz := len(strs)
    for ;cur_len<max_len;cur_len++{
    	ch := strs[0][cur_len]
    	for i :=0;i<sz;i++ {
            if(strs[i][cur_len]!=ch) {
    			return strs[0][0:cur_len]
            }
    	}
    }
    return strs[0][0:max_len]
}

func min(a int,b int) int {
	if a<=b {
		return a
	}
	return b
}