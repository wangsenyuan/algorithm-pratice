package lcp60

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMaxLayerSum(root *TreeNode) int {

	var levels [][]*State

	var collect func(d int, node *TreeNode) int

	var remove [][]int

	collect = func(d int, node *TreeNode) int {
		if node == nil {
			return 0
		}
		for d+1 >= len(levels) {
			levels = append(levels, []*State{NewState()})
		}

		lst := last(levels[d])
		cur := State{lst.sum + node.Val, -1, getSize(levels, d+1), -1}
		levels[d] = append(levels[d], &cur)
		if collect(d+1, node.Left)+collect(d+1, node.Right) != 2 {
			remove = append(remove, []int{getSize(levels, d) - 1, d})
		}
		cur.right = getSize(levels, d+1) - 1
		return 1
	}

	collect(0, root)

	height := len(levels)
	res := last(levels[0]).sum

	for i := 1; i < height-1; i++ {
		res = max(res, last(levels[i]).sum)
	}

	for i := 0; i < len(remove); i++ {
		node, start := remove[i][0], remove[i][1]
		left, right := node, node
		lost := levels[start][left].sum - levels[start][left-1].sum
		for level := start; level < height; level++ {
			if left > right {
				break
			}
			if right-left+1 == getSize(levels, level)-1 {
				break
			}
			a := levels[level][left]
			b := levels[level][right]
			if a.flag != -1 && a.flag == b.flag {
				break
			}
			a.flag = i
			b.flag = i
			var add int
			if a.left <= b.right {
				add = getAt(levels, level+1, b.right).sum - getAt(levels, level+1, a.left-1).sum
			}
			res = max(res, last(levels[level]).sum-lost+add)
			left = a.left
			right = b.right
			lost = add
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type State struct {
	sum   int
	flag  int
	left  int
	right int
}

func getAt(arr [][]*State, i, j int) *State {
	if i < 0 || len(arr) <= i {
		return NewState()
	}
	if j < 0 || len(arr[i]) <= j {
		return NewState()
	}

	return arr[i][j]
}

func last(arr []*State) *State {
	n := len(arr)
	return arr[n-1]
}

func getSize(arr [][]*State, i int) int {
	if len(arr) <= i {
		return 0
	}
	return len(arr[i])
}

func NewState() *State {
	return &State{0, -1, -1, -1}
}
