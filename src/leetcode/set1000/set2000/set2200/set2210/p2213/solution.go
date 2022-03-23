package p2213

import "container/heap"

func allOne(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}
	return res
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	n := len(s)
	if n == 1 {
		return allOne(len(queryIndices))
	}

	X := make([]int, 2*n)
	Y := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		X[i] = 0
		Y[i] = n
	}
	set := func(p int) {
		p += n
		X[p] = p - n
		Y[p] = p - n

		for p > 0 {
			X[p>>1] = max(X[p], X[p^1])
			Y[p>>1] = min(Y[p], Y[p^1])
			p >>= 1
		}
	}

	unset := func(p int) {
		p += n
		X[p] = 0
		Y[p] = n

		for p > 0 {
			X[p>>1] = max(X[p], X[p^1])
			Y[p>>1] = min(Y[p], Y[p^1])
			p >>= 1
		}
	}

	get := func(l, r int) (x int, y int) {
		x = 0
		y = n
		l += n
		r += n
		for l < r {
			if l&1 == 1 {
				x = max(x, X[l])
				y = min(y, Y[l])
				l++
			}
			if r&1 == 1 {
				r--
				x = max(x, X[r])
				y = min(y, Y[r])
			}
			l >>= 1
			r >>= 1
		}
		return
	}

	buf := []byte(s)

	pq := make(PQ, 0, n)
	items := make([]*Item, n)

	res := make([]int, len(queryIndices))

	erase := func(p int) {
		item := items[p]
		pq.update(item, n+1)
		heap.Pop(&pq)
		items[p] = nil
		unset(p)
	}

	upset := func(p int, l int) {
		if items[p] != nil {
			pq.update(items[p], l)
		} else {
			item := new(Item)
			item.pos = p
			item.length = l
			heap.Push(&pq, item)
			items[p] = item
		}
		set(p)
	}

	for i := 0; i < n; {
		j := i
		for i < n && buf[j] == buf[i] {
			i++
		}
		upset(j, i-j)
	}

	getLeft := func(p int) int {
		x, _ := get(0, p)
		return x
	}
	getRight := func(p int) int {
		_, y := get(p, n)
		return y
	}

	for i, j := range queryIndices {
		c := queryCharacters[i]
		if c != buf[j] {
			if j > 0 && buf[j] != buf[j-1] {
				erase(j)
			}

			if j+1 < n && buf[j] != buf[j+1] {
				erase(j + 1)
			}

			buf[j] = c
			l, r := j, j
			if j > 0 {
				x := getLeft(j)
				if buf[j] != buf[j-1] {
					// now differ
					upset(x, j-x)
				} else {
					// now same
					l = x
				}
			}
			if j+1 < n {
				if buf[j] != buf[j+1] {
					y := getRight(j+1) - 1
					upset(j+1, y-(j+1)+1)
				} else {
					r = getRight(j+1) - 1
				}
			}

			upset(l, r-l+1)
		}

		res[i] = pq[0].length
		buf[j] = c
	}

	return res
}

type Item struct {
	pos    int
	length int
	index  int
}

func (this Item) Less(that Item) bool {
	return this.length > that.length
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].Less(*pq[j])
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	ret := old[n-1]
	ret.index = -1
	*pq = old[:n-1]
	return ret
}

func (pq *PQ) update(item *Item, length int) {
	item.length = length
	heap.Fix(pq, item.index)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
