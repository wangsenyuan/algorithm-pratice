package p715

/**
* this is a AVL tree
 */
type Node struct {
	key         int
	height      int
	cnt         int
	left, right *Node
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func NewNode(key int) *Node {
	node := new(Node)
	node.key = key
	node.height = 1
	node.cnt = 1
	return node
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	y.height = max(y.left.Height(), y.right.Height()) + 1
	x.height = max(x.left.Height(), x.right.Height()) + 1

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.height = max(x.left.Height(), x.right.Height()) + 1
	y.height = max(y.left.Height(), y.right.Height()) + 1

	return y
}

func (node *Node) GetBalance() int {
	if node == nil {
		return 0
	}
	return node.left.Height() - node.right.Height()
}

func Insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}
	if node.key == key {
		node.cnt++
		return node
	}

	if key < node.key {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
	}

	node.height = max(node.left.Height(), node.right.Height()) + 1
	balance := node.GetBalance()

	if balance > 1 && key < node.left.key {
		return rightRotate(node)
	}

	if balance < -1 && key > node.right.key {
		return leftRotate(node)
	}

	if balance > 1 && key > node.left.key {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && key < node.right.key {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func Ceiling(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if root.key >= key {
		res := Ceiling(root.left, key)
		if res != nil {
			return res
		}
		return root
	}
	return Ceiling(root.right, key)
}

func Floor(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if root.key <= key {

		res := Floor(root.right, key)

		if res != nil {
			return res
		}

		return root
	}
	return Floor(root.left, key)
}

func MinValueNode(root *Node) *Node {
	cur := root

	for cur.left != nil {
		cur = cur.left
	}

	return cur
}

func Delete(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if key < root.key {
		root.left = Delete(root.left, key)
	} else if key > root.key {
		root.right = Delete(root.right, key)
	} else {
		root.cnt--
		if root.cnt > 0 {
			return root
		}
		if root.left == nil || root.right == nil {
			tmp := root.left
			if root.left == nil {
				tmp = root.right
			}
			root = tmp
		} else {
			tmp := MinValueNode(root.right)

			root.key = tmp.key
			root.cnt = tmp.cnt
			// make sure tmp node deleted after call delete on root.right
			tmp.cnt = 1
			root.right = Delete(root.right, tmp.key)
		}
	}

	if root == nil {
		return root
	}

	root.height = max(root.left.Height(), root.right.Height()) + 1
	balance := root.GetBalance()

	if balance > 1 && root.left.GetBalance() >= 0 {
		return rightRotate(root)
	}

	if balance > 1 && root.left.GetBalance() < 0 {
		root.left = leftRotate(root.left)
		return rightRotate(root)
	}

	if balance < -1 && root.right.GetBalance() <= 0 {
		return leftRotate(root)
	}

	if balance < -1 && root.right.GetBalance() > 0 {
		root.right = rightRotate(root.right)
		return leftRotate(root)
	}

	return root
}

type RangeModule struct {
	tr     *Node
	values map[int]int
}

func Constructor() RangeModule {
	return RangeModule{nil, make(map[int]int)}
}

func (this *RangeModule) AddRange(left int, right int) {
	it := Floor(this.tr, left)

	if it != nil && this.values[it.key] >= left {
		left = it.key
	}

	it = Ceiling(this.tr, left)
	for it != nil && it.key <= right {
		key := it.key
		right = max(right, this.values[key])

		this.tr = Delete(this.tr, key)
		delete(this.values, key)

		it = Ceiling(this.tr, key)
	}

	// it == nil || it.key > right
	this.tr = Insert(this.tr, left)
	this.values[left] = right
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	it := Floor(this.tr, left)
	return it != nil && this.values[it.key] >= right
}

func (this *RangeModule) RemoveRange(left int, right int) {
	it := Floor(this.tr, left)
	if it != nil && this.values[it.key] >= left {
		if this.values[it.key] > right {
			this.tr = Insert(this.tr, right)
			this.values[right] = this.values[it.key]
		}
		if it.key < left {
			this.values[it.key] = left
		}
		// 如果it.key == left, 那么后续会被处理掉
	}

	it = Ceiling(this.tr, left)

	for it != nil && it.key < right {
		key := it.key
		// left <= it.key < right
		if right < this.values[key] {
			this.tr = Insert(this.tr, right)
			this.values[right] = this.values[key]
		}
		this.tr = Delete(this.tr, key)
		delete(this.values, key)
		it = Ceiling(this.tr, key)
	}
}
