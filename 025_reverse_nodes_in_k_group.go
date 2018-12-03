
func reverseKGroup(head *ListNode, k int) *ListNode {
    pre_heading := ListNode{0,head}
	pre,cur := &pre_heading,&pre_heading

	for pre!=nil{
        pre = cur //notice this,it's easy to making mistake
		for i:=0;i<k && pre!=nil;i++ {
            pre = pre.Next
		}
		if(pre!=nil) {
			next,loop_end := cur.Next.Next,cur.Next
			for i:=0;i<k-1;i++{
				temp := next.Next
				next.Next = cur.Next
				cur.Next = next
				loop_end.Next = temp // important, easy to ingore the element 
				next = temp
			}
            cur = loop_end
		}
    }
	return pre_heading.Next
}