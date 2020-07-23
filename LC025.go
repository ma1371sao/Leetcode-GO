/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    lLen := 0
    l := head
    for l != nil {
        lLen++
        l = l.Next
    }
    if k < 2 || lLen < k {
        return head
    }
    
    preHead := &ListNode{}
    preHead.Next = head
    pre := preHead
    l = head
    n := lLen / k
    var preK, nxt *ListNode
    for i := 0; i < n; i++ {
        preK = pre
        pre = l
        l = l.Next
        for j := 0; j < k - 1; j++ {
            nxt = l.Next
            l.Next = pre
            pre = l
            l = nxt
        }
        preK.Next.Next = l
        tmp := preK.Next
        preK.Next = pre
        pre = tmp
    }
    return preHead.Next
}
