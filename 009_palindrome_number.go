func isPalindrome(x int) bool {
    if x<10 {
        return x>=0
    }
    if x%10 == 0 {
        //notice 100,200这些需要额外处理
        return false
    }
    
    res := 0
    for x>res {
        res = res*10+x%10
        if (x==res) {
            return true
        }
        x /= 10
    }
    return x==res
    
}