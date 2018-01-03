package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] == ' ' {
		i++
	}
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		var n, m int
		pos := readInt(scanner.Bytes(), 0, &n)
		readInt(scanner.Bytes(), pos+1, &m)
		nums := make([]int, n)
		scanner.Scan()
		pos = -1
		for j := 0; j < n; j++ {
			pos = readInt(scanner.Bytes(), pos+1, &nums[j])
		}
		query := make([][]int, m)
		for j := 0; j < m; j++ {
			scanner.Scan()
			var a, b, c int
			pos = readInt(scanner.Bytes(), 0, &a)
			pos = readInt(scanner.Bytes(), pos+1, &b)
			readInt(scanner.Bytes(), pos+1, &c)
			query[j] = []int{a, b, c}
		}
		res := solve(n, nums, query)
		for j := 0; j < len(res); j++ {
			fmt.Println(res[j])
		}
	}
}

func solve(n int, nums []int, query [][]int) []int {
	root := initNode(n, nums)
	ans := make([]int, len(query))
	p := 0
	for i := 0; i < len(query); i++ {
		if query[i][0] == 1 {
			left, right := query[i][1]-1, query[i][2]-1
			x := calculate(1, -(nums[left] + nums[right]))
			y := int(x)

			cand1 := root.query(left, right+1, y)
			cand2 := cand1
			if x-float64(y) >= 0.5 {
				cand2 = root.query(left, right+1, y+1)
			}

			res := (cand1 - nums[left]) * (nums[right] - cand1)
			if cand2 != cand1 {
				tmp := (cand2 - nums[left]) * (nums[right] - cand2)
				if tmp > res {
					res = tmp
				}
			}
			ans[p] = res
			p++
		} else {
			idx, val := query[i][1]-1, query[i][2]
			nums[idx] = val
			root.update(idx, val)
		}
	}
	return ans[:p]
}

func calculate(a, b int) float64 {
	x, y := float64(a), float64(b)
	return -y / (2 * x)
}

type Node struct {
	start int
	end   int
	min   int
	max   int
	left  *Node
	right *Node
}

func initNode(n int, nums []int) *Node {
	var help func(start int, end int) *Node

	help = func(start int, end int) *Node {
		if start+1 == end {
			return &Node{start: start, end: end, min: nums[start], max: nums[start]}
		}

		mid := start + (end-start)/2
		left := help(start, mid)
		right := help(mid, end)
		minVal := min(left.min, right.min)
		maxVal := max(left.max, right.max)
		return &Node{start: start, end: end, min: minVal, max: maxVal, left: left, right: right}
	}

	return help(0, n)
}

func (root *Node) update(i int, x int) {
	var help func(node *Node)
	help = func(node *Node) {
		if node.start == i && node.end == i+1 {
			node.min = x
			node.max = x
			return
		}

		mid := node.start + (node.end-node.start)/2
		if i >= mid {
			help(node.right)
		} else {
			help(node.left)
		}
		// node.val = max(node.left.val, node.right.val)
		node.min = min(node.left.min, node.right.min)
		node.max = max(node.left.max, node.right.max)
	}

	help(root)
}

func (root *Node) query(start int, end int, x int) int {
	var help func(node *Node) (int, bool)
	help = func(node *Node) (int, bool) {
		if node.end <= start || node.start >= end {
			return -1, false
		}

		if node.start >= start && node.end <= end {
			// total in the range
			if x >= node.max || x <= node.min {
				// x in the outside
				if abs(x-node.max) >= abs(x-node.min) {
					return node.min, true
				}
				return node.max, true
			}
		}

		leftRes, leftCan := help(node.left)
		rightRes, rightCan := help(node.right)
		if !leftCan {
			return rightRes, rightCan
		}
		if !rightCan {
			return leftRes, leftCan
		}
		if abs(x-leftRes) >= abs(x-rightRes) {
			return rightRes, true
		}
		return leftRes, true
	}
	res, _ := help(root)
	return res
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
