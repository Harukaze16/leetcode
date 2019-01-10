func setZeroes(matrix [][]int)  {
    if len(matrix)==0 || len(matrix[0]) == 0 {
    	return
    }

    rows,cols :=len(matrix),len(matrix[0])
    var firLine0 bool
    //cope first line
    for i:=0;i<cols;i++ {
    	if matrix[0][i]==0 {
    		firLine0 = true
    		break
    	}
    }
    for i:=1;i<rows;i++ {
    	for j:=0;j<cols;j++ {
    		if matrix[i][j] == 0 {
    			matrix[i][0] = 0
    			matrix[0][j] = 0
    		}
    	}
    }
    //set 0
    //cols
    for j:=1;j<cols;j++ {
    	if matrix[0][j] == 0 {
    		for i:=1;i<rows;i++ {
    			matrix[i][j] = 0
    		}
    	}
    }
    //rows
    for i:=1;i<rows;i++ {
    	if matrix[i][0] == 0 {
    		for j:=1;j<cols;j++ {
    			matrix[i][j] = 0
    		}
    	}
    }
    //column
    if matrix[0][0] == 0{
        for i:=1;i<rows;i++ {
            matrix[i][0] = 0
        }
    }
    
    //firstline
    if firLine0 {
    	for i:=1;i<cols;i++ {
    		matrix[0][i] = 0
    	}
    }
}