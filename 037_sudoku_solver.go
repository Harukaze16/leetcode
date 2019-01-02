func solveSudoku(board [][]byte)  {
    var rowRec  [9][10]bool
    var colRec  [9][10]bool 
    var gridRec [9][10]bool

    emptyGrid := []pair{}
    for i:=0;i<9;i++ {
        for j:=0;j<9;j++ {
            if board[i][j]=='.' {
                emptyGrid = append(emptyGrid,pair{i,j})
            } else {
                rowRec[i][board[i][j]-'0'] = true
                colRec[j][board[i][j]-'0'] = true
                gridRec[(i/3)*3+j/3][board[i][j]-'0'] = true
            }
        }
    }
    solve(board,emptyGrid,rowRec,colRec,gridRec)
}


func solve(board [][]byte,wait []pair,rowRec [9][10]bool,colRec [9][10]bool,gridRec [9][10]bool) bool {
    if len(wait)==0 {
        return true
    }
    cur_pair:= wait[len(wait)-1]
    r,c := cur_pair.r,cur_pair.c
    for i:=byte(1);i<10;i++ {
        if !rowRec[r][i] && !colRec[c][i] && !gridRec[(r/3)*3+c/3][i] {
            board[r][c]=i+'0'
            rowRec[r][i] = true
            colRec[c][i] = true
            gridRec[r/3*3+c/3][i] = true

            if solve(board,wait[0:len(wait)-1],rowRec,colRec,gridRec) {
                return true
            }

            rowRec[r][i] = false
            colRec[c][i] = false
            gridRec[r/3*3+c/3][i] = false
            board[r][c]='.'
        }
    }
    return false
}
type pair struct{
    r,c int
}