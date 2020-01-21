package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
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
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		edges := make([][]int, n-1)

		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(scanner, 3)
		}
		A := readNNums(scanner, n)
		res := solve(n, edges, A)
		fmt.Println(res)
	}
}

const INF = 1000000009

func solve(n int, edges [][]int, A []int) int {
	// sort.Ints(A)

	B := make([]int, n-1)

	conn := make([][]int, n)
	degree := make([]int, n)
	connect := func(u, v int) {
		if conn[u] == nil {
			conn[u] = make([]int, 0, 10)
		}
		conn[u] = append(conn[u], v)
		degree[u]++
	}

	for i, edge := range edges {
		B[i] = edge[2]
		u, v := edge[0], edge[1]
		u--
		v--
		connect(u, v)
		connect(v, u)
	}

	sort.Ints(B)

	que := make([]int, n)
	var end int

	for i := 0; i < n; i++ {
		// a leaf
		if degree[i] == 1 {
			que[end] = i
			end++
		}
	}

	pq := make(IntHeap, n)

	for i := 0; i < n; i++ {
		pq[i] = A[i]
	}

	heap.Init(&pq)

	var front int
	var ii int

	var ans int

	for front < end && ii < len(B) {
		u := que[front]
		front++
		// try assign a num at edge [u, p], where p is u's parent
		num := B[ii]
		ii++

		// then A[u] >= num

		for {
			x := heap.Pop(&pq).(int)
			if x >= num {
				break
			}

			ans++
			// increase x to INF and put it back to heap
			heap.Push(&pq, INF)
		}

		for _, v := range conn[u] {
			degree[v]--
			if degree[v] == 1 {
				que[end] = v
				end++
			}
		}
	}

	X := heap.Pop(&pq).(int)

	if X < B[n-2] {
		ans++
	}

	return ans

}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
