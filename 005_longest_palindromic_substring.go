func longestPalindrome(s string) string {
    array := []byte(s)
    beg,end := 0,len(s)
    cur := beg
    self_divide := true
    max_begin,max_len := 0,0
    for cur < end {
    	var cur_beg, cur_end int
    	if self_divide {
    		cur_beg,cur_end = cur, cur
    	} else {
    		cur_beg,cur_end = cur,cur+1
    		cur++
    	}
    	self_divide = !self_divide
    	for cur_beg>=beg && cur_end<end && array[cur_beg]==array[cur_end]{
    		cur_beg--;
    		cur_end++;
    	}
    	if(cur_end-cur_beg-1>max_len) {
    		max_len = cur_end-cur_beg-1;
    		max_begin = cur_beg+1
    	}
    }
    fmt.Println(max_begin,max_len)
    return string(array[max_begin:(max_len+max_begin)])
}