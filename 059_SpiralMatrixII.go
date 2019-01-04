func generateMatrix(n int) [][]int {
    res := make([][]int,n)
    for i:=0;i<n;i++ {
    	res[i] = make([]int,n)
    }

    rb,cb,re,ce := 0,0,n-1,n-1
    
    num := 1
    for rb<re && cb < ce {
    	for i:=cb;i<=ce;i,num = i+1,num+1{
    		res[rb][i] = num
    	}
    	for i:= rb+1;i<=re;i,num = i+1,num+1 {
    		res[i][ce] = num
    	}
    	for i:= ce-1;i>=cb;i,num = i-1,num+1 {
    		res[re][i] = num
    	}
    	for i:=re-1;i>rb;i,num = i-1,num+1 {
    		res[i][cb] = num
    	}
        rb,cb,re,ce = rb+1,cb+1,re-1,ce-1
    }
    for ;rb==re && cb<=ce;cb,num = cb+1,num+1 {
    	res[rb][cb] = num
    }
    for ;cb==ce && rb<=re;rb,num = rb+1,num+1 {
    	res[rb][cb] = num
    }
    return res
}