/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
    heading := ListNode{0,head}
    pre,follow := &heading,&heading
    for i:=0;i<n && pre!=nil;i++ {
    	pre = pre.Next
    }
    // n小于链表长度
    if(pre==nil) {
    	return nil
    }
    for pre.Next!=nil {
    	pre=pre.Next
    	follow=follow.Next
    }
    follow.Next = follow.Next.Next

    return heading.Next
}