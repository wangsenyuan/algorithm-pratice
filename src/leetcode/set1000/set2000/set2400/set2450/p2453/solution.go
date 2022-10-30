package p2453

func secondGreaterElement(nums []int) []int {
	n := len(nums)

	nodes := make([]*Node, n+1)
	nodes[n] = NewNode(-1, 0, n-1)

	for i := n - 1; i >= 0; i-- {
		nodes[i] = Set(nodes[i+1], i, nums[i])
	}

	stack := make([]int, n)
	var p int
	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = -1
		for p > 0 && nums[stack[p-1]] < nums[i] {
			x := Find(nodes[i+1], nums[stack[p-1]]+1)
			res[stack[p-1]] = x
			p--
		}
		stack[p] = i
		p++
	}

	return res
}

type Node struct {
	begin int
	end   int
	left  *Node
	right *Node
	val   int
}

func NewNode(v int, begin int, end int) *Node {
	node := new(Node)
	node.begin = begin
	node.end = end
	node.val = v
	return node
}

func (node *Node) Merge() {
	if node.left != nil {
		node.val = node.left.val
	}
	if node.right != nil {
		node.val = max(node.val, node.right.val)
	}
}

func (node *Node) GetLeft() *Node {
	if node == nil {
		return nil
	}
	return node.left
}

func (node *Node) GetRight() *Node {
	if node == nil {
		return nil
	}
	return node.right
}

// Set value at position p
func Set(root *Node, p int, v int) *Node {

	var loop func(src *Node, begin int, end int) *Node

	loop = func(src *Node, begin int, end int) *Node {
		res := NewNode(v, begin, end)
		if begin == end {
			return res
		}
		mid := (begin + end) / 2
		if p <= mid {
			res.left = loop(src.GetLeft(), begin, mid)
			res.right = src.GetRight()
		} else {
			res.left = src.GetLeft()
			res.right = loop(src.GetRight(), mid+1, end)
		}
		res.Merge()

		return res
	}

	return loop(root, root.begin, root.end)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Find(node *Node, v int) int {
	if node == nil || node.val < v {
		return -1
	}
	if node.begin == node.end {
		return node.val
	}
	if node.left != nil && node.left.val >= v {
		return Find(node.left, v)
	}
	return Find(node.right, v)
}
