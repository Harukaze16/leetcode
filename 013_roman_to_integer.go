func romanToInt(s string) int {
    dict := map[byte]int{
        'I':1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    res,pre := 0,0
	for i,sz := 0,len(s);i<sz;i++ {
        //通过存储pre的办法，与赋值为val的方式，避免了每次都重新从map中获取，从字符串中获取，减少runtime
        val := dict[s[i]] 
        if val>pre {
            res-=pre*2
        }
	    res += val
        pre = val
    }
	return res
}