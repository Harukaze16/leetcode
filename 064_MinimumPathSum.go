func minPathSum(grid [][]int) int {
    if len(grid)==0 || len(grid[0])==0 {
    	return 0
    }

    rows,cols := len(grid),len(grid[0])

    //notice 尾列的特殊处理
    for r:=rows-2;r>=0;r--{
        grid[r][cols-1]+=grid[r+1][cols-1]
    }
    for c:=cols-2;c>=0;c-- {
        grid[rows-1][c] += grid[rows-1][c+1]
        for r:=rows-2;r>=0;r-- {
            grid[r][c] += min(grid[r+1][c],grid[r][c+1])
        }
    }
    return grid[0][0]
}

func min(a,b int) int {
	if a<=b {
		return a
	}
	return b
}