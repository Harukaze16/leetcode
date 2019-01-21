func combine(n int, k int) [][]int {
    if k>n {
    	return [][]int{}
    }

    var res [][]int
    
    cur := make([]int,k)
    for i:=0;i<k;i++ {
    	cur[i]=i+1
    }

    fmt.Print(cur)
    var stopIndex = k
    for stopIndex>=0{
    	temp := make([]int,len(cur))
    	copy(temp,cur)
    	res = append(res,temp)

    	//find next cur
    	for stopIndex=k-1;stopIndex>=0;stopIndex-- {
    		if cur[stopIndex] < n-k+stopIndex+1 {
    			cur[stopIndex]++
                for i:=stopIndex+1;i<k;i++ {
                    cur[i]=cur[stopIndex]+i-stopIndex
                }
                break
    		}
    	}
    }
    
    return res
}