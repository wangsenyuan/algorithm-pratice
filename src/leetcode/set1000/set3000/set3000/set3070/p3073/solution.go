package p3073

func resultArray(nums []int) []int {
	var t1 = Insert(nil, nums[0])
	a1 := []int{nums[0]}
	var t2 = Insert(nil, nums[1])
	a2 := []int{nums[1]}

	for i := 2; i < len(nums); i++ {
		c1 := CountGreater(t1, nums[i])
		c2 := CountGreater(t2, nums[i])
		if c1 > c2 || c1 == c2 && len(a1) <= len(a2) {
			t1 = Insert(t1, nums[i])
			a1 = append(a1, nums[i])
		} else {
			t2 = Insert(t2, nums[i])
			a2 = append(a2, nums[i])
		}
	}

	return append(a1, a2...)
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int
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

func NewNode(key int) *Node {
	node := new(Node)
	node.key = key
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

func Insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}
	if node.key == key {
		node.cnt++
		node.size++
		return node
	}

	if node.key > key {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
	}

	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.size = node.left.Size() + node.right.Size() + node.cnt

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

func MinValueNode(root *Node) *Node {
	cur := root

	for cur.left != nil {
		cur = cur.left
	}

	return cur
}

func (root *Node) MinValue() int {
	if root == nil {
		return 0
	}
	node := MinValueNode(root)

	return node.key
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
		root.size--
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
		return nil
	}

	root.height = max(root.left.Height(), root.right.Height()) + 1
	root.size = root.left.Size() + root.right.Size() + root.cnt
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

func CountGreater(node *Node, key int) int {
	if node == nil {
		return 0
	}
	if node.key <= key {
		return CountGreater(node.right, key)
	}
	return node.right.Size() + node.cnt + CountGreater(node.left, key)
}
