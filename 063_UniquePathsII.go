//o(M*N) space
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if len(obstacleGrid)==0 || len(obstacleGrid[0])==0 {
        return 0
    }
    rows,cols := len(obstacleGrid),len(obstacleGrid[0])
    
    count := make([][]int,rows+1)
    for i:=0;i<rows+1;i++ { //notice this should be rows+1
        count[i] = make([]int,cols+1)
    }
    
    //set 1
    count[rows-1][cols] = 1
    
    for c := cols-1;c>=0;c-- {
        for r:=rows-1;r>=0;r-- {
            if obstacleGrid[r][c] == 1 {
                count[r][c]  = 0
            } else {
                count[r][c] = count[r+1][c]+count[r][c+1]
            }
        }
    }
    return count[0][0]
}

//o(N) space
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if len(obstacleGrid)==0 || len(obstacleGrid[0])==0 {
        return 0
    }
    rows,cols := len(obstacleGrid),len(obstacleGrid[0])
    
    count := make([]int,rows+1)

    count[rows-1] = 1
    for c := cols-1;c>=0;c-- {
        for r:=rows-1;r>=0;r-- {
            if obstacleGrid[r][c] == 1 {
                count[r]  = 0
            } else {
                count[r] = count[r]+count[r+1]
            }
        }
    }
    return count[0]
}