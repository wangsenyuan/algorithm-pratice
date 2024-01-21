package p3015

import "sort"

func minimumPushes(word string) int {
	freq := make([]int, 26)

	for i := 0; i < len(word); i++ {
		freq[int(word[i]-'a')]++
	}

	type Pair struct {
		first  int
		second int
	}
	var arr []Pair
	for k, v := range freq {
		arr = append(arr, Pair{v, k})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].first > arr[j].first
	})

	pos := make([]int, 26)
	for i, p := range arr {
		pos[p.second] = i/8 + 1
	}
	var res int

	for k, v := range freq {
		res += pos[k] * v
	}
	return res
}
