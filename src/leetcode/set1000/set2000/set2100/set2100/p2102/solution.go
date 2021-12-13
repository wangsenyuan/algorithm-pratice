package p2102

type Item struct {
	key int
	val string
}

func (this Item) Less(that Item) bool {
	return this.key > that.key || this.key == that.key && this.val < that.val
}

/**
* this is a AVL tree
 */
type Node struct {
	item        Item
	height      int
	cnt         int
	size        int
	left, right *Node
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (node *Node) Size() int {
	if node == nil {
		return 0
	}
	return node.size
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func NewNode(item Item) *Node {
	node := new(Node)
	node.item = item
	node.height = 1
	node.cnt = 1
	node.size = 1
	return node
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	y.height = max(y.left.Height(), y.right.Height()) + 1
	y.size = y.left.Size() + y.right.Size() + y.cnt

	x.height = max(x.left.Height(), x.right.Height()) + 1
	x.size = x.left.Size() + x.right.Size() + x.cnt

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.height = max(x.left.Height(), x.right.Height()) + 1

	x.size = x.left.Size() + x.right.Size() + x.cnt

	y.height = max(y.left.Height(), y.right.Height()) + 1

	y.size = y.left.Size() + y.right.Size() + y.cnt

	return y
}

func (node *Node) GetBalance() int {
	if node == nil {
		return 0
	}
	return node.left.Height() - node.right.Height()
}

func FindEqualOrGreater(node *Node, item Item) *Node {
	if node == nil {
		return nil
	}
	if node.item.Less(item) {
		return FindEqualOrGreater(node.right, item)
	}

	res := FindEqualOrGreater(node.left, item)
	if res == nil {
		return node
	}
	return res
}

func FindEqualOrLess(node *Node, item Item) *Node {
	if node == nil {
		return nil
	}
	if item.Less(node.item) {
		return FindEqualOrLess(node.left, item)
	}

	res := FindEqualOrLess(node.right, item)
	if res == nil {
		return node
	}
	return res
}

func Insert(node *Node, item Item) *Node {
	if node == nil {
		return NewNode(item)
	}
	if node.item == item {
		node.cnt++
		node.size++
		return node
	}

	if item.Less(node.item) {
		node.left = Insert(node.left, item)
	} else {
		node.right = Insert(node.right, item)
	}

	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.size = node.left.Size() + node.right.Size() + node.cnt
	balance := node.GetBalance()

	if balance > 1 && item.Less(node.left.item) {
		return rightRotate(node)
	}

	if balance < -1 && node.right.item.Less(item) {
		return leftRotate(node)
	}

	if balance > 1 && node.left.item.Less(item) {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && item.Less(node.right.item) {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func MinValueNode(root *Node) *Node {
	cur := root

	for cur.left != nil {
		cur = cur.left
	}

	return cur
}

func FindKth(root *Node, k int) *Node {
	if root.left.Size() >= k {
		return FindKth(root.left, k)
	}
	// root.left.Size() < k
	if root.left.Size()+root.cnt >= k {
		return root
	}

	return FindKth(root.right, k-root.left.Size()-root.cnt)
}

type SORTracker struct {
	root *Node
	cnt  int
}

func Constructor() SORTracker {
	return SORTracker{nil, 0}
}

func (this *SORTracker) Add(name string, score int) {
	item := new(Item)
	item.key = score
	item.val = name
	this.root = Insert(this.root, *item)
}

func (this *SORTracker) Get() string {
	this.cnt++
	node := FindKth(this.root, this.cnt)
	return node.item.val
}

/**
 * Your SORTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(name,score);
 * param_2 := obj.Get();
 */
