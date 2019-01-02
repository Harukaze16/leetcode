/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head==nil {
        return head
    }
    for cur := head;cur!=nil;cur=cur.Next{
        for cur.Next!=nil && cur.Next.Val==cur.Val {
            cur.Next=cur.Next.Next
        }
    }
    return head
}