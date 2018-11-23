func longestPalindrome(s string) string {   
    array := coping_str(s) //对s进行预处理
    sz := len(array)

    len_arr := make([]int, sz)//创建了长度为5初值为0的切片
    max_center,max_right := 0,0  //最大长度的len,扩展范围最大的最右端index,中心处index

    max_len_i ,max_len := 0,0

    for i:= 0; i<sz;i++{
        var cur_len int
        if max_right > i {
            symmer_i := max_center*2 - i
            cur_len = min(max_right-i,len_arr[symmer_i])                                     
        }
        for i-cur_len>= 0 && i+cur_len <sz && array[i-cur_len]==array[i+cur_len] {
            cur_len++;
        }
        len_arr[i] = cur_len;
        //最右扩展变更
        if(i+cur_len>max_right) {
            max_right = i+cur_len;
            max_center = i
        }
        if(cur_len>max_len) {
            max_len = cur_len;
            max_len_i = i
        }
    }

    begi,endi := max_len_i-max_len+1,max_len+max_len_i-1
    return s[begi/2:endi/2]   
}
func coping_str(s string) []byte {
    array := []byte{'#'}
    copy := []byte(s)
    for _,val := range copy {
        array = append(array,val)
        array = append(array,'#')
    }
    return array
}
func min(a int, b int) int {
    if a<=b {
        return a
    }
    return b
}