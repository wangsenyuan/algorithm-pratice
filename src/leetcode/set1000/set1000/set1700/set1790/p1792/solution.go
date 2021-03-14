package p1792

import "container/heap"

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	n := len(classes)
	// avg pass = (sum(ratio) / n)
	// n no change how to dispatch extra students
	// so to improve avg, need to increase ratio
	// which means to add x to some i, make (pass + x) / (total + x) - pass / total
	// (pass + x) * total - pass * (total + x) = (total - pass) x
	// (total + x) * total
	// total 越小，增加x的作用越大,
	// total - pass 越大，增加x的作用越大
	// 1 / 2, 1 / 3
	// 2/3, 2/4 =
	pq := make(PQ, 0, n)
	for i := 0; i < n; i++ {
		cur := new(Item)
		cur.first = classes[i][0]
		cur.second = classes[i][1]
		heap.Push(&pq, cur)
	}

	for extraStudents > 0 {
		best := pq[0]
		best.first++
		best.second++
		heap.Fix(&pq, 0)
		extraStudents--
	}

	var res float64

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		tmp := float64(cur.first) / float64(cur.second)
		res += tmp
	}
	return res / float64(n)
}

type Item struct {
	first, second int
	index         int
}

func (this Item) Less(that Item) bool {
	a := this.second - this.first
	b := this.second * (this.second + 1)
	c := that.second - that.first
	d := that.second * (that.second + 1)
	return a*d > b*c
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

func (pq *PQ) Push(item interface{}) {
	x := item.(*Item)
	x.index = len(*pq)
	*pq = append(*pq, x)
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
