package p2816

func minAbsoluteDifference(nums []int, x int) int {
	// abs(i - j) >= x, and min value of abs(nums[i] - nums[j]) 最小
	// 假设 i < j, 对于i，也就是要找到 j >= i + x访问内的和nums[i]最接近的数
	n := len(nums)
	ans := 1 << 30
	// i + x == n - 1
	var root *Node
	for i := n - 1 - x; i >= 0; i-- {
		j := i + x
		// j <= n - 1
		root = Insert(root, nums[j])
		tmp := UpperBound(root, nums[i])
		// tmp.key > nums[i]
		if tmp != nil {
			ans = min(ans, tmp.key-nums[i])
		}
		tmp = LowerBound(root, nums[i])
		// tmp.key <= nums[i]
		if tmp != nil {
			ans = min(ans, nums[i]-tmp.key)
		}
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

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

	if node.key > key {
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

func UpperBound(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key > key {
		// node.key >= key
		tmp := UpperBound(node.left, key)
		if tmp != nil {
			return tmp
		}
		return node
	}
	// node.key <= key
	return UpperBound(node.right, key)
}

func LowerBound(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key <= key {
		tmp := LowerBound(node.right, key)
		if tmp != nil {
			return tmp
		}
		return node
	}
	// node.key > key
	return LowerBound(node.left, key)
}
