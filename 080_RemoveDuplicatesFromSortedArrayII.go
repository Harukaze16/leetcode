func removeDuplicates(nums []int) int {
    cur := 0
    for _,num := range nums {
        if cur<2 || num>nums[cur-2] {
            nums[cur] = num
            cur++
        }
    }
    return cur
}