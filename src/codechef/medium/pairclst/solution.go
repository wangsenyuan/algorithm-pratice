package main

import (
	"container/heap"
	"bufio"
	"os"
	"fmt"
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
	firstLine := readNNums(scanner, 3)
	N, M, K := firstLine[0], firstLine[1], firstLine[2]
	special := readNNums(scanner, K)
	edges := make([][]int, M)
	for i := 0; i < M; i++ {
		edges[i] = readNNums(scanner, 3)
	}
	fmt.Println(solve(N, M, K, special, edges))
}

func solve(N, M, K int, special []int, edges [][]int) int {
	graph := make([]map[int]int, N)

	for i := 0; i < N; i++ {
		graph[i] = make(map[int]int)
	}

	for _, edge := range edges {
		u, v, w := edge[0]-1, edge[1]-1, edge[2]
		graph[u][v] = w
		graph[v][u] = w
	}

	dists := make([]int, N)

	for i := 0; i < N; i++ {
		dists[i] = -1
	}

	pq := make(PriorityQueue, 0, N)
	parent := make([]int, N)

	for i := 0; i < K; i++ {
		parent[special[i]-1] = special[i] - 1
		dists[special[i]-1] = 0
		heap.Push(&pq, &Item{value: special[i] - 1, priority: 0})
	}

	seen := make([]bool, N)

	for pq.Len() > 0 {
		head := heap.Pop(&pq).(*Item)
		u, d := head.value, head.priority
		if seen[u] {
			continue
		}
		seen[u] = true

		for v, w := range graph[u] {
			if dists[v] < 0 || dists[v] > d+w {
				dists[v] = d + w
				parent[v] = u
				heap.Push(&pq, &Item{value: v, priority: dists[v]})
			}
		}
	}

	var boss func(u int) int
	boss = func(u int) int {
		if parent[u] == u {
			return u
		}
		parent[u] = boss(parent[u])
		return parent[u]
	}

	best := -1
	for u := 0; u < N; u++ {
		for v, w := range graph[u] {
			if boss(u) == boss(v) {
				continue
			}
			tmp := dists[u] + dists[v] + w
			if best < 0 || best > tmp {
				best = tmp
			}
		}
	}
	return best
}

func solve1(N, M, K int, special []int, edges [][]int) int {
	graph := make([]map[int]int, N)

	for i := 0; i < N; i++ {
		graph[i] = make(map[int]int)
	}

	specialFlag := make([]bool, N)

	for i := 0; i < K; i++ {
		specialFlag[special[i]-1] = true
	}

	for _, edge := range edges {
		u, v, w := edge[0]-1, edge[1]-1, edge[2]
		graph[u][v] = w
		graph[v][u] = w
	}

	INF := 1 << 25

	dists := make([]int, N)

	best := INF

	pq := make(PriorityQueue, 0, N)

	seen := make([]bool, N)

	bfs := func(x int) {
		for i := 0; i < N; i++ {
			dists[i] = INF
			seen[i] = false
		}
		dists[x] = 0
		heap.Push(&pq, &Item{value: x, priority: 0})

		for pq.Len() > 0 {
			hd := heap.Pop(&pq).(*Item)
			u, d := hd.value, hd.priority
			if seen[u] {
				continue
			}
			seen[u] = true

			if u != x && specialFlag[u] && d < best {
				best = d
			}

			for v, w := range graph[u] {
				if dists[v] > d+w {
					dists[v] = d + w
					heap.Push(&pq, &Item{value: v, priority: dists[v]})
				}
			}
		}
	}

	for i := 0; i < K; i++ {
		bfs(special[i] - 1)
	}

	return best
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
