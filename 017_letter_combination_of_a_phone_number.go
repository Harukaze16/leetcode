func letterCombinations(digits string) []string {
    if(len(digits)==0){
        return []string{}
    }
    res := []string{""}
    num_ch := []string{"","","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"}

    for _,num := range digits {
        index := int(num-'0')
        last_sz := len(res)
        each_sz := len(res[0])
        for res_index,cur_str := range res {
            cur_byte_arr := []byte(cur_str)
            cur_byte_arr = append(cur_byte_arr,num_ch[index][0])
            //变量名重复必须应避免，should be careful
            res[res_index] = string(cur_byte_arr)
        }
        for i :=1;i<len(num_ch[index]);i++{
            for j:=0;j<last_sz;j++ {
                cur_arr := []byte(res[j])
                cur_arr[each_sz] = num_ch[index][i]
                res = append(res,string(cur_arr))
            }
        }
    }
    return res
}