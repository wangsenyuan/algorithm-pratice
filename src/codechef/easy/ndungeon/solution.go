package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var m, n int
	scanner.Scan()
	bs := scanner.Bytes()
	readInt(bs, readInt(bs, 0, &m)+1, &n)

	rooms := make([][]int, m)
	for i := 0; i < m; i++ {
		rooms[i] = make([]int, n)
		scanner.Scan()
		bs = scanner.Bytes()
		x := 0
		for j := 0; j < n; j++ {
			x = readInt(bs, x, &rooms[i][j]) + 1
		}
	}
	var a, b, T int
	scanner.Scan()
	bs = scanner.Bytes()
	readInt(bs, readInt(bs, readInt(bs, 0, &a)+1, &b)+1, &T)

	left, can := solve(m, n, rooms, a-1, b-1, T)
	if can {
		fmt.Println("YES")
		fmt.Println(left)
	} else {
		fmt.Println("NO")
	}
}

var dirs = []int{-1, 0, 1, 0, -1}

func solve(m int, n int, rooms [][]int, c int, d int, T int) (int, bool) {
	ds := make([][]int, m)
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		ds[i] = make([]int, n)
		visited[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			ds[i][j] = -1
		}
	}
	ds[0][0] = rooms[0][0]
	que := PriorityQueue(make([]*Item, 0, m*n))
	heap.Init(&que)
	heap.Push(&que, &Item{0, 0, rooms[0][0]})
	for que.Len() > 0 {
		v := heap.Pop(&que).(*Item)
		x, y, p := v.x, v.y, v.priority
		if visited[x][y] {
			continue
		}
		visited[x][y] = true
		for k := 0; k < 4; k++ {
			a, b := x+dirs[k], y+dirs[k+1]
			if a >= 0 && a < m && b >= 0 && b < n && (ds[a][b] == -1 || ds[a][b] > p+rooms[a][b]) {
				ds[a][b] = p + rooms[a][b]
				heap.Push(&que, &Item{a, b, ds[a][b]})
			}
		}
	}

	// fmt.Println("debug-----")
	// for _, row := range ds {
	// fmt.Println(row)
	// }
	// fmt.Println("debug done")
	took := ds[c][d]
	if took > T {
		return -1, false
	}
	return T - took, true
}

// An Item is something we manage in a priority queue.
type Item struct {
	x        int
	y        int
	priority int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
