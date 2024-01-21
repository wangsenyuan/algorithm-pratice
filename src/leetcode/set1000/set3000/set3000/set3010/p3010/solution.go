package p3010

func minimumCost(nums []int, k int, dist int) int64 {
	var root *Node

	n := len(nums)

	var best int64 = -1

	for i, r := n-1, n-1; i > 0; i-- {
		if r-i > dist {
			// 超出了距离
			root = Delete(root, Pair{nums[r], r})
			r--
		}

		if root.Size() >= k-2 {
			tmp := PrefixSum(root, k-2) + int64(nums[0]+nums[i])
			if best < 0 || best > tmp {
				best = tmp
			}
		}
		root = Insert(root, Pair{nums[i], i})
	}

	return best
}

/**
* this is a AVL tree
 */
type Node struct {
	item        Pair
	height      int
	size        int
	sum         int64
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

func (node *Node) Sum() int64 {
	if node == nil {
		return 0
	}
	return node.sum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}

func (this Pair) Less(that Pair) bool {
	return this.first < that.first || this.first == that.first && this.second < that.second
}
func NewNode(item Pair) *Node {
	node := new(Node)
	node.item = item
	node.height = 1
	node.size = 1
	node.sum = int64(item.first)
	return node
}

func (node *Node) update() {
	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.size = node.left.Size() + node.right.Size() + 1
	node.sum = node.left.Sum() + node.right.Sum() + int64(node.item.first)
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

func Insert(node *Node, item Pair) *Node {
	if node == nil {
		return NewNode(item)
	}

	if item.Less(node.item) {
		node.left = Insert(node.left, item)
	} else {
		node.right = Insert(node.right, item)
	}

	node.update()

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

func Delete(root *Node, item Pair) *Node {
	if root == nil {
		return nil
	}

	if item.Less(root.item) {
		root.left = Delete(root.left, item)
	} else if root.item.Less(item) {
		root.right = Delete(root.right, item)
	} else {
		if root.left == nil || root.right == nil {
			tmp := root.left
			if root.left == nil {
				tmp = root.right
			}
			root = tmp
		} else {
			tmp := MinValueNode(root.right)

			root.item = tmp.item
			// make sure tmp node deleted after call delete on root.right
			root.right = Delete(root.right, tmp.item)
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

func PrefixSum(node *Node, cnt int) int64 {
	if cnt == 0 {
		return 0
	}
	if node.Size() <= cnt {
		return node.Sum()
	}
	// node.Size() > cnt
	if node.left.Size() >= cnt {
		return PrefixSum(node.left, cnt)
	}
	res := node.left.Sum() + int64(node.item.first)
	res += PrefixSum(node.right, cnt-node.left.Size()-1)

	return res
}
