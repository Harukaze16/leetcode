func spiralOrder(matrix [][]int) []int {
    if len(matrix)==0 || len(matrix[0])==0  {
    	return []int{}
    }
    rowSz,colSz := len(matrix),len(matrix[0])
    var res []int
    rb,re,cb,ce := 0,rowSz-1,0,colSz-1
    for rb<re && cb <ce {
    	for c:=cb;c<=ce;c++ {
    		res = append(res,matrix[rb][c])
    	}
    	for r:=rb+1;r<=re;r++ {
    		res = append(res,matrix[r][ce])
    	}
    	for c:=ce-1;c>=cb;c--{
    		res = append(res,matrix[re][c])
    	}
    	for r:=re-1;r>rb;r--{
    		res = append(res,matrix[r][cb])
    	}
    	rb,re = rb+1,re-1
    	cb,ce = cb+1,ce-1
    }

    if rb<=re && cb==ce {
    	for i:=rb;i<=re;i++ {
    		res = append(res,matrix[i][cb])
    	}
    } else if cb<ce && rb==re {
    	for j:=cb;j<=ce;j++ {
    		res = append(res,matrix[rb][j])
    	}
    }
    return res
}