func twoSum(nums []int, target int) []int {
    seek_index := make(map[int]int)
    for index,num := range nums {
        seeked_index,res := seek_index[num]
        if res {
            return []int{seeked_index,index}
        } 
        seek_index[target-num]=index
    }
    return []int{}
}