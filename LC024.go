/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    preHead := &ListNode{}
    preHead.Next = head
    l := preHead
    for l.Next != nil && l.Next.Next != nil {
        nextNode := l.Next
        nextNextNode := l.Next.Next
        nextNode.Next = nextNextNode.Next
        nextNextNode.Next = nextNode
        l.Next = nextNextNode
        l = nextNode
    }
    return preHead.Next
}
