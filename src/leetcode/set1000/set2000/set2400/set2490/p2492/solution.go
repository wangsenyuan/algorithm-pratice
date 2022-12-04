package p2492

func minScore(n int, roads [][]int) int {
	set := NewSet(n)
	for _, road := range roads {
		u, v := road[0], road[1]
		u--
		v--
		set.Union(u, v)
	}

	res := 1 << 30

	for _, road := range roads {
		u, w := road[0], road[2]
		u--
		if set.Find(u) == set.Find(0) {
			if w < res {
				res = w
			}
		}
	}

	return res
}

type Set struct {
	arr []int
	cnt []int
}

func NewSet(n int) *Set {
	s := new(Set)
	s.arr = make([]int, n)
	s.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		s.arr[i] = i
		s.cnt[i] = 1
	}
	return s
}

func (s *Set) Find(a int) int {
	if s.arr[a] != a {
		s.arr[a] = s.Find(s.arr[a])
	}
	return s.arr[a]
}

func (s *Set) Union(a, b int) bool {
	a = s.Find(a)
	b = s.Find(b)
	if a == b {
		return false
	}
	if s.cnt[a] < s.cnt[b] {
		a, b = b, a
	}
	s.cnt[a] += s.cnt[b]
	s.arr[b] = a
	return true
}
