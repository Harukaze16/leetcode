func isValid(s string) bool {
    var mis_match []byte
    par_dict := map[byte]byte{
    	')':'(',
    	']':'[',
    	'}':'{',
    }
    for  _,temp := range s{
        ch := byte(temp)
    	if(par_dict[ch]==0) {
    		mis_match = append(mis_match,ch)
        } else if(len(mis_match)>0 && par_dict[ch]==mis_match[len(mis_match)-1]) {
            //notice the case input ) ,but stack is empty
            mis_match = mis_match[:len(mis_match)-1]
    	} else {
    		return false;
    	}
    }
    if(len(mis_match)==0) {
    	return true
    }
    return false
}