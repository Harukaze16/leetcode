func intToRoman(num int) string {
    if(num<=0 || num>3999) {
    	return ""
    }
    var res_arr []byte
    roman_str := "IVXLCDM"
    for cur_index,base := 6,1000;base>0;cur_index,base=cur_index-2,base/10{
    	if(num/base>0) {
    		cur := num/base
            num = num%base  //it's important ,so the num won't be influent the next loop of assigning of cur
    		if cur == 9 || cur == 4 {
    			res_arr = append(res_arr,roman_str[cur_index])
    			res_arr = append(res_arr,roman_str[cur_index+1+cur/5])
    		} else {
    			if cur>=5 {
    				res_arr = append(res_arr,roman_str[cur_index+1])
    				cur-=5
    			}
    			for ;cur>0;cur--{
    				res_arr = append(res_arr,roman_str[cur_index])
    			}
    		}
    	}
    }
    return string(res_arr)
}