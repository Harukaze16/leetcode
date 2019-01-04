func solveNQueens(n int) [][]string {
	if n==0 {
		return [][]string{}
	}
    var res [][]string
    
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
    recur_solve(n,0,&res,&cur,&colExi,&gridExi,&gridExi2)
    
    return res
}

// in i row
func recur_solve(n int,i int, res* [][]string,cur* [][]byte,colExi* []bool,gridExi* []bool,gridExi2 * []bool) {
	if n==i {
		cur_str := make([]string,n)
		for i:=0;i<n;i++ {
            cur_str[i]= string((*cur)[i])
		}
		*res = append(*res,cur_str)
	}

	for j:=0;j<n;j++ {
        if !(*colExi)[j] && !(*gridExi)[j-i+n] && !(*gridExi2)[i+j] {
			(*colExi)[j]=true
            (*gridExi)[j-i+n]=true
            (*gridExi2)[i+j]=true

			(*cur)[i][j]='Q'
			recur_solve(n,i+1,res,cur,colExi,gridExi,gridExi2)
            
            (*cur)[i][j]='.'
            (*colExi)[j]=false
            (*gridExi)[j-i+n]=false
            (*gridExi2)[i+j]=false
		} 
	}
}