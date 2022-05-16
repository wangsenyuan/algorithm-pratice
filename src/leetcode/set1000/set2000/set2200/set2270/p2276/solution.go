package p2276

type CountIntervals struct {
	root *Node
}

const INF = 1000000000

func Constructor() CountIntervals {
	root := NewNode(1, INF)
	return CountIntervals{root}
}

func (this *CountIntervals) Add(left int, right int) {
	this.root.Add(left, right)
}

func (this *CountIntervals) Count() int {
	return this.root.cnt
}

type Node struct {
	L, R  int
	left  *Node
	right *Node
	cnt   int
}

func NewNode(L, R int) *Node {
	node := new(Node)
	node.L = L
	node.R = R
	node.cnt = 0
	return node
}

func (node *Node) push() {
	if node.L < node.R {
		if node.left == nil {
			mid := (node.L + node.R) / 2
			node.left = NewNode(node.L, mid)
			node.right = NewNode(mid+1, node.R)
		}
	}
}

func (node *Node) Add(left, right int) {
	if right < node.L || node.R < left {
		return
	}
	if left <= node.L && node.R <= right {
		node.cnt = node.R - node.L + 1
		return
	}

	if node.cnt == node.R-node.L+1 {
		// no need go down
		return
	}

	node.push()

	node.left.Add(left, right)
	node.right.Add(left, right)

	node.cnt = node.left.cnt + node.right.cnt
}

/**
 * Your CountIntervals object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(left,right);
 * param_2 := obj.Count();
 */
