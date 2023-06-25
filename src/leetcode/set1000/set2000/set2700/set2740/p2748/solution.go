package main

import "sort"

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	n := len(positions)
	robots := make([]Robot, n)
	for i := 0; i < n; i++ {
		robots[i] = Robot{positions[i], healths[i], directions[i], i}
	}

	sort.Slice(robots, func(i, j int) bool {
		return robots[i].pos < robots[j].pos
	})

	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = -1
	}
	stack := make([]int, n)

	process := func(iter func(func(int)), x byte) {
		var p int
		iter(func(i int) {
			if robots[i].dir == x {
				stack[p] = robots[i].hp
				p++
			} else {
				// p > 0
				cur := robots[i].hp
				for p > 0 && cur > 0 && stack[p-1] < cur {
					cur--
					p--
				}
				if p == 0 {
					res[robots[i].id] = cur
					return
				}
				// p > 0
				if stack[p-1] == cur {
					p--
					return
				}
				// stack[p-1] > cur
				stack[p-1]--
			}
		})
	}

	process(func(f func(int)) {
		for i := 0; i < n; i++ {
			f(i)
		}
	}, 'R')

	process(func(f func(int)) {
		for i := n - 1; i >= 0; i-- {
			f(i)
		}
	}, 'L')

	var ans []int
	for i := 0; i < n; i++ {
		if res[i] > 0 {
			ans = append(ans, res[i])
		}
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Robot struct {
	pos int
	hp  int
	dir byte
	id  int
}
