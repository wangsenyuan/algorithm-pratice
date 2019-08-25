package p1170

import "sort"

func numSmallerByFrequency(queries []string, words []string) []int {
	wf := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		wf[i] = freq(words[i])
	}

	sort.Ints(wf)

	ans := make([]int, len(queries))

	for i := 0; i < len(ans); i++ {
		f := freq(queries[i])
		j := sort.Search(len(wf), func(j int) bool {
			return wf[j] > f
		})

		ans[i] = len(wf) - j
	}
	return ans
}

func freq(s string) int {
	x := s[0]
	cnt := 1
	for i := 1; i < len(s); i++ {
		if s[i] < x {
			x = s[i]
			cnt = 1
		} else if s[i] == x {
			cnt++
		}
	}
	return cnt
}
