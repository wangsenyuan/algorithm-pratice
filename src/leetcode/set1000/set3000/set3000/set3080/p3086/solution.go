package p3086

import "sort"

func minimumDeletions(word string, k int) int {
	freq := make([]int, 26)

	for i := 0; i < len(word); i++ {
		freq[int(word[i]-'a')]++
	}
	sort.Ints(freq)

	for i := 0; i < len(freq); i++ {
		if freq[i] > 0 {
			freq = freq[i:]
			break
		}
	}
	best := len(word)

	// 现在删除，应该是删除部分最大值，且最小值完全删除掉
	var prev int

	n := len(freq)
	for i := 0; i < n; i++ {
		var suf int
		for j := n - 1; j > i; j-- {
			suf += max(0, freq[j]-(freq[i]+k))
		}

		best = min(best, prev+suf)
		prev += freq[i]
	}

	return best
}
