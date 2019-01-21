func sortColors(nums []int)  {
    lessRight := -1
    for i,j:=0,len(nums)-1;i<=j; {
    	if nums[i]<1 {
            swap(&nums[i],&nums[lessRight+1])
    		lessRight++
    		i++
    	} else if nums[i]>1 {
    		swap(&nums[i],&nums[j])
    		j--
    	} else {
    		i++
    	}
   	}
}


func swap(a,b *int) {
	temp := *a
	*a = *b
	*b = temp
}