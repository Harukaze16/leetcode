/*
my recursive solution
*/
func myPow(x float64, n int) float64 { 
    if n<0 {
    	x = 1/x
    	n = -n
    }
    return pow(x,n)
}

//assert n >=0
func pow(x float64,n int) float64 {
	if n==0 {
		return 1
	} else if n%2 ==0 { //even
		val := pow(x,n/2)
		return val*val
	} else { //odd
		val := pow(x,n/2)
		return val*val*x
	}
}


//iterative o(1) space
func myPow(x float64, n int) float64 {    
    if n<0 {
    	x = 1/x
    	n = -n
    }

    res := float64(1)
    for n>0 {
    	if n & 1 >0{
    		res *= x
    	}
    	x = x*x
    	n>>=1
    }
    return res
}