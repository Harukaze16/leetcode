func insert(intervals []Interval, newInterval Interval) []Interval {
    i,sz := 0,len(intervals)
    var res []Interval
    for i<sz && intervals[i].End<newInterval.Start {
    	res = append(res,intervals[i])
    	i++
    }
    // 交融部分
    for i<sz && newInterval.End>=intervals[i].Start {
    	newInterval.Start = min(newInterval.Start,intervals[i].Start)
    	newInterval.End = max(newInterval.End,intervals[i].End)
    	i++
    }
    res = append(res,newInterval)
    for i<sz {
    	res = append(res,intervals[i])
    	i++
    }
    return res
}

func min(a,b int) int{
    if a<=b {
        return a
    }
    return b
}
func max(a,b int) int{
    if a>=b {
        return a
    }
    return b
}