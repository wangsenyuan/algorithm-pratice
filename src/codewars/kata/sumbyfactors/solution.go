package sumbyfactors

import (
	"bytes"
	"fmt"
	"sort"
)

const MAX_A = 1000000

var LF [MAX_A + 1]int

func init() {
	for x := 2; x < MAX_A; x++ {
		if LF[x] == 0 {
			for y := x; y <= MAX_A; y += x {
				LF[y] = x
			}
		}
	}
}

func SumOfDivided(lst []int) string {
	// your code
	res := make(map[int]int64)
	cnt := make(map[int]int)

	for _, num := range lst {
		cur := abs(num)
		for cur > 1 {
			if cnt[LF[cur]] == 0 {
				res[LF[cur]] += int64(num)
			}
			cnt[LF[cur]]++
			cur /= LF[cur]
		}
		cur = abs(num)
		for cur > 1 {
			cnt[LF[cur]]--
			cur /= LF[cur]
		}
	}
	ps := make([]Pair, 0, len(res))

	for k, v := range res {
		ps = append(ps, Pair{k, v})
	}
	sort.Sort(Pairs(ps))
	var buf bytes.Buffer
	for i := 0; i < len(ps); i++ {
		buf.WriteString(fmt.Sprintf("(%d %d)", ps[i].first, ps[i].second))
	}
	return buf.String()
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type Pair struct {
	first  int
	second int64
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
