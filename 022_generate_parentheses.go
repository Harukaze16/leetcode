func generateParenthesis(n int) []string {
    return generate_recur(n,0,0,string(""))
}

func generate_recur(n int,lpar int,rpar int,str string) []string{
    var res []string
    if  lpar == n && rpar == n {
        return []string{str}
    }
    //以下结构是经优化后的，更简单，考虑的更少，可以不必历经那么多else if
    if rpar<lpar { //need the r
        res =generate_recur(n,lpar,rpar+1,str+")")
    } 
    if(lpar<n) {
         res = append(res,generate_recur(n,lpar+1,rpar,str+"(")...)
    }
    return res
}