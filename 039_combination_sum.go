func combinationSum(candidates []int, target int) [][]int {
    var res [][]int
    sort.Ints(candidates)
    recur_combination(candidates,[]int{},target,&res)
    return res
}

func recur_combination(candidates []int, cur_candidates []int,target int,res* [][]int) {
    if target == 0 {
        now := make([]int,len(cur_candidates))
        copy(now,cur_candidates)
        *res = append(*res,now)
    }
    if len(candidates)==0 || target<0 {
        return
    }
    sz := len(candidates)
    for i:=0; i<sz && target>0;i++ {
       cur_candidates = append(cur_candidates,candidates[i])
       recur_combination(candidates[i:],cur_candidates,target-candidates[i],res)
       cur_candidates = cur_candidates[:len(cur_candidates)-1]
    }
}