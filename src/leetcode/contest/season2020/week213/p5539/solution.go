package p5539

import "sort"

func frequencySort(nums []int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}
	ps := make([]Pair, 0, len(freq))
	for k, v := range freq {
		ps = append(ps, Pair{v, k})
	}

	sort.Sort(Pairs(ps))

	res := make([]int, 0, len(nums))

	for i := 0; i < len(ps); i++ {
		for j := 0; j < ps[i].first; j++ {
			res = append(res, ps[i].second)
		}
	}
	return res
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first || ps[i].first == ps[j].first && ps[i].second > ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
