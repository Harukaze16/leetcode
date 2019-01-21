func exist(board [][]byte, word string) bool {
    if len(board)==0 || len(board[0])==0 || len(word)==0{
        return false
    }
    rows,cols := len(board),len(board[0])
    for i:=0;i<rows;i++ {
        for j:=0;j<cols;j++ {
            if board[i][j] == word[0] {
                if recur_find(board,word,1,rows,cols,i,j) {
                 	return true
            	}
            }
        }
    }
    return false
}

//find if it's adjacent have wanted element
func recur_find(board [][]byte,word string,index int,rows int,cols int,i int,j int) bool{
    if index>= len(word){
		return true
	}
    temp := board[i][j]
	board[i][j] = ' '
	if i-1>=0 && board[i-1][j]==word[index] && 
		recur_find(board,word,index+1,rows,cols,i-1,j) {
			return true
	} else if i+1<rows && board[i+1][j]==word[index] && 
		recur_find(board,word,index+1,rows,cols,i+1,j) {
			return true
	} else if j-1>=0 && board[i][j-1]==word[index] && 
		recur_find(board,word,index+1,rows,cols,i,j-1) {
			return true
	} else if j+1<cols && board[i][j+1]==word[index] && 
		recur_find(board,word,index+1,rows,cols,i,j+1) {
			return true
	} else {
		board[i][j] = temp
		return false
	}
}