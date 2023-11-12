package p2931

import "sort"

const H = 31

func maximumStrongPairXor(nums []int) int {
	sort.Ints(nums)
	// y - x <= x
	// y <= 2 * x
	// 对于y来说，就是要找到y/2到y之间的x

	tr := new(Node)

	remove := func(num int) {
		node := tr
		for i := H - 1; i >= 0; i-- {
			x := (num >> i) & 1
			node.children[x].cnt--
			node = node.children[x]
		}
	}

	find := func(num int) int {
		node := tr
		var res int
		for i := H - 1; i >= 0 && node != nil; i-- {
			x := (num >> i) & 1
			if node.children[1^x] != nil && node.children[1^x].cnt > 0 {
				res ^= 1 << i
				node = node.children[1^x]
			} else {
				node = node.children[x]
			}
		}

		return res
	}

	add := func(num int) {
		node := tr
		for i := H - 1; i >= 0; i-- {
			x := (num >> i) & 1
			if node.children[x] == nil {
				node.children[x] = new(Node)
			}
			node.children[x].cnt++
			node = node.children[x]
		}
	}

	var best int
	n := len(nums)
	for l, r := 0, 0; r < n; r++ {
		for l < r && nums[l]*2 < nums[r] {
			remove(nums[l])
			l++
		}
		best = max(best, find(nums[r]))
		add(nums[r])
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Node struct {
	cnt      int
	children [2]*Node
}
