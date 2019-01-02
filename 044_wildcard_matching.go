func isMatch(s string, p string) bool {
	plen,slen := len(p),len(s)
	dp := make([][]bool,plen+1)
	for i:=0;i<=plen;i++ {
        dp[i] =make([]bool,slen+1)
	}
	dp[plen][slen]=true
	for i:=plen-1;i>=0 && p[i]=='*';i-- {
		dp[i][slen]=true
	}

	//dp solution
	for i:=plen-1;i>=0;i-- {
		for j:=slen-1;j>=0;j--{
			if p[i]=='*' {
				dp[i][j]=dp[i+1][j]||dp[i][j+1]
			} else if p[i]=='?' || p[i]==s[j] {
				dp[i][j]=dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}