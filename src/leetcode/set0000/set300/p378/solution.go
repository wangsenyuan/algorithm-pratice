package main

import (
	"container/heap"
	"fmt"
)

func main() {
	/*matrix := [][]int{
		[]int{1, 5, 9},
		[]int{10, 11, 13},
		[]int{12, 13, 15},
	}
	matrix := [][]int{
		[]int{1, 2},
		[]int{3, 3},
	} */
	matrix := [][]int{
		[]int{1, 3, 5},
		[]int{6, 7, 12},
		[]int{11, 14, 14},
	}

	r := kthSmallest(matrix, 3)
	fmt.Println(r)
}

func kthSmallest(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])

	count := func(num int) int {
		var cnt int

		var r, c = 0, n - 1

		for r < m && c >= 0 {
			if matrix[r][c] > num {
				c--
			} else {
				cnt += c + 1
				r++
			}
		}
		return cnt
	}

	left := matrix[0][0]
	right := matrix[m-1][n-1] + 1

	for left < right {
		mid := left + (right-left)/2
		c := count(mid)
		if c >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func kthSmallest1(matrix [][]int, k int) int {
	n := len(matrix)
	pq := make(CellPQ, 1)
	pq[0] = &Cell{0, 0, matrix[0][0]}
	heap.Init(&pq)

	for len(pq) > 0 {
		cell := heap.Pop(&pq).(*Cell)
		k--
		if k == 0 {
			return cell.val
		}
		if cell.i+1 < n {
			heap.Push(&pq, &Cell{cell.i + 1, cell.j, matrix[cell.i+1][cell.j]})
		}

		if cell.i == 0 && cell.j+1 < n {
			heap.Push(&pq, &Cell{cell.i, cell.j + 1, matrix[cell.i][cell.j+1]})
		}
	}
	panic("should not go here")
}

type Cell struct {
	i, j int
	val  int
}

type CellPQ []*Cell

func (pq CellPQ) Len() int { return len(pq) }

func (pq CellPQ) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].val < pq[j].val
}

func (pq CellPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *CellPQ) Push(x interface{}) {
	*pq = append(*pq, x.(*Cell))
}

func (pq *CellPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	cell := old[n-1]
	*pq = old[0 : n-1]
	return cell
}
