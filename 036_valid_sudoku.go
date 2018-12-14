func isValidSudoku(board [][]byte) bool {
    //check rows
    for i:=0;i<9;i++ {
        dict := make(map[byte]bool)
    	for j:=0;j<9;j++ {
            if board[i][j]!='.' {
                if dict[board[i][j]]!=false {
                    return false
                } else {
                    dict[board[i][j]]=true
                }
            }
        }
    }
    //check columns
    for j:=0;j<9;j++ {
        dict := make(map[byte]bool)
    	for i:=0;i<9;i++ {
            if board[i][j]!='.' {
                if dict[board[i][j]]!=false {
                    return false
                } else {
                    dict[board[i][j]]=true
                }
            }
        }
    }
    //check grid
    for i:=0;i<9;i++ {
    	m,n := (i/3)*3,(i%3)*3
        dict := make(map[byte]bool)
    	for k:=0;k<3;k++ {
    		for v :=0 ;v<3;v++ {
                if board[m+k][n+v]!='.' {
                    if dict[board[m+k][n+v]]!=false{
                        return false
                    } else {
                        dict[board[m+k][n+v]] =true
                    }
                }
            }
        }
    }
    return true 
}