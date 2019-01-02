func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    next := compute_next(candidates)
    var res [][]int
    recur_combination(candidates,[]int{},target,next,&res,0)
    return res
}

func recur_combination(candidates []int,cur_candi []int,target int,next []int,res* [][]int,index int){
    if target==0 {
        temp := make([]int,len(cur_candi))
        copy(temp,cur_candi)
        *res = append(*res,temp)
    }
    sz := len(candidates)
    for i:=index;i<sz;i=next[i] {
        cur_candi=append(cur_candi,candidates[i])       
        if target-candidates[i]<0 {
            return
        }
        recur_combination(candidates,cur_candi,target - candidates[i],next,res,i+1)
        //next存储的是原始current下标，而非偏移
        cur_candi = cur_candi[:len(cur_candi)-1]
    }
}


func compute_next(candidates []int) []int {
    next := make([]int,len(candidates))
    for i:=len(candidates)-1;i>=0;i--{
        next[i]=i+1
        for cur:=i+1;i>0 && candidates[i-1]==candidates[i];i--{
            next[i-1]=cur
        }
    }
    return next
}