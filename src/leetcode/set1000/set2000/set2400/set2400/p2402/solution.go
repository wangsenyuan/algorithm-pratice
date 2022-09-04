package p2402

import (
	"container/heap"
	"sort"
)

const INF = 1 << 31

func mostBooked(n int, meetings [][]int) int {
	m := len(meetings)

	mees := make([]Meeting, m)
	for i := 0; i < m; i++ {
		mees[i] = Meeting{meetings[i][0], meetings[i][1], i}
	}

	sort.Slice(mees, func(i, j int) bool {
		return mees[i].Less(mees[j])
	})

	arr := make([]int, 2*n)
	for i := n; i < 2*n; i++ {
		arr[i] = i - n
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = min(arr[i*2], arr[i*2+1])
	}

	get := func(l, r int) int {
		l += n
		r += n
		res := INF
		for l < r {
			if l&1 == 1 {
				res = min(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = min(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	set := func(p int, v int) {
		p += n
		arr[p] = v
		for p > 0 {
			arr[p>>1] = min(arr[p], arr[p^1])
			p >>= 1
		}
	}

	pq := make(PQ, 0, n)

	cnt := make([]int, n)

	for _, cur := range mees {

		for pq.Len() > 0 && pq[0].free_at <= cur.start {
			x := heap.Pop(&pq).(Room)
			set(x.id, x.id)
		}

		id := get(0, n)

		if id < n {
			// a free room
			cnt[id]++
			heap.Push(&pq, Room{id, cur.end})
			set(id, INF)
		} else {
			// no free room
			x := heap.Pop(&pq).(Room)
			// x will free at, and assign it to cur
			cnt[x.id]++
			next_free_at := x.free_at + (cur.end - cur.start)
			heap.Push(&pq, Room{x.id, next_free_at})
		}
	}

	var ans int

	for i := 0; i < n; i++ {
		if cnt[i] > cnt[ans] {
			ans = i
		}
	}
	return ans
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

type Meeting struct {
	start int
	end   int
	id    int
}

func (this Meeting) Less(that Meeting) bool {
	return this.start < that.start
}

type Room struct {
	id      int
	free_at int
}

type PQ []Room

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].free_at < pq[j].free_at || pq[i].free_at == pq[j].free_at && pq[i].id < pq[j].id
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Room))
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	ret := old[n-1]
	*pq = old[:n-1]
	return ret
}
