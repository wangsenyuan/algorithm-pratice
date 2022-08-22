package p2382

const INF = 1 << 60

func maximumSegmentSum(nums []int, removeQueries []int) []int64 {
	n := len(nums)

	root := NewNode(0, n-1)

	ans := make([]int64, n)

	for i := n - 1; i >= 0; i-- {
		ans[i] = root.val

		id := removeQueries[i]

		Set(root, id, int64(nums[id]))
	}

	return ans
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

type Node struct {
	start int
	end   int
	left  *Node
	right *Node
	val   int64
	l_sum int64
	r_sum int64
	cnt   int
}

func NewNode(start int, end int) *Node {
	node := new(Node)
	node.start = start
	node.end = end
	node.val = 0
	node.l_sum = 0
	node.r_sum = 0
	node.cnt = 0

	return node
}

func (node *Node) push() {
	if node.left == nil {
		mid := (node.start + node.end) / 2
		node.left = NewNode(node.start, mid)
		node.right = NewNode(mid+1, node.end)
	}
}

func (node *Node) GetValue() int64 {
	if node == nil {
		return 0
	}
	return node.val
}

func (node *Node) GetCount() int {
	if node == nil {
		return 0
	}
	return node.cnt
}

func (node *Node) IsFull() bool {
	return node != nil && node.cnt == (node.end-node.start+1)
}

func (node *Node) GetLeftSum() int64 {
	if node == nil {
		return 0
	}
	return node.l_sum
}

func (node *Node) GetRightSum() int64 {
	if node == nil {
		return 0
	}
	return node.r_sum
}

func Set(node *Node, p int, value int64) {
	if node.start == node.end {
		node.val = value
		node.l_sum = value
		node.r_sum = value
		node.cnt = 1
		return
	}
	node.push()

	mid := (node.start + node.end) / 2
	if p <= mid {
		Set(node.left, p, value)
	} else {
		Set(node.right, p, value)
	}

	node.val = max(node.left.GetValue(), node.right.GetValue())
	node.cnt = node.left.GetCount() + node.right.GetCount()

	if !node.left.IsFull() {
		node.l_sum = node.left.GetLeftSum()
	} else {
		node.l_sum = node.left.GetLeftSum() + node.right.GetLeftSum()
	}

	if !node.right.IsFull() {
		node.r_sum = node.right.GetRightSum()
	} else {
		node.r_sum = node.right.GetRightSum() + node.left.GetRightSum()
	}
	node.val = max(node.val, max(max(node.l_sum, node.r_sum), node.left.GetRightSum()+node.right.GetLeftSum()))
}
