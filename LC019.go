/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    preHead := &ListNode{}
    preHead.Next = head
    tail := preHead
    delNode := preHead
    for i := 0; i < n && tail != nil; i++ {
        tail = tail.Next
    }
    if tail == nil {
        return nil
    }
    
    for tail.Next != nil {
        tail = tail.Next
        delNode = delNode.Next
    }
    delNode.Next = delNode.Next.Next
    return preHead.Next
}
