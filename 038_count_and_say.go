func countAndSay(n int) string {
    res := "1"
    for i:=1;i<n;i++{
        sz := len(res)
        temp:= res
        res = ""
        for index:=0;index<sz; index++{
            src_index := index
            for index < sz-1 && temp[index]==temp[index+1] {
                index++
            }
            res += fmt.Sprintf("%d%c",index-src_index+1,temp[src_index])
        }
    }
    return res
}