package c

const MOD = 1000000007

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	if a < 0 {
		a += MOD
	}
	return a
}

func numsGame(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	var root *Node
	var sum int
	for i := 0; i < n; i++ {
		x := nums[i] - i
		sum = sum + x
		root = Insert(root, x)
		if i == 0 {
			continue
		}
		h := (i + 1) / 2
		y := root.GetSum(h)
		z := sum - y
		if (i+1)&1 == 1 {
			//odd cnt
			z = sum - root.GetSum(h+1)
		}
		res[i] = (z - y) % MOD
		if res[i] < 0 {
			res[i] += MOD
		}
	}

	return res
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int
	height      int
	sum         int
	size        int
	left, right *Node
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func NewNode(key int) *Node {
	node := new(Node)
	node.key = key
	node.height = 1
	node.size = 1
	node.sum = key
	return node
}

func (node *Node) Size() int {
	if node == nil {
		return 0
	}
	return node.size
}

func (node *Node) Sum() int {
	if node == nil {
		return 0
	}
	return node.sum
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func update(x *Node) {
	x.height = max(x.left.Height(), x.right.Height()) + 1
	x.size = x.left.Size() + x.right.Size() + 1
	x.sum = x.key + x.left.Sum() + x.right.Sum()
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2

	update(y)
	update(x)

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	update(x)
	update(y)

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

	if node.key > key {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
	}

	update(node)

	balance := node.GetBalance()

	if balance > 1 && key <= node.left.key {
		return rightRotate(node)
	}

	if balance < -1 && key >= node.right.key {
		return leftRotate(node)
	}

	if balance > 1 && key >= node.left.key {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && key <= node.right.key {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func (node *Node) GetSum(cnt int) int {
	if cnt == 0 {
		return 0
	}
	if node.Size() == cnt {
		return node.Sum()
	}
	// node.size > cnt
	if node.left.Size() >= cnt {
		return node.left.GetSum(cnt)
	}
	// node.left.size < cnt
	var res = node.left.Sum()
	res += node.key
	res += node.right.GetSum(cnt - node.left.Size() - 1)
	return res
}
