package main

import (
	"fmt"
	"sort"
)

func rearrangeString(str string, k int) string {
	if k <= 1 {
		return str
	}

	mp := make(map[rune]*cnt)

	for _, r := range str {
		if v, found := mp[r]; found {
			v.count++
		} else {
			mp[r] = &cnt{1, r}
		}
	}

	as := make([]*cnt, 0, len(mp))

	for _, a := range mp {
		as = append(as, a)
	}

	sort.Sort(sort.Reverse(cnts(as)))

	res := make([]rune, len(str), len(str))
	index := 0
	cycles := 0
	for _, a := range as {
		for i := 0; i < a.count; i++ {
			res[index] = a.index
			if index > 0 && res[index] == res[index-1] {
				return ""
			}
			index += k

			if index >= len(str) {
				cycles++
				index = cycles
			}
		}
	}

	return string(res)
}

type cnt struct {
	count int
	index rune
}

type cnts []*cnt

func (a cnts) Len() int {
	return len(a)
}

func (a cnts) Less(i, j int) bool {
	if a[i].count == a[j].count {
		return a[i].index < a[j].index
	}

	return a[i].count < a[j].count
}

func (a cnts) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	fmt.Println(rearrangeString("aabbcc", 3))
	fmt.Println(rearrangeString("bbabcaccaaabababbaaaaccbbcbacbacacccbbcaabcbcacaaccbabbbbbcacccaccbabaccbacabcabcacb", 2))
}
