/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head==nil || k==0 {
        return head
    }
    sz := 0
    for cur:=head;cur!=nil;cur=cur.Next {
        sz++
    }
    k = k%sz
    
    pre,follow := head,head
    for i:=0;i<k;i++ {
        pre = pre.Next
    }
    for pre.Next!=nil {
        pre=pre.Next
        follow = follow.Next
    } 
    
    pre.Next = head
    head = follow.Next
    follow.Next = nil
    return head
}