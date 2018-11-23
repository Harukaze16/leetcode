func lengthOfLongestSubstring(s string) int {
    array := []byte(s)
    dict := make(map[byte]int)
	begin_index,max_len:=  0,0
    for index,c := range(array) {
		old_index,ok := dict[c]
		if ok && old_index>=begin_index { // 此值在当前序列出现过
			max_len = max(max_len,index-begin_index)
			begin_index = old_index+1;
        }
        dict[c]=index
	}
	return max(max_len,len(s)-begin_index)
}

func max(a int,b int) int {
    if a>=b {
        return a
    }
    return b
}