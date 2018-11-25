package p947

func removeStones(stones [][]int) int {
	n := len(stones)
	set := CreateSet(n)
	for i := 0; i < n; i++ {
		x, y := stones[i][0], stones[i][1]
		for j := i + 1; j < n; j++ {
			u, v := stones[j][0], stones[j][1]
			if x == u || y == v {
				set.Union(i, j)
			}
		}
	}

	keys := set.Keys()
	var res int

	for i := 0; i < len(keys); i++ {
		res += set.Size(keys[i]) - 1
	}
	return res
}

type Set struct {
	arr  []int
	size []int
}

func CreateSet(n int) *Set {
	set := new(Set)
	set.arr = make([]int, n)
	set.size = make([]int, n)
	for i := 0; i < n; i++ {
		set.arr[i] = i
		set.size[i] = 1
	}
	return set
}

func (set *Set) Find(x int) int {
	var loop func(x int) int
	loop = func(x int) int {
		if set.arr[x] == x {
			return x
		}
		set.arr[x] = loop(set.arr[x])
		return set.arr[x]
	}
	return loop(x)
}

func (set *Set) Union(a, b int) bool {
	pa := set.Find(a)
	pb := set.Find(b)
	if pa == pb {
		return false
	}
	if set.size[pa] >= set.size[pb] {
		set.arr[pb] = pa
		set.size[pa] += set.size[pb]
	} else {
		set.arr[pa] = pb
		set.size[pb] += set.size[pa]
	}
	return true
}

func (set *Set) Size(a int) int {
	pa := set.Find(a)
	return set.size[pa]
}

func (set Set) Keys() []int {
	n := len(set.arr)
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if set.arr[i] == i {
			res = append(res, i)
		}
	}
	return res
}
