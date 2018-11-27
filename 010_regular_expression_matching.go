func isMatch(s string, p string) bool {
    rsz,csz := len(p),len(s)
    dp := make([][]bool,rsz+1)
    for i:= 0;i<rsz+1;i++ {
        dp[i] = make([]bool,csz+1)
    }
    dp[rsz][csz]=true
    
    //notice this,以处理pattern末尾为*时，对于其匹配到s末尾的初始值赋值，it's important
    for i :=rsz-1;i>0 && p[i]=='*';i-=2{
            dp[i-1][csz]=true
    }
    // dp process
    for j :=csz-1;j>=0;j--{ //column
        for i:=rsz-1;i>=0;i--{ //row
            if(p[i]=='*') {
                continue
            }
            if (i+1<rsz && p[i+1]=='*') {
                if p[i]==s[j] || p[i]=='.' {
                    dp[i][j]=dp[i+2][j]||dp[i+2][j+1]||dp[i][j+1]
                } else {
                    dp[i][j]=dp[i+2][j]
                }
            } else {
                dp[i][j]=(p[i]==s[j] || p[i]=='.') && dp[i+1][j+1]
            }
        }
    }
    return dp[0][0]
}