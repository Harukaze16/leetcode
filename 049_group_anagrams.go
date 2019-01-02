func groupAnagrams(strs []string) [][]string {
	dict := make(map[string]int)
    var res [][]string
    
    index := 0
    for _,s := range strs {
    	sorted_s := sortedStr(s)
        cur,ok := dict[sorted_s]
    	if !ok {
    		dict[sorted_s]=index;
    		res = append(res,[]string{})
    		res[index]=append(res[index],s)
    		index++
    	} else {
    		res[cur] = append(res[cur],s)
    	}
    }
    return res
}

func sortedStr(str string) string{
    bytearr := []byte(str)
    sort.Slice(bytearr,func(a,b int) bool {
        return bytearr[a]<bytearr[b]
    })
	return string(bytearr)
}