/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    
    res := &ListNode{}
    l := res
    flag := 0
    for l1 != nil || l2 != nil {
        sum := 0
        if l1 == nil {
            sum = l2.Val + flag
            l2 = l2.Next
        } else if l2 == nil {
            sum = l1.Val + flag
            l1 = l1.Next
        } else {
            sum = l1.Val + l2.Val + flag
            l1 = l1.Next
            l2 = l2.Next
        }
        flag = sum / 10
        l.Next = &ListNode{Val: sum % 10}
        l = l.Next
    }
    if flag > 0 {
        l.Next = &ListNode{Val: flag}
    }
    return res.Next
}
