/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func merge(intervals []Interval) []Interval {
    sort.Slice(intervals,func (i,j int) bool {
   		return intervals[i].Start<intervals[j].Start
    })
   
    if len(intervals) == 0 {
    	return []Interval{}
    }
    sz := len(intervals)
    cur := intervals[0]
    var res []Interval
    for i:=1;i<sz;i++ {
    	if intervals[i].Start<=cur.End && intervals[i].End>cur.End {
    		cur.End = intervals[i].End
    	} else if intervals[i].Start>cur.End {
    		res = append(res,cur)
    		cur = intervals[i]
    	}
    }
    res = append(res,cur)
    return res
}

