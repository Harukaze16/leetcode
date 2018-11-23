/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head ListNode
	cur := &head
	accr := 0
    for l1!=nil && l2!=nil{
    	cur.Next = &ListNode{(l1.Val+l2.Val+accr)%10,nil}
    	accr = (l1.Val+l2.Val+accr)/10
    	cur = cur.Next
        l1=l1.Next
        l2=l2.Next
    }
    var Remain *ListNode
    if(l1!=nil) {
        Remain = l1
    } else {
        Remain = l2
    }
    for Remain !=nil{
    	cur.Next = &ListNode{(Remain.Val+accr)%10,nil}
    	accr = (Remain.Val + accr)/10
    	cur = cur.Next
        Remain = Remain.Next
    }
    if accr!=0 {
        cur.Next = &ListNode{accr,nil}
    }
    return head.Next
}