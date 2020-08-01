/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    oldNewNode := make(map[*Node]*Node)
    oldNewNode[nil] = nil
    p := head
    var newP, newHead *Node
    for p != nil {
        if newP == nil {
            newP = &Node{}
            newHead = newP
            oldNewNode[p] = newP
        }
        
        newP.Val = p.Val
        if _, ok := oldNewNode[p.Next]; ok {
            newP.Next = oldNewNode[p.Next]
        } else {
            newP.Next = &Node{}
            oldNewNode[p.Next] = newP.Next
        }
        
        if _, ok := oldNewNode[p.Random]; ok {
            newP.Random = oldNewNode[p.Random]
        } else {
            newP.Random = &Node{}
            oldNewNode[p.Random] = newP.Random
        }
        
        p = p.Next
        newP = newP.Next
    }
    return newHead
}
