func getPermutation(n int, k int) string {
	//compute permutate num
	perm := make([]int,n+1)
	perm[0] = 1;
    for i:=1;i<=n;i++ {
		perm[i]=i*perm[i-1]
	}

    cur := make([]int,n)
    for i,_ := range cur {
    	cur[i]=i+1
    }
    
    var res []int    
    for k--;k>0; {
        t := k/perm[n-1]
    	res = append(res,cur[t])
        cur = append(cur[:t],cur[t+1:]...)
    	k = k%perm[n-1]
    	n--
    }
    res = append(res,cur...)
    
    var final_res string
    for _,num := range res {
        final_res += strconv.Itoa(num)
    }
    return final_res
}