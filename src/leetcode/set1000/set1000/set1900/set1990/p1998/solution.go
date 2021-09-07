package p1998

import "sort"

const X = 100002

var pf []int

func init() {
	pf = make([]int, X)

	for x := 2; x < X; x++ {
		if pf[x] == 0 {
			for y := x; y < X; y += x {
				if pf[y] == 0 {
					pf[y] = x
				}
			}
		}
	}
}

func gcdSort(nums []int) bool {
	pos := make([]int, X)
	n := len(nums)
	set := NewUFSet(n)
	for i, num := range nums {
		tmp := num
		for tmp > 1 {
			j := pos[pf[tmp]]
			if j > 0 {
				set.Union(i, j-1)
			}
			tmp /= pf[tmp]
		}

		for num > 1 {
			pos[pf[num]] = i + 1
			num /= pf[num]
		}
	}

	ps := make(map[int][]int)

	for i := 0; i < n; i++ {
		p := set.Find(i)
		sz := set.size[p]
		if _, found := ps[p]; !found {
			ps[p] = make([]int, 0, sz)
		}
		ps[p] = append(ps[p], i)
	}
	arr := make([]int, n)

	for _, v := range ps {
		w := make([]int, len(v))
		copy(w, v)
		sort.Slice(v, func(i, j int) bool {
			a, b := v[i], v[j]
			return nums[a] <= nums[b]
		})
		for i := 0; i < len(v); i++ {
			arr[w[i]] = nums[v[i]]
		}
	}

	return sort.IntsAreSorted(arr)
}

type UFSet struct {
	arr  []int
	size []int
	cnt  int
}

func NewUFSet(n int) *UFSet {
	arr := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		size[i] = 1
	}

	return &UFSet{arr, size, n}
}

func (set *UFSet) Find(x int) int {
	if set.arr[x] != x {
		set.arr[x] = set.Find(set.arr[x])
	}
	return set.arr[x]
}

func (set *UFSet) Union(a, b int) bool {
	pa := set.Find(a)
	pb := set.Find(b)
	if pa == pb {
		return false
	}
	set.cnt--
	if set.size[pa] < set.size[pb] {
		pa, pb = pb, pa
	}
	set.size[pa] += set.size[pb]
	set.arr[pb] = pa
	return true
}
