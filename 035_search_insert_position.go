func searchInsert(nums []int, target int) int {
    beg,end := 0,len(nums)
    for beg<end {
        mid := (beg+end)/2
        if(nums[mid]<target) {
            beg=mid+1
        } else if(nums[mid]>target){
            end=mid
        } else {
            return mid
        }
    }
    return end
}