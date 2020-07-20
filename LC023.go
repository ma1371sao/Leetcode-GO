/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type priorityQueue []*ListNode

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool { return pq[i].Val < pq[j].Val }

func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(*ListNode))
}

func (pq *priorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    node := old[n - 1]
    *pq = old[:n-1]
    return node
}

func mergeKLists(lists []*ListNode) *ListNode {
    preHead := &ListNode{}
    l := preHead
    pq := make(priorityQueue, 0)
    for i := range lists {
        if lists[i] != nil {
            heap.Push(&pq, lists[i])
        }
    }
    
    for pq.Len() > 0 {
        node := heap.Pop(&pq).(*ListNode)
        l.Next = node
        l = l.Next
        if node.Next != nil {
            heap.Push(&pq, node.Next)
        }
    }
    return preHead.Next
}
