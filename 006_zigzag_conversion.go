func convert(s string,numRows int) string {
	if numRows<=1 {
		return s;
	}  
	//row 0
    var res []byte
    sz := len(s)
	for j :=0;j<sz;j+=2*numRows-2{
        res = append(res,byte(s[j]))
	}
	//i==1 to numRows -2
	for i :=1;i<numRows-1;i++{
        for j:=i;j<sz;{
            res = append(res,s[j])
            j+=2*(numRows-i-1)
            if j<sz {
                res = append(res,s[j])
            }
            j+=2*i //对应关系开始没有认知清晰
        }
	} 
    for j :=numRows-1;j<sz;j+=2*numRows-2{
        res = append(res,byte(s[j]))
	}
    return string(res)
}