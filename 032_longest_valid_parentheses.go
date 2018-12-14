func longestValidParentheses(s string) int {
    var mismatch_par []int
    sz := len(s)
    for i:=0;i<sz;i++ {
        mis_end := len(mismatch_par)-1
        if mis_end>=0 && s[i]==')' && s[mismatch_par[mis_end]]=='(' {
            mismatch_par = mismatch_par[0:mis_end]
        } else {
            mismatch_par = append(mismatch_par,i)
        }
    }
    
    //notice
    mismatch_par = append(mismatch_par,sz)
    //find max    
    //the special cope for head and tail should notice 
    max := mismatch_par[0]+1
    
    for i:=len(mismatch_par)-1;i>0;i-- {
        if mismatch_par[i]-mismatch_par[i-1]>max {
            max = mismatch_par[i]-mismatch_par[i-1]
        }
    }
    return max-1
}