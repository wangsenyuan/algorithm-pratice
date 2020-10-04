package p1612

import (
	"sort"
	"strconv"
	"strings"
)

func alertNames(keyName []string, keyTime []string) []string {
	n := len(keyName)
	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{keyName[i], parseAsTime(keyTime[i])}
	}

	sort.Sort(Pairs(ps))

	mem := make(map[string]int)

	res := make(map[string]bool)
	var j int
	for i := 0; i < n; i++ {
		cur := ps[i]
		mem[cur.first]++
		for j < i && ps[j].second+60 < cur.second {
			mem[ps[j].first]--
			j++
		}

		if mem[cur.first] > 2 {
			res[cur.first] = true
		}
	}
	ans := make([]string, 0, len(res))

	for k := range res {
		ans = append(ans, k)
	}
	sort.Strings(ans)
	return ans
}

func parseAsTime(s string) int {
	ss := strings.Split(s, ":")
	h, _ := strconv.Atoi(ss[0])
	m, _ := strconv.Atoi(ss[1])

	return h*60 + m
}

type Pair struct {
	first  string
	second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].second < ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
