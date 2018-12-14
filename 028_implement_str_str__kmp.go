func strStr(haystack string, needle string) int {
	need_sz,sz := len(needle),len(haystack)
	if(need_sz>sz) {
		return -1
	} else if(need_sz==0) {
		return 0
	}

    next := make([]int,len(needle))
    cur := 0
	for i:=2;i<need_sz;{
		if(needle[i-1]==needle[cur]) {
			next[i]=cur+1
            cur ++
            i++
        } else if cur>0{   //important
            cur = next[cur]
        } else {
            next[i]=0
            i++
        }
	}

	i,j:=0,0
	for i<sz && j<need_sz{
		if(haystack[i]==needle[j]) {
			i++
			j++
		} else if(j==0) {
			i++
		} else {
			j=next[j]
		}
	}
    if(j==need_sz) {
        return i-need_sz
    } else {
        return -1
    }
}