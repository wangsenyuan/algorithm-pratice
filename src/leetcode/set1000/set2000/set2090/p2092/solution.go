package p2092

import "sort"

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	// n <= 10 ^^ 5
	// len(meetings) <= 10  ^^ 5
	set := NewUF(n)
	set.Union(0, firstPerson)
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	for i := 0; i < len(meetings); {
		j := i
		// 需要保证先收到消息
		for i < len(meetings) && meetings[i][2] == meetings[j][2] {
			u, v := meetings[i][0], meetings[i][1]
			set.Union(u, v)
			i++
		}

		i = j

		for i < len(meetings) && meetings[i][2] == meetings[j][2] {
			u, v := meetings[i][0], meetings[i][1]
			pu := set.Find(u)
			pf := set.Find(firstPerson)
			if pu != pf {
				// need to reset it
				set.Reset(u)
				set.Reset(v)
			}

			i++
		}
	}
	var res []int

	for i := 0; i < n; i++ {
		pi := set.Find(i)
		pf := set.Find(firstPerson)
		if pi == pf {
			res = append(res, i)
		}
	}

	return res
}

type UF struct {
	arr []int
	cnt []int
}

func NewUF(n int) *UF {
	set := new(UF)
	set.arr = make([]int, n)
	set.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		set.arr[i] = i
		set.cnt[i] = 1
	}
	return set
}

func (set *UF) Find(x int) int {
	if set.arr[x] != x {
		set.arr[x] = set.Find(set.arr[x])
	}
	return set.arr[x]
}

func (set *UF) Union(a, b int) {
	pa := set.Find(a)
	pb := set.Find(b)
	if pa == pb {
		return
	}
	if set.cnt[pa] < set.cnt[pb] {
		pa, pb = pb, pa
	}
	set.cnt[pa] += set.cnt[pb]
	set.arr[pb] = pa
}

func (set *UF) Reset(x int) {
	set.arr[x] = x
	set.cnt[x] = 1
}
