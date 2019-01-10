func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    res := make([]byte,len(num1)+len(num2))
    for i:=len(res)-1;i>=0;i-- {
        res[i] = '0'
    }
    end1,end2 := len(num1)-1,len(num2)-1
    end := end1+end2+1
    for i:=0;i<=end1;i++ {
      var carry byte = 0
      for j:=0;j<=end2;j++ {
        val := (num1[end1-i]-'0')*(num2[end2-j]-'0') + (res[end-i-j]-'0') + carry
        res[end-i-j] = val % 10 + '0'
        carry = val /10 
      }
      res[end1-i] = carry + '0' //noti
    }
    if res[0] == '0' {
        return string(res[1:])
    } else {
        return string(res)
    }
}