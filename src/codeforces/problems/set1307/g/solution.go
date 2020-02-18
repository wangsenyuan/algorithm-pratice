package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
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

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, m := readTwoNums(scanner)

	edges := make([][]int, m)

	for i := 0; i < m; i++ {
		edges[i] = readNNums(scanner, 3)
	}

	solver := NewSolver(n, m, edges)

	q := readNum(scanner)

	for q > 0 {
		q--

		x := readNum(scanner)

		r := solver.Query(x)

		fmt.Printf("%.7f\n", r)
	}

}

type Solver struct {
	paths []int
}

func (solver Solver) Query(x int) float64 {
	X := float64(x)

	paths := solver.paths

	check := func(y float64) bool {
		used := y

		h := used + float64(paths[0])

		for i := 1; i < len(paths); i++ {
			cur := float64(paths[i])
			if cur >= h {
				break
			}
			used += h - cur
		}

		return used <= X
	}

	r := float64(paths[0])

	var left float64
	var right = X + 0.1

	for left < right {
		diff := right - left
		if math.Abs(diff) < 1e-8 {
			return left + r
		}
		mid := (left + right) / 2
		if !check(mid) {
			right = mid
		} else {
			left = mid
		}
	}

	return r
}

const INF = 1 << 30

func NewSolver(n, m int, edges [][]int) Solver {
	used := make([]bool, m)

	outs := make([][]int, n)

	for i := 0; i < n; i++ {
		outs[i] = make([]int, 0, 3)
	}

	for i, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		outs[u] = append(outs[u], i)
		outs[v] = append(outs[v], i)
	}

	items := make([]*Item, n)
	inEdge := make([]int, n)

	bfs := func() int {
		pq := make(PriorityQueue, n)

		for i := 0; i < n; i++ {
			item := NewItem(i, INF, i)
			items[i] = item
			pq[i] = item
			inEdge[i] = -1
		}

		heap.Init(&pq)

		pq.update(items[0], 0)

		for pq.Len() > 0 {
			cur := heap.Pop(&pq).(*Item)
			u := cur.value
			d := cur.priority

			for _, j := range outs[u] {
				if used[j] {
					continue
				}

				edge := edges[j]
				v, w := edge[1]-1, edge[2]

				if d+w < items[v].priority && items[v].index >= 0 {
					inEdge[v] = j
					pq.update(items[v], d+w)
				}
			}
		}

		if inEdge[n-1] < 0 {
			return -1
		}

		var res int

		x := n - 1

		for x > 0 {
			j := inEdge[x]

			used[j] = true

			u, v, w := edges[j][0]-1, edges[j][1]-1, edges[j][2]

			res += w

			y := u
			if u == x {
				y = v
			}
			x = y
		}

		return res
	}

	var paths []int

	for {
		res := bfs()
		if res < 0 {
			break
		}
		paths = append(paths, res)
	}

	return Solver{paths}
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func NewItem(value, prioirity, index int) *Item {
	item := new(Item)
	item.value = value
	item.priority = prioirity
	item.index = index
	return item
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
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
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
