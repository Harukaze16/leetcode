func nextPermutation(nums []int)  {
	end := len(nums)-1
	cur := end
	for cur>0 && nums[cur]<=nums[cur-1]{
		cur--
	}
	if cur==0 {
		reverse(nums[cur:end+1])
	} else {
		pre_index := cur-1
		for cur<end && nums[cur+1]>nums[pre_index] {
			cur++
		}
		swap(&nums[cur],&nums[pre_index])
		reverse(nums[pre_index+1:end+1])
	}
}

func reverse(cur []int) {
    beg,end := 0,len(cur)-1
    for beg<end {
        swap(&cur[beg],&cur[end])
        beg++
        end--
    }
}
func swap(a *int, b *int) {
    temp := *a
    *a = *b
    *b = temp
}