func uniquePaths(m int, n int) int {
    count := make([][]int,n)
    for i:=0;i<n;i++ {
        count[i] = make([]int,m)
    }
    for i:=0;i<m;i++ {
        count[n-1][i]=1
    }
    for i:=0;i<n;i++ {
        count[i][m-1]=1
    }
    
    for c:=m-2;c>=0;c-- {
        for r:=n-2;r>=0;r-- {
            count[r][c]=count[r][c+1]+count[r+1][c]
        }
    }
    return count[0][0]
}


//排列组合
func uniquePaths(m int, n int) int {
    m,n = m-1,n-1
    if m>n {
        swap(&m,&n)
    }
    res := 1
    for i:=1;i<=m;i++ {
        res = res *(n+i)/i
    }
    return res
}

func swap(a,b *int){
    temp := *a
    *a = *b
    *b = temp
}