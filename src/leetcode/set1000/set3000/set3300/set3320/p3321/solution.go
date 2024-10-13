package p3321

func findXSum(nums []int, k int, x int) []int64 {
	freq := make(map[int]int)
	var tr *Node

	add := func(num int) {
		if freq[num] > 0 {
			tr = Delete(tr, Pair{freq[num], num})
		}
		freq[num]++
		tr = Insert(tr, Pair{freq[num], num})
	}

	rem := func(num int) {
		tr = Delete(tr, Pair{freq[num], num})
		freq[num]--
		if freq[num] > 0 {
			tr = Insert(tr, Pair{freq[num], num})
		}
	}

	var ans []int64
	for i, num := range nums {
		if i >= k {
			rem(nums[i-k])
		}
		add(num)
		if i >= k-1 {
			ans = append(ans, int64(Get(tr, x)))
		}
	}
	return ans
}

type Pair struct {
	first  int
	second int
}

func (this Pair) Less(that Pair) bool {
	return this.first < that.first || this.first == that.first && this.second < that.second
}

/**
* this is a AVL tree
 */
type Node struct {
	key         Pair
	sum         int
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

func NewNode(key Pair) *Node {
	node := new(Node)
	node.key = key
	node.height = 1
	node.cnt = 1
	node.sum = key.first * key.second
	return node
}

func (node *Node) Sum() int {
	if node == nil {
		return 0
	}
	return node.sum
}

func (node *Node) Count() int {
	if node == nil {
		return 0
	}
	return node.cnt
}

func (node *Node) update() {
	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.sum = node.left.Sum() + node.right.Sum() + node.key.first*node.key.second
	// num 表示子树中一共有多少个节点
	node.cnt = node.left.Count() + node.right.Count() + 1
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

func Insert(node *Node, key Pair) *Node {
	if node == nil {
		return NewNode(key)
	}

	if key.Less(node.key) {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
	}

	node.update()

	balance := node.GetBalance()

	if balance > 1 && key.Less(node.left.key) {
		return rightRotate(node)
	}

	if balance < -1 && node.right.key.Less(key) {
		return leftRotate(node)
	}

	if balance > 1 && node.left.key.Less(key) {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && key.Less(node.right.key) {
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

func Delete(root *Node, key Pair) *Node {
	if root == nil {
		return nil
	}

	if key.Less(root.key) {
		root.left = Delete(root.left, key)
	} else if root.key.Less(key) {
		root.right = Delete(root.right, key)
	} else {
		// equal
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
			root.sum = tmp.sum
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

func Get(node *Node, k int) int {
	if node == nil || k == 0 {
		return 0
	}
	if k >= node.cnt {
		return node.sum
	}
	if node.right.Count() >= k {
		return Get(node.right, k)
	}
	res := node.right.Sum() + node.key.first*node.key.second
	k -= node.right.Count() + 1
	return res + Get(node.left, k)
}
