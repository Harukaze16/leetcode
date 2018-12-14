func strStr(haystack string, needle string) int {
    index,need_i,needle_len,sz  := 0,0,len(needle),len(haystack)
    //考虑needle长度>haystack的情况，不特殊处理会导致后续比较时越界
    if(needle_len>sz) {
        return -1
    }
    for index = 0;index<sz-needle_len+1;index++ {   //notice this scope ,可能会导致越界
        for need_i = 0;need_i<needle_len && haystack[index+need_i]==needle[need_i];need_i++ {
        }
        if(need_i == needle_len) {
            return index
        }
    }
    return -1
}