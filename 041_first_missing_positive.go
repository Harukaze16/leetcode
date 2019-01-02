func firstMissingPositive(nums []int) int {
    i,sz := 0,len(nums)
    for i =0;i<sz;i++ {
        for nums[i]>0 && nums[i]<sz && nums[i]!=i+1 && nums[nums[i]-1]!=nums[i] {
            swap(&nums[i],&nums[nums[i]-1])
        }
    }
    for i=0;i<sz && nums[i]==i+1;i++ {
    }
    return i+1;
}

func swap(a *int,b *int) {
    temp := *a
    *a = *b
    *b = temp
}