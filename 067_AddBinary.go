func addBinary(a string, b string) string {
	res := make([]byte,max(len(a),len(b))+1)

	var carry byte = 0
	i,j := len(a)-1,len(b)-1
	for k:=len(res)-1;k>=0;k-- {
		var x,y byte
		if i>=0 {
			x = a[i]-'0'
			i--
		} else {
			x = 0
		}

		if j>=0 {
			y = b[j]-'0'
			j--
		} else {
			y = 0
		}
        res[k] = (x + y + carry)%2 + '0'
		carry = (x + y + carry)/2
	}
	if res[0] == '0' {
		return string(res[1:])
	} else {
		return string(res)
	}
}