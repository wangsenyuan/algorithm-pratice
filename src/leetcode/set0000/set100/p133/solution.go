package p133

/**
 * Definition for a Node.
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {

	var clone func(node *Node) *Node

	mem := make(map[*Node]*Node)

	clone = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		if v, found := mem[node]; found {
			return v
		}
		res := new(Node)

		mem[node] = res

		res.Val = node.Val
		res.Neighbors = make([]*Node, len(node.Neighbors))
		for i := 0; i < len(node.Neighbors); i++ {
			res.Neighbors[i] = clone(node.Neighbors[i])
		}

		return res
	}

	return clone(node)
}
