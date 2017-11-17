package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d", &n)
		dices := make([][]int, n)
		for j := 0; j < n; j++ {
			var a, b, c, d, e, f int
			fmt.Scanf("%d %d %d %d %d %d", &a, &b, &c, &d, &e, &f)
			dices[j] = []int{a, b, c, d, e, f}
		}
		res := solve(n, dices)
		fmt.Printf("Case #%d: %d\n", i+1, res)
	}
}

func solve(n int, dices [][]int) int {
	numMap := make(map[int]int)

	for _, dice := range dices {
		for _, num := range dice {
			numMap[num] = 1
		}
	}

	nums := make([]int, len(numMap))
	idx := 0
	for num := range numMap {
		nums[idx] = num
		idx++
	}
	sort.Ints(nums)

	graph := make([]map[int]bool, len(nums))

	for i, num := range nums {
		graph[i] = make(map[int]bool)
		numMap[num] = i
	}

	for i, dice := range dices {
		for _, num := range dice {
			j := numMap[num]
			graph[j][i] = true
		}
	}

	var bpm func(v int, flag int) bool
	was := make([]int, len(nums))
	p := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		p[i] = -1
	}

	q := make([]int, n)
	for i := 0; i < n; i++ {
		q[i] = -1
	}

	bpm = func(v int, flag int) bool {
		if was[v] == flag {
			return false
		}
		was[v] = flag
		for w := range graph[v] {
			if q[w] < 0 || bpm(q[w], flag) {
				q[w] = v
				p[v] = w
				return true
			}
		}
		return false
	}
	var ct = 0
	var ans int
	for i, j := 0, 0; i < len(nums); i++ {
		for j < len(nums) && nums[j] == nums[i]+j-i {
			ct++
			if !bpm(j, ct) {
				break
			}
			j++
		}

		if j-i > ans {
			ans = j - i
		}

		q[p[i]] = -1
	}

	return ans
}
