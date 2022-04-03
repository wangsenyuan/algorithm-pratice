package p2225

import "sort"

func findWinners(matches [][]int) [][]int {
	cnt := make(map[int]int)

	for _, m := range matches {
		x := m[1]
		cnt[x]++
	}

	res := make([][]int, 2)
	res[0] = make([]int, 0, 2)
	res[1] = make([]int, 0, 2)

	for _, m := range matches {
		a := m[0]
		if cnt[a] == 0 {
			res[0] = append(res[0], a)
		}
		b := m[1]
		if cnt[b] == 1 {
			res[1] = append(res[1], b)
		}
	}

	sort.Ints(res[0])
	res[0] = removeDuplicate(res[0])
	sort.Ints(res[1])
	res[1] = removeDuplicate(res[1])

	return res
}

func removeDuplicate(arr []int) []int {
	var j int
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i] > arr[i-1] {
			arr[j] = arr[i-1]
			j++
		}
	}
	return arr[:j]
}
