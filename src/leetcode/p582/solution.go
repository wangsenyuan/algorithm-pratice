package main

import "fmt"

func main() {
	pid := []int{1, 3, 10, 5}
	ppid := []int{3, 0, 5, 3}
	kill := 5
	fmt.Println(killProcess(pid, ppid, kill))
}

func killProcess(pid []int, ppid []int, kill int) []int {
	idx := make(map[int]int)

	for i, p := range pid {
		idx[p] = i
	}

	parent := make(map[int]int)

	var connect func(cur int) int

	connect = func(cur int) int {
		if cur == kill || cur == 0 {
			parent[cur] = cur
			return cur
		}
		if pp, found := parent[cur]; found {
			return pp
		}

		pp := ppid[idx[cur]]
		ppp := connect(pp)
		parent[cur] = ppp
		return ppp
	}

	for _, cur := range pid {
		connect(cur)
	}

	res := make([]int, 0, len(pid))

	for _, cur := range pid {
		if parent[cur] == kill || cur == kill {
			res = append(res, cur)
		}
	}
	if kill == 0 {
		res = append(res, kill)
	}
	return res
}
