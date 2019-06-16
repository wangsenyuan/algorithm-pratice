package p764

import "sort"

func orderOfLargestPlusSign(N int, mines [][]int) int {
	rows := make([]*S, N)
	cols := make([]*S, N)

	for i := 0; i < N; i++ {
		rows[i] = &S{start: 0, end: N - 1, idx: i}
		cols[i] = &S{start: 0, end: N - 1, idx: i}
	}

	for _, mine := range mines {
		r, c := mine[0], mine[1]
		rows[r] = insertHole(rows[r], c, r)
		cols[c] = insertHole(cols[c], r, c)
	}

	rs := make([]*S, 0, 10)
	for i := 0; i < N; i++ {
		tmp := rows[i]
		for tmp != nil {
			rs = append(rs, tmp)
			tmp = tmp.next
		}
	}
	cs := make([]*S, 0, 10)
	for i := 0; i < N; i++ {
		tmp := cols[i]
		for tmp != nil {
			cs = append(cs, tmp)
			tmp = tmp.next
		}
	}

	sort.Sort(SS(rs))
	sort.Sort(SS(cs))

	xs := make(map[int]map[int]int)
	for j := 0; j < len(cs); j++ {
		c := cs[j]
		if _, found := xs[c.idx]; !found {
			xs[c.idx] = make(map[int]int)
		}
		xs[c.idx][c.start] = c.end
	}

	check := func(ord int) bool {
		l := 2*(ord-1) + 1
		if l > N {
			return false
		}
		i := sort.Search(len(rs), func(i int) bool {
			return rs[i].end-rs[i].start+1 >= l
		})
		if i == len(rs) {
			return false
		}

		j := sort.Search(len(cs), func(j int) bool {
			return cs[j].end-cs[j].start+1 >= l
		})
		if j == len(cs) {
			return false
		}

		for i < len(rs) {
			r := rs[i]
			rn := r.idx
			mid := r.start + ord - 1
			for mid+ord-1 <= r.end {
				if x, found := xs[mid]; found {
					for start, end := range x {
						if start+ord-1 <= rn && rn+ord-1 <= end {
							return true
						}
					}
				}
				mid++
			}
			i++
		}
		return false
	}

	l, r := 0, N/2+2
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l - 1
}

type S struct {
	prev  *S
	next  *S
	start int
	end   int
	idx   int
}

func insertHole(head *S, end int, idx int) *S {
	tmp := head
	for tmp.end < end {
		tmp = tmp.next
	}

	if tmp.end == end {
		tmp.end--
	} else {
		cur := &S{start: end + 1, end: tmp.end, idx: idx}
		tmp.end = end - 1
		next := tmp.next
		if next != nil {
			cur.next = next
			next.prev = cur
		}
		tmp.next = cur
		cur.prev = tmp
	}

	return head
}

type SS []*S

func (this SS) Len() int {
	return len(this)
}

func (this SS) Less(i, j int) bool {
	return this[i].end-this[i].start < this[j].end-this[j].start
}

func (this SS) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
