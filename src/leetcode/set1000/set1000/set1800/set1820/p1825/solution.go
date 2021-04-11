package p1825

type MKAverage struct {
	root *Node
	m    int
	k    int
	nums []int
}

func Constructor(m int, k int) MKAverage {
	nums := make([]int, 0, 100010)
	return MKAverage{nil, m, k, nums}
}

func (this *MKAverage) AddElement(num int) {
	this.nums = append(this.nums, num)
	this.root = Insert(this.root, num)
	n := len(this.nums)
	if n > this.m {
		this.root = Delete(this.root, this.nums[n-this.m-1])
	}
}

func (this *MKAverage) CalculateMKAverage() int {
	n := len(this.nums)
	if n < this.m {
		return -1
	}
	s1 := FirstKSum(this.root, this.k)
	s2 := FirstKSum(this.root, this.m-this.k)
	s := int(s2 - s1)
	return s / (this.m - 2*this.k)
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int
	height      int
	cnt         int
	sum         int64
	sz          int
	left, right *Node
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (node *Node) Sum() int64 {
	if node == nil {
		return 0
	}
	return node.sum
}

func (node *Node) Size() int {
	if node == nil {
		return 0
	}
	return node.sz
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
	node.sum = int64(key)
	node.sz = 1
	return node
}

func (node *Node) update() {
	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.sum = node.left.Sum() + node.right.Sum() + int64(node.key)*int64(node.cnt)
	node.sz = node.left.Size() + node.right.Size() + node.cnt
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

func Insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}
	if node.key == key {
		node.cnt++
		node.sz++
		node.sum += int64(key)
		return node
	}

	if node.key > key {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
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
		root.sum -= int64(key)
		root.sz--
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
			//root.sum = tmp.key
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

func FirstKSum(root *Node, k int) int64 {
	if k == 0 {
		return 0
	}
	if root.left.Size() >= k {
		return FirstKSum(root.left, k)
	}

	res := root.left.Sum()

	if root.cnt >= k-root.left.Size() {
		res += int64(root.key) * int64(k-root.left.Size())
		return res
	}
	res += int64(root.key) * int64(root.cnt)
	// need some from right
	res += FirstKSum(root.right, k-root.left.Size()-root.cnt)
	return res
}

/**
 * Your MKAverage object will be instantiated and called as such:
 * obj := Constructor(m, k);
 * obj.AddElement(num);
 * param_2 := obj.CalculateMKAverage();
 */
