func totalNQueens(n int) int {
    cur := make([][]byte,n)
    for i:=0;i<n;i++ {
        cur[i]=make([]byte,n)
        for j:=0;j<n;j++ {
            cur[i][j]='.'
        }
    }

    colExi := make([]bool,n)
    gridExi := make([]bool,2*n)
    gridExi2 := make([]bool,2*n)
    return recur_solve(n,0,&cur,&colExi,&gridExi,&gridExi2)
}
// in i row
func recur_solve(n int,i int, cur* [][]byte,colExi* []bool,gridExi* []bool,gridExi2 * []bool) int {
    cur_num := 0
	if n==i {
        return cur_num+1
	}

	for j:=0;j<n;j++ {
        if !(*colExi)[j] && !(*gridExi)[j-i+n] && !(*gridExi2)[i+j] {
			(*colExi)[j]=true
            (*gridExi)[j-i+n]=true
            (*gridExi2)[i+j]=true

			(*cur)[i][j]='Q'
			cur_num+=recur_solve(n,i+1,cur,colExi,gridExi,gridExi2)
            
            (*cur)[i][j]='.'
            (*colExi)[j]=false
            (*gridExi)[j-i+n]=false
            (*gridExi2)[i+j]=false
		} 
	}
    return cur_num
}