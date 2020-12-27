package p1706

import "sort"

func maximizeXor(nums []int, queries [][]int) []int {
	qs := make([]Query, len(queries))
	for i := 0; i < len(queries); i++ {
		qs[i] = Query{queries[i][0], queries[i][1], i}
	}

	sort.Slice(qs, func(i, j int) bool {
		return qs[i].m < qs[j].m
	})

	sort.Ints(nums)

	root := NewNode()

	ans := make([]int, len(queries))

	var j int
	for _, q := range qs {
		for j < len(nums) && nums[j] <= q.m {
			AddNum(root, nums[j])
			j++
		}
		if j == 0 {
			ans[q.i] = -1
		} else {
			ans[q.i] = FindMax(root, q.x)
		}
	}

	return ans
}

type Node struct {
	children []*Node
}

func NewNode() *Node {
	node := new(Node)
	node.children = make([]*Node, 2)
	return node
}

const H = 30

func AddNum(node *Node, num int) {
	cur := node
	for i := H - 1; i >= 0; i-- {
		j := (num >> i) & 1
		if cur.children[j] == nil {
			cur.children[j] = NewNode()
		}
		cur = cur.children[j]
	}
}

func FindMax(node *Node, num int) int {
	var res int
	cur := node
	for i := H - 1; i >= 0; i-- {
		j := (num >> i) & 1
		if cur.children[j^1] != nil {
			res |= 1 << i
			cur = cur.children[j^1]
		} else {
			cur = cur.children[j]
		}
	}
	return res
}

type Query struct {
	x int
	m int
	i int
}
