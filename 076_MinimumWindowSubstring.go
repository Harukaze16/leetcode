func minWindow(s string, t string) string {
    dict := make(map[byte]int)
    for i:= 0;i<len(t);i++ {
        dict[t[i]]++
    }
    
    minLen,sz,count := len(s)+1,len(s),len(t)
    var beg,end,head int
    for beg,end = 0,0; end<sz;end++ {
        if dict[s[end]]>0 {
            count--
        }
        dict[s[end]]--
        for count==0 {
            if end-beg+1 < minLen {
                minLen = end-beg+1
                head = beg
            }            
            if dict[s[beg]]==0 { //每次beg延后一个，判断一个beg
                count++
            }
            dict[s[beg]]++ //notice
            beg++
        }
    }
    if minLen>len(s) {
        return ""
    } else {
        return s[head:head+minLen]
    }
    
}