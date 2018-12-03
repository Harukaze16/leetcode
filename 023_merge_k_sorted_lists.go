func mergeKLists(lists []*ListNode) *ListNode {
    sz := len(lists)
    if sz == 0 {
        return nil
    }
    for sz > 1 {
        cur_loop_sz := sz/2
        for i:=0;i<cur_loop_sz;i++ {
            lists[i] = merge2Lists(lists[2*i],lists[2*i+1])
        }
        if(sz%2 != 0) { //odd,this is sz ,not cur_loop_sz,notice!
            lists[cur_loop_sz] = lists[sz-1]
        }
        sz = (sz+1)/2
    }
    return lists[0]
}

func merge2Lists(l1 *ListNode,l2 *ListNode) *ListNode {
    var head ListNode
    cur := &head
    for l1!=nil && l2!=nil {
        if(l1.Val<=l2.Val) {
            cur.Next = l1;
            l1 = l1.Next
        } else {
            cur.Next = l2;
            l2 = l2.Next
        }
        cur = cur.Next
    }
    if(l1==nil) {
        cur.Next = l2;
    } else {
        cur.Next = l1;
    }
    return head.Next
}