package p767

import (
	"sort"
)

func reorganizeString(S string) string {
	n := len(S)
	if n == 0 {
		return ""
	}

	cnt := make([]int, 26)
	var maj int
	for i := 0; i < n; i++ {
		cnt[S[i]-'a']++
		if cnt[S[i]-'a'] > maj {
			maj = cnt[S[i]-'a']
		}
	}

	if n%2 == 0 && maj > n/2 {
		return ""
	}
	if n%2 == 1 && maj > n/2+1 {
		return ""
	}

	items := make(Items, 26)
	for i := 0; i < 26; i++ {
		items[i] = Item{letter: byte(i + 'a'), cnt: cnt[i]}
	}

	sort.Sort(items)

	ans := make([]byte, n)

	i := 0
	for j := 25; j >= 0 && items[j].cnt > 0; j-- {
		for k := 0; k < items[j].cnt; k++ {
			ans[i] = items[j].letter
			i += 2
			if i >= n {
				i = 1
			}
		}
	}
	return string(ans)
}

type Item struct {
	letter byte
	cnt    int
}

type Items []Item

func (this Items) Len() int {
	return len(this)
}

func (this Items) Less(i, j int) bool {
	return this[i].cnt < this[j].cnt
}

func (this Items) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
