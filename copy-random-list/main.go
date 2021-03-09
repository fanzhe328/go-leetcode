package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func main() {

}

// 拼接+拆分方法
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	p := head
	var next *Node
	for p.Next != nil {
		next = p.Next
		node := Node{p.Val, next, nil}
		p.Next = &node
		p = next
	}
	node := Node{p.Val, nil, nil}
	p.Next = &node

	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}

	p = head
	newHead := p.Next
	newP := newHead
	for newP.Next != nil {
		p.Next = newP.Next
		p = p.Next
		newP.Next = p.Next
		newP = newP.Next
	}
	p.Next = nil
	newP.Next = nil
	return newHead
}

// 哈希表方法
// func copyRandomList(head *Node) *Node {
// 	if head == nil {
// 		return nil
// 	}
// 	m := map[*Node][]*Node{}
// 	newHead := &Node{}
// 	p, newP := head, newHead
// 	newPre := newP
// 	for p != nil {
// 		if p.Random != nil {
// 			if m[p.Random] == nil {
// 				m[p.Random] = []*Node{}
// 			}
// 			m[p.Random] = append(m[p.Random], newP)
// 		}
// 		newP.Val = p.Val
// 		newP.Next = &Node{}
// 		p = p.Next
// 		newPre = newP
// 		newP = newP.Next
// 	}
// 	newPre.Next = nil
// 	p, newP = head, newHead
// 	for p != nil {
// 		if _, ok := m[p]; ok {
// 			for _, item := range m[p] {
// 				item.Random = newP
// 			}
// 		}
// 		p = p.Next
// 		newP = newP.Next
// 	}
// 	return newHead
// }
