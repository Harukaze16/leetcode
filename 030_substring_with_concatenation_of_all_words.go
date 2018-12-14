func findSubstring(s string, words []string) []int {
	if s=="" || len(words)==0 {
		return []int{}
	}

	//compute the words
	count := make(map[string]int)
	for _,cur := range words {
		count[cur]++
	}

	var res []int
	ssz,word_num,word_len := len(s),len(words),len(words[0])
	for i:=0;i<word_len;i++ {
		j ,cur_num,begin := i,0,i
		seen := make(map[string]int)
        for ;j<=ssz-word_len*(word_num-cur_num);j+=word_len{  //
			cur_str := s[j:j+word_len]
			if count[cur_str]==0 { //didn't exist in words
				cur_num=0;
				seen = make(map[string]int)
				begin = j+word_len
			} else if seen[cur_str]==count[cur_str] { //duplicate
                seen[cur_str]++
                cur_num++   //verty important, 保证cur_num的有效性
				for t:= begin;t<j;t+=word_len {
                    cur := s[t:t+word_len]
					seen[cur]--
					begin += word_len
					cur_num--
					if cur_str==cur {
						break;
					}
				}
			} else {
				seen[cur_str]++
				cur_num++
				if cur_num == word_num {
					res = append(res,begin)
                    begin_str := s[begin:begin+word_len]
					seen[begin_str]--
					begin+=word_len
					cur_num--
				}
			}
		}
	}
	return res
}