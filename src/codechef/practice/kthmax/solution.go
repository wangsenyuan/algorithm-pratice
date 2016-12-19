package main

import (
	"fmt"
	"sort"
)

func main() {
	t := 0
	fmt.Scanf("%d\n", &t)
	for t > 0 {
		t--

		n, m := 0, 0
		fmt.Scanf("%d %d", &n, &m)
		nums := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &nums[i])
		}

		solver := solve(n, nums)
		//fmt.Println(solver.ys)
		for i := 0; i < m; i++ {
			q := int64(0)
			fmt.Scanf("%d", &q)
			fmt.Println(solver.query(q))
		}
	}
}

type Solver struct {
	nums  []int64
	pairs []Pair
	xs    []int64
	ys    []int64
}

func (solver *Solver) query(p int64) int64 {
	n := len(solver.nums)
	l, r := 0, n
	for r-l > 1 {
		mid := (r + l) / 2
		if solver.ys[mid] < p {
			l = mid
		} else {
			r = mid
		}
	}

	return solver.nums[solver.pairs[n-r].idx]
}

func solve(n int, nums []int64) *Solver {
	xs, ys := make([]int64, n+1), make([]int64, n+1)

	pairs := make([]Pair, n)

	for i := 0; i < n; i++ {
		pairs[i] = Pair{i, nums[i]}
	}

	sort.Sort(Pairs(pairs))

	root := insert(nil, -1)

	root = insert(root, n)

	for i := n - 1; i >= 0; i-- {
		pair := pairs[i]
		j := pair.idx

		left := lower(root, j)
		right := higher(root, j)

		xs[n-i] = int64(j-left.val) * int64(right.val-j)

		root = insert(root, j)
	}

	for i := 1; i <= n; i++ {
		ys[i] = ys[i-1] + xs[i]
	}

	return &Solver{nums, pairs, xs, ys}
}

type Pair struct {
	idx int
	val int64
}

type Pairs []Pair

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Less(i, j int) bool {
	if p[i].val < p[j].val {
		return true
	}
	return p[i].idx < p[j].idx
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}

	if val < root.val {
		root.left = insert(root.left, val)
	} else {
		root.right = insert(root.right, val)
	}

	return root
}

func lower(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val > root.val {
		right := lower(root.right, val)
		if right != nil {
			return right
		}
		return root
	}

	return lower(root.left, val)
}

func higher(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val < root.val {
		left := higher(root.left, val)
		if left != nil {
			return left
		}

		return root
	}
	return higher(root.right, val)
}
