func numDecodings(s string) int {
    last1,last2 := 1,1
    for i:=len(s)-1;i>=0;i-- {
        if s[i]=='0' {
            last1,last2 = 0,last1
        } else {
            temp := last1
            if (i+1<len(s))  && (s[i]<'2' || s[i]=='2' && s[i+1]<='6') {
                temp+=last2
            }
            last1,last2 = temp,last1
        }
    }
    return last1
}