func minDistance(word1 string, word2 string) int {
    sz1,sz2 := len(word1),len(word2)
    dp := make([][]int,sz1+1)
    for i:=0;i<=sz1;i++ {
    	dp[i] = make([]int,sz2+1)
    }

    for i:=0;i<=sz1;i++ {
    	dp[i][0]=i
    }
    for i:=0;i<=sz2;i++ {
    	dp[0][i]=i
    }
    for i:=1;i<=sz1;i++ {
        for j:=1;j<=sz2;j++ {
    		if word1[i-1]==word2[j-1] {
    			dp[i][j] = dp[i-1][j-1]
    		} else {
    			dp[i][j] = min(dp[i-1][j-1],dp[i-1][j],dp[i][j-1]) + 1
    		}
    	}
    }
    return dp[sz1][sz2]
}

func min(a int,b int,c int) int{
	if a<=b && a<=c {
		return a
	} else if b<=a && b<= c {
		return b
	} else {
		return c
	}
}



//o(n) space solution
func minDistance(word1 string, word2 string) int {
    sz1,sz2 := len(word1), len(word2)
    dp := make([]int,sz2+1)
    for i:=0;i<=sz2;i++ {
        dp[i] = i
    }

    for i:=1;i<=sz1;i++ {
        last := dp[0]
        dp[0] = i
        for j:=1;j<=sz2;j++ {
            if word1[i-1] == word2[j-1] {
                temp := dp[j]
                dp[j] = last
                last = temp
            } else {
                temp := dp[j]
                dp[j] = min(dp[j],last,dp[j-1])+1
                last = temp
            }
        }
    }
    return dp[sz2]
}

func min(a, b ,c int) int{
    if a<=b && a<=c {
        return a
    } else if b<=a && b<= c {
        return b
    } else {
        return c
    }
}