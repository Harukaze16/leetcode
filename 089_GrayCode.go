func grayCode(n int) []int {
    res := []int{0}
    cur_num := 1
    for i:=1;i<=n;i++ {
    	for cur:=len(res)-1;cur>=0;cur--{
    		res = append(res,res[cur]+cur_num)
    	}
    	cur_num <<= 1
    }
    return res
}	