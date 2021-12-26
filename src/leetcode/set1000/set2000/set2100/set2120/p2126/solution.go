package p2126

import "sort"

func recoverArray(nums []int) []int {
	N := len(nums)
	n := N / 2
	P := make([]Pair, N)

	for i := 0; i < N; i++ {
		P[i] = Pair{nums[i], i}
	}

	sort.Slice(P, func(i, j int) bool {
		return P[i].first < P[j].first
	})

	update := func(pos map[int][]int, i int, p Pair) {
		f := p.first
		if _, found := pos[f]; !found {
			pos[f] = make([]int, 0, 2)
		}
		pos[f] = append(pos[f], i)
	}

	arr := make([]int, n)

	used := make([]bool, N)

	check := func(k int) bool {
		// a == 0
		pos := make(map[int][]int)
		for i, p := range P {
			update(pos, i, p)
			used[i] = false
		}
		var cnt int
		for i := 0; i < N && cnt < n; i++ {
			if used[i] {
				continue
			}
			v := P[i].first
			arr[cnt] = v + k
			w := v + 2*k
			if ws, found := pos[w]; !found || len(ws) == 0 {
				return false
			} else {
				used[ws[0]] = true
				pos[w] = ws[1:]
				cnt++
			}
		}

		return cnt == n
	}

	for i := 1; i+n <= N; i++ {
		k := P[i].first - P[0].first
		if k == 0 || k%2 == 1 {
			continue
		}
		k /= 2
		if check(k) {
			return arr
		}
	}

	return nil
}

type Pair struct {
	first, second int
}
