package p1258

import (
	"sort"
	"strings"
)

func generateSentences(synonyms [][]string, text string) []string {
	ii := make(map[string]int)
	si := make(map[int]string)
	var index int

	addWord := func(word string) {
		if _, found := ii[word]; !found {
			ii[word] = index
			si[index] = word
			index++
		}
	}

	for _, syn := range synonyms {
		addWord(syn[0])
		addWord(syn[1])
	}
	n := len(ii)
	uf := NewUF(n)

	for _, syn := range synonyms {
		a := syn[0]
		b := syn[1]
		uf.Union(ii[a], ii[b])
	}
	pos := make(map[int]int)
	roots := uf.GetRoots()
	ss := make([][]string, len(roots))
	for i := 0; i < len(roots); i++ {
		pos[roots[i]] = i
		set := uf.GetSet(roots[i])
		ss[i] = make([]string, len(set))
		for j, k := range set {
			ss[i][j] = si[k]
		}
		sort.Strings(ss[i])
	}

	words := strings.Split(text, " ")
	m := len(words)

	buf := make([]string, m)

	var dfs func(i int)
	var res []string
	dfs = func(i int) {
		if i == m {
			res = append(res, strings.Join(buf, " "))
			return
		}

		if j, found := ii[words[i]]; found {
			// have syns
			pj := uf.Find(j)
			r := pos[pj]
			xs := ss[r]

			for _, x := range xs {
				buf[i] = x
				dfs(i + 1)
			}

			return
		}

		buf[i] = words[i]
		dfs(i + 1)
	}

	dfs(0)

	return res
}

type UF struct {
	arr []int
	cnt []int
}

func NewUF(n int) UF {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return UF{arr, cnt}
}

func (uf *UF) Find(x int) int {
	if uf.arr[x] != x {
		uf.arr[x] = uf.Find(uf.arr[x])
	}
	return uf.arr[x]
}

func (uf *UF) Union(a, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		return false
	}
	if uf.cnt[pa] >= uf.cnt[pb] {
		uf.arr[pb] = pa
		uf.cnt[pa] += uf.cnt[pb]
	} else {
		uf.arr[pa] = pb
		uf.cnt[pb] += uf.cnt[pa]
	}
	return true
}

func (uf UF) GetRoots() []int {
	var res []int
	for i := 0; i < len(uf.arr); i++ {
		if uf.arr[i] == i {
			res = append(res, i)
		}
	}
	return res
}

func (uf *UF) GetSet(x int) []int {
	px := uf.Find(x)
	var res []int
	for i := 0; i < len(uf.arr); i++ {
		if uf.arr[i] == px {
			res = append(res, i)
		}
	}
	return res
}
