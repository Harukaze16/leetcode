func mySqrt(x int) int {
    res := x
    for res*res > x {
        res = (x/res+res)/2
    }
	return res 
}