func fullJustify(words []string, maxWidth int) []string {

    var finalRes []string
   	for i,sz := 0,len(words);i<sz; {
        var res []byte
   		j,curlen := i,0
   		for j<sz && curlen+len(words[j])+j-i<=maxWidth { //noti
   			curlen,j = curlen+len(words[j]),j+1
   		}
        emptyCount := j-i-1
        if emptyCount == 0 {
            res = append(res,words[i]...)
            res = append(res,makeSpace(maxWidth - curlen)...)
            i++
        } else if j == sz{ //last line
            res = append(res,words[i]...)
            for i=i+1;i<j;i++ {
                res = append(res,' ')
                res = append(res,words[i]...)
            }
            res = append(res,makeSpace(maxWidth-len(res))...)
        } else {
            emptyArr := makeSpace( (maxWidth - curlen) / emptyCount)
            res = append(res,words[i]...)
            remainSapce := (maxWidth - curlen)%emptyCount

            for i=i+1;i<j;i++ {
                if remainSapce > 0 {
                  res = append(res,' ')
                  remainSapce--
                }
                res = append(res,emptyArr...)
                res = append(res,words[i]...)
            }
        }
        finalRes = append(finalRes,string(res))
   	}
   	return finalRes
}

func max(a,b int) int {
	if a>=b {
		return a
	}
	return b
}

func makeSpace(n int) []byte {
	res := make([]byte,n)
	for i:=0;i<n;i++ {
		res[i] = ' '
	}
	return res
}