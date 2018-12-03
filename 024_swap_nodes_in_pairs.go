func swapPairs(head *ListNode) *ListNode {
    pre_head := ListNode{0,head}
    cur := &pre_head

    for cur.Next!=nil && cur.Next.Next!=nil {
        temp := cur.Next.Next;
    	cur.Next.Next = temp.Next
    	temp.Next = cur.Next
    	cur.Next = temp
    	cur = cur.Next.Next
    }
    return pre_head.Next
}