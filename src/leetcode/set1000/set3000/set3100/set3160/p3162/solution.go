package p3162

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func maximumSumSubsequence(nums []int, queries [][]int) int {
	tr := Build(nums)

	n := len(nums)

	var res int

	for _, cur := range queries {
		i, x := cur[0], cur[1]
		nums[i] = x
		tr.Update(i, x, n)
		tmp := tr.f11 % mod
		res = add(res, tmp)
	}

	return res
}

const inf = 1 << 60

// f00 表示第一个数一定不选，最后一个数一定不选
// f01 表示第一个数一定不选，最后一个数可选可不选
// f10 表示第一个数可选可不选，最后一个数一定不选
// f11 表示第一个数可选可不选，最后一个数可选可不选，也就是没有任何限制

type Node struct {
	f00         int
	f01         int
	f10         int
	f11         int
	left, right *Node
}

func (node *Node) merge() {
	node.f00 = max(node.left.f00+node.right.f10, node.left.f01+node.right.f00)
	node.f01 = max(node.left.f00+node.right.f11, node.left.f01+node.right.f01)
	node.f10 = max(node.left.f10+node.right.f10, node.left.f11+node.right.f00)
	node.f11 = max(node.left.f10+node.right.f11, node.left.f11+node.right.f01)
}

func Build(arr []int) *Node {
	var loop func(l int, r int) *Node

	loop = func(l int, r int) *Node {
		node := new(Node)
		if l == r {
			node.f11 = max(0, arr[l])
			return node
		}
		mid := (l + r) / 2
		node.left = loop(l, mid)
		node.right = loop(mid+1, r)
		node.merge()
		return node
	}
	return loop(0, len(arr)-1)
}

func (root *Node) Update(p int, v int, n int) {
	var loop func(node *Node, l int, r int)
	loop = func(node *Node, l int, r int) {
		if l == r {
			node.f11 = max(0, v)
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(node.left, l, mid)
		} else {
			loop(node.right, mid+1, r)
		}
		node.merge()
	}

	loop(root, 0, n-1)
}
