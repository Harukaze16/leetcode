func isNumber(s string) bool {
    var num,exp,point bool
    var i,j int
    //skip space
    for i=0;i<len(s) && s[i]==' ';i++{
    }
    //skip sign
    if i<len(s) && (s[i]=='+' || s[i]=='-') {
        i++
    }
    //skip space
    for j=len(s)-1;j>=i && s[j]==' ';j--{
    }

    for ;i<=j;i++ {
    	switch s[i] {
            case '0','1','2','3','4','5','6','7','8','9':
                num = true //较优，要应对.digit的情况
            case '.':
                if point || exp{
                    return false
                } else {
                    point = true
                }
            case 'e':
                if exp || !num || i==j{
                    return false
                } else {
                    if i+1<j && (s[i+1]=='+' || s[i+1] == '-') {
                        i++
                    }
                    exp = true
                }
            default:
                return false
    	}
    }
    return num
}