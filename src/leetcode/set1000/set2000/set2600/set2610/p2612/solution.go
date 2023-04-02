package p2612

func minReverseOperations(n int, p int, banned []int, k int) []int {
	// n <= 1e5
	ban := make([]bool, n)
	for _, i := range banned {
		ban[i] = true
	}

	// 假设当前 i在dp[i]次操作后被设置位1
	// 如果某个点j满足条件 ((k - 1) - (j - i) ) % 2 == 0
	// 那么 dp[j] = dp[i]+1
	// 也就是说 	k - 1 和 j - i需要相同的奇偶性
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}

	trs := make([]*Node, 2)

	for i := 0; i < n; i++ {
		if i != p && !ban[i] {
			trs[i%2] = Insert(trs[i%2], i)
		}
	}

	trs[0] = Insert(trs[0], n)
	trs[1] = Insert(trs[1], n)

	que := make([]int, n)
	var front, end int

	que[end] = p
	end++

	for step := 0; front < end; step++ {
		cur := end
		for front < cur {
			i := que[front]
			front++
			res[i] = step
			mn := max(i+k-(i*2+1), i-k+1)
			mx := min(i-k+((n-1-i)*2+1), i+k-1)
			j := mn % 2
			for {
				node := Search(trs[j], mn)

				if node == nil || node.key > mx {
					break
				}
				que[end] = node.key
				end++
				trs[j] = Delete(trs[j], node.key)
			}
		}
	}

	return res
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

func Search(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key < key {
		return Search(node.right, key)
	}
	// node.key >= key
	res := Search(node.left, key)
	if res != nil {
		return res
	}
	return node
}
