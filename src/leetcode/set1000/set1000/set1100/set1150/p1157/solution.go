package p1157

import "sort"

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].second > this[j].second
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type Node struct {
	nums Pairs
}

type MajorityChecker struct {
	tree []*Node
	size int
}

func merge(a, b *Node) *Node {
	if a == nil && b == nil {
		return nil
	}
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	nums_a := a.nums
	nums_b := b.nums

	cnt := make(map[int]int)

	for i := 0; i < len(nums_a); i++ {
		cnt[nums_a[i].first] = nums_a[i].second
	}

	for i := 0; i < len(nums_b); i++ {
		cnt[nums_b[i].first] += nums_b[i].second
	}

	nums := make([]Pair, 0, 6)

	for k, v := range cnt {
		nums = append(nums, Pair{k, v})
	}

	sort.Sort(Pairs(nums))

	if len(nums) > 3 {
		nums = nums[:5]
	}

	return &Node{nums}
}

func Constructor(arr []int) MajorityChecker {
	n := len(arr)
	tree := make([]*Node, 4*n)

	var loop func(i int, left int, right int)

	loop = func(i int, left int, right int) {
		if left == right {
			node := new(Node)
			node.nums = Pairs(make([]Pair, 1))
			node.nums[0] = Pair{arr[left], 1}
			tree[i] = node
			return
		}
		mid := (left + right) / 2
		loop(2*i+1, left, mid)
		loop(2*i+2, mid+1, right)
		tree[i] = merge(tree[2*i+1], tree[2*i+2])
	}

	loop(0, 0, n-1)

	return MajorityChecker{tree, n}
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	tree := this.tree

	var loop func(i int, start int, end int) *Node

	loop = func(i int, start int, end int) *Node {
		if end < left || right < start {
			return nil
		}
		if left <= start && end <= right {
			return tree[i]
		}
		mid := (start + end) / 2
		a := loop(2*i+1, start, mid)
		b := loop(2*i+2, mid+1, end)

		return merge(a, b)
	}

	res := loop(0, 0, this.size-1)

	if res == nil {
		return -1
	}

	theNum := res.nums[0]

	if theNum.second >= threshold {
		return theNum.first
	}
	return -1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */
