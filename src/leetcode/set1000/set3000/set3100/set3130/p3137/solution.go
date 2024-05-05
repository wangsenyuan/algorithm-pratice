package p3137

func minimumOperationsToMakeKPeriodic(word string, k int) int {
	n := len(word)
	// 这样子可能会造成超时
	// hash 也是可以的
	tr := NewNode()
	var best int
	for i := 0; i < n; i += k {
		tmp := tr
		for j := i; j < i+k; j++ {
			x := int(word[j] - 'a')
			if tmp.children[x] == nil {
				tmp.children[x] = NewNode()
			}
			tmp = tmp.children[x]
		}
		tmp.cnt++
		best = max(best, tmp.cnt)
	}

	return n/k - best
}

func NewNode() *Node {
	node := new(Node)
	node.children = make([]*Node, 26)
	return node
}

type Node struct {
	children []*Node
	cnt      int
}
