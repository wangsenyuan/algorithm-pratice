package p2424

func numberOfPairs(nums1 []int, nums2 []int, diff int) int64 {
	// nums1[i] - nums1[j] <= nums2[i] - nums2[j] + diff
	// nums[i] - nums2[i] <= nums1[j] - nums2[j] + diff
	n := len(nums1)

	A := make([]int, n)

	for i := 0; i < n; i++ {
		A[i] = nums1[i] - nums2[i]
	}
	// find (i, j) A[i] <= A[j] + diff, (i < j)
	var res int64

	var root *Node

	for i := n - 1; i >= 0; i-- {
		cnt := CountKey(root, A[i]-diff)
		res += int64(cnt)
		root = Insert(root, A[i])
	}

	return res
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

func (node *Node) Count() int {
	if node == nil {
		return 0
	}
	return node.cnt
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

func NewNode(key int) *Node {
	node := new(Node)
	node.key = key
	node.height = 1
	node.cnt = 1
	node.size = 1
	return node
}

func (node *Node) Update() {
	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.size = node.left.Size() + node.right.Size() + node.cnt
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	y.Update()
	x.Update()

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.Update()

	y.Update()

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

	node.Update()

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

func CountKey(root *Node, key int) int {
	if root == nil {
		return 0
	}
	if root.key < key {
		return CountKey(root.right, key)
	}
	// root.key >= key
	return CountKey(root.left, key) + root.cnt + root.right.Size()
}
