package p3158

const S = 5 * 1e4

func getResults(queries [][]int) []bool {
	n := len(queries)

	tr := Insert(nil, min(S, 3*n)*2, min(S, 3*n)*2)

	change := func(x int) {
		// 先找到比x大的值
		node := UpperBound(tr, x)
		// node != nil
		// split [y - v, y] into (y - v, x], (x, y)
		y := node.key
		v := node.val
		tr = Delete(tr, y)
		tr = Insert(tr, y, y-x)
		tr = Insert(tr, x, v-(y-x))
	}

	check := func(x int, sz int) bool {
		if sz > x {
			return false
		}
		return CheckValueBeforeKey(tr, x, sz)
	}

	var ans []bool

	for _, cur := range queries {
		if cur[0] == 1 {
			change(cur[1])
		} else {
			ans = append(ans, check(cur[1], cur[2]))
		}
	}

	return ans
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int
	val         int
	maxValue    int
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

func (node *Node) MaxValue() int {
	if node == nil {
		return 0
	}
	return node.maxValue
}

func NewNode(key, val int) *Node {
	node := new(Node)
	node.key = key
	node.height = 1
	node.cnt = 1
	node.val = val
	node.maxValue = val
	return node
}

func (node *Node) update() {
	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.maxValue = max(node.left.MaxValue(), node.right.MaxValue(), node.val)
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	y.update()
	x.update()

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.update()
	y.update()

	return y
}

func (node *Node) GetBalance() int {
	if node == nil {
		return 0
	}
	return node.left.Height() - node.right.Height()
}

func Insert(node *Node, key, val int) *Node {
	if node == nil {
		return NewNode(key, val)
	}
	if node.key == key {
		node.cnt++
		return node
	}

	if node.key > key {
		node.left = Insert(node.left, key, val)
	} else {
		node.right = Insert(node.right, key, val)
	}

	node.update()

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
			root.val = tmp.val
			root.maxValue = tmp.maxValue
			// make sure tmp node deleted after call delete on root.right
			tmp.cnt = 1
			root.right = Delete(root.right, tmp.key)
		}
	}

	if root == nil {
		return root
	}

	root.update()

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

func UpperBound(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key <= key {
		return UpperBound(node.right, key)
	}
	res := UpperBound(node.left, key)
	if res == nil {
		res = node
	}
	return res
}

func CheckValueBeforeKey(node *Node, key int, val int) bool {
	if node == nil {
		return false
	}
	if key >= node.key {
		if node.val >= val || node.left.MaxValue() >= val {
			return true
		}
		return CheckValueBeforeKey(node.right, key, val)
	}

	// key < node.key
	if key-val >= node.key-node.val {
		return true
	}

	return CheckValueBeforeKey(node.left, key, val)
}
