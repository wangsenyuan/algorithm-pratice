package p755

import "container/heap"

func pourWater(heights []int, V int, K int) []int {
	n := len(heights)
	if n == 1 {
		heights[0] += V
		return heights
	}
	cols := make([]*Col, n)
	for i := 0; i < n; i++ {
		dist := K - i
		if dist < 0 {
			dist = i - K
		}
		col := &Col{heights[i], i, dist}
		cols[i] = col
	}
	leftCols := make(Cols, 0, K)
	heap.Init(&leftCols)
	left := K - 1
	for left >= 0 && heights[left] <= heights[left+1] {
		heap.Push(&leftCols, cols[left])
		left--
	}
	rightCols := make(Cols, 0, n-K)
	heap.Init(&rightCols)
	right := K + 1
	for right < n && heights[right] <= heights[right-1] {
		heap.Push(&rightCols, cols[right])
		right++
	}

	for v := 0; v < V; v++ {
		var goLeft bool
		if leftCols.Len() > 0 {
			lowest := heap.Pop(&leftCols).(*Col)
			if lowest.height < cols[K].height {
				lowest.height++
				goLeft = true
			}
			//put it back
			heap.Push(&leftCols, lowest)
		}
		if !goLeft {
			var goRight bool
			if rightCols.Len() > 0 {
				lowest := heap.Pop(&rightCols).(*Col)
				if lowest.height < cols[K].height {
					lowest.height++
					goRight = true
				}
				heap.Push(&rightCols, lowest)
			}
			if !goRight {
				cols[K].height++
			}
		}
		for left >= 0 && cols[left].height <= cols[left+1].height {
			heap.Push(&leftCols, cols[left])
			left--
		}
		for right < n && cols[right].height <= cols[right-1].height {
			heap.Push(&rightCols, cols[right])
			right++
		}
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = cols[i].height
	}
	return res
}

type Col struct {
	height int
	at     int
	dist   int
}

type Cols []*Col

func (cols Cols) Len() int {
	return len(cols)
}

func (cols Cols) Less(i, j int) bool {
	return cols[i].height < cols[j].height || (cols[i].height == cols[j].height && cols[i].dist < cols[j].dist)
}

func (cols Cols) Swap(i, j int) {
	cols[i], cols[j] = cols[j], cols[i]
}

func (cols *Cols) Push(x interface{}) {
	col := x.(*Col)
	*cols = append(*cols, col)
}

func (cols *Cols) Pop() interface{} {
	old := *cols
	n := len(old)
	col := old[n-1]
	*cols = old[:n-1]
	return col
}
