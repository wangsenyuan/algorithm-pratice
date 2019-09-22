package p1202

import (
	"bytes"
	"sort"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	set := NewUFSet(n)

	for _, pair := range pairs {
		set.Union(pair[0], pair[1])
	}

	groups := make(map[int][]int)

	for i := 0; i < n; i++ {
		pi := set.Find(i)
		if _, found := groups[pi]; !found {
			groups[pi] = make([]int, 0, 10)
		}
		groups[pi] = append(groups[pi], i)
	}

	res := []byte(s)

	var buf bytes.Buffer

	for _, group := range groups {
		buf.Reset()

		for i := 0; i < len(group); i++ {
			buf.WriteByte(s[group[i]])
		}
		bs := buf.Bytes()

		sort.Sort(SortBytes(bs))

		for i := 0; i < len(group); i++ {
			res[group[i]] = bs[i]
		}
	}

	return string(res)
}

type SortBytes []byte

func (ss SortBytes) Len() int {
	return len(ss)
}

func (ss SortBytes) Less(i, j int) bool {
	return ss[i] < ss[j]
}

func (ss SortBytes) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

type UF struct {
	arr []int
	cnt []int
}

func NewUFSet(n int) UF {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return UF{arr, cnt}
}

func (this *UF) Find(a int) int {
	if this.arr[a] != a {
		this.arr[a] = this.Find(this.arr[a])
	}
	return this.arr[a]
}

func (this *UF) Union(a, b int) {
	pa := this.Find(a)
	pb := this.Find(b)
	if pa != pb {
		if this.cnt[pa] > this.cnt[pb] {
			this.arr[pb] = pa
			this.cnt[pa] += this.cnt[pb]
		} else {
			this.arr[pa] = pb
			this.cnt[pb] += this.cnt[pa]
		}
	}
}
