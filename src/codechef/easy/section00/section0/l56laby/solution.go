package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n, m := readTwoNums(scanner)
		grid := make([][]int, n)
		for i := 0; i < n; i++ {
			grid[i] = readNNums(scanner, m)
		}
		res := solve(n, m, grid)
		var buf bytes.Buffer
		for i := 0; i < n; i++ {
			buf.Write(res[i])
			buf.WriteString("\n")
		}
		fmt.Print(buf.String())
		t--
	}
}

func solve(n, m int, grid [][]int) [][]byte {
	res := make([][]byte, n)

	cells := make([]*Cell, 0, n*m)
	for i := 0; i < n; i++ {
		res[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			if grid[i][j] > 0 {
				cells = append(cells, &Cell{i, j, grid[i][j]})
			}
			if grid[i][j] < 0 {
				res[i][j] = 'B'
			} else {
				res[i][j] = 'N'
			}
		}
	}
	pq := PriorityQueue(cells)
	heap.Init(&pq)

	dd := []int{-1, 0, 1, 0, -1}

	for pq.Len() > 0 {
		cell := heap.Pop(&pq).(*Cell)
		i, j, x := cell.x, cell.y, cell.v
		res[i][j] = 'Y'
		x--
		if x < 0 {
			continue
		}
		for k := 0; k < 4; k++ {
			a, b := i+dd[k], j+dd[k+1]
			if a >= 0 && a < n && b >= 0 && b < m && res[a][b] == 'N' {
				heap.Push(&pq, &Cell{a, b, x})
			}
		}
	}

	return res
}

type Cell struct {
	x int
	y int
	v int
}

type PriorityQueue []*Cell

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].v > pq[j].v
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Cell)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
