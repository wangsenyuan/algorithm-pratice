package p1418

import (
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	foods := make(map[string]int)
	mem := make(map[string]map[string]int)

	for _, order := range orders {
		t, f := order[1], order[2]

		if _, found := mem[t]; !found {
			mem[t] = make(map[string]int)
		}
		mem[t][f]++
		foods[f]++
	}

	names := make([]string, 0, len(mem))

	for k := range mem {
		names = append(names, k)
	}

	items := make([]string, 0, len(foods))

	for k := range foods {
		items = append(items, k)
	}
	sort.Sort(Strings(names))
	sort.Strings(items)

	getHeader := func() []string {
		row := make([]string, len(items)+1)
		row[0] = "Table"
		copy(row[1:], items)
		return row
	}

	res := make([][]string, len(names)+1)
	res[0] = getHeader()

	for i := 0; i < len(names); i++ {
		row := make([]string, len(items)+1)
		row[0] = names[i]
		for j := 0; j < len(items); j++ {
			cnt := mem[names[i]][items[j]]
			row[j+1] = strconv.Itoa(cnt)
		}
		res[i+1] = row
	}

	return res
}

type Strings []string

func (ss Strings) Len() int {
	return len(ss)
}

func (ss Strings) Less(i, j int) bool {
	a, _ := strconv.Atoi(ss[i])
	b, _ := strconv.Atoi(ss[j])
	return a < b
}

func (ss Strings) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
