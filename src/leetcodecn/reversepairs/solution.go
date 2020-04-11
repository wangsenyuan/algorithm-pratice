package reversepairs

import "sort"

func reversePairs(nums []int) int {
	n := len(nums)
	ps := make([]Pair, n)

	for i := 0; i < n; i++ {
		ps[i] = Pair{nums[i], i}
	}

	sort.Sort(Pairs(ps))

	arr := make([]int, n+1)

	set := func(pos int) {
		pos++
		for pos <= n {
			arr[pos]++
			pos += pos & -pos
		}
	}

	get := func(pos int) int {
		var res int
		pos++
		for pos > 0 {
			res += arr[pos]
			pos -= pos & -pos
		}
		return res
	}
	var res int
	for i := 0; i < n; i++ {
		cur := ps[i]

		res += i - get(cur.second)

		set(cur.second)
	}

	return res
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first || ps[i].first == ps[j].first && ps[i].second < ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func reversePairs1(nums []int) int {
	n := len(nums)

	if n <= 1 {
		return 0
	}

	back := make([]int, n)

	var res int

	merge := func(a, b, c int) {
		i := a
		j := b
		k := a
		for i < b && j < c {
			if nums[i] <= nums[j] {
				back[k] = nums[i]
				k++
				i++
			} else {
				// reverse
				res += b - i
				back[k] = nums[j]
				k++
				j++
			}
		}

		for i < b {
			back[k] = nums[i]
			k++
			i++
		}

		for j < c {
			back[k] = nums[j]
			k++
			j++
		}

		for i := a; i < c; i++ {
			nums[i] = back[i]
		}
	}

	var sort func(i, j int)

	sort = func(i, j int) {
		if i == j {
			return
		}

		mid := (i + j) / 2
		sort(i, mid)
		sort(mid+1, j)

		merge(i, mid+1, j+1)
	}

	sort(0, n-1)

	return res
}
