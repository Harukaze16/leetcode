func rotate(matrix [][]int)  {
    n := len(matrix)
    for i:=0;i<n;i++ {
        for j:=i+1;j<n;j++ {
            swap(&matrix[i][j],&matrix[j][i])
        }
    }

    for r:=0;r<n;r++ {
        for i,j:=0,n-1;i<j;i,j = i+1,j-1 {
            swap(&matrix[r][i],&matrix[r][j])
        }
    }

}

func swap(a *int,b *int) {
    temp := *b
    *b = *a
    *a = temp
}