package main

import (
	"bufio"
	"container/heap"
	"fmt"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, k := readTwoNums(scanner)

		grid := make([][]int, n)

		for i := 0; i < n; i++ {
			grid[i] = readNNums(scanner, n)
		}

		res := solve(n, k, grid)

		fmt.Println(res)
	}
}

const INF = 1 << 30

func solve(n int, k int, grid [][]int) int {
	points := groupByHeight(n, k, grid)

	que := make([]int, n*n)
	var front, end int
	dist := make([][]int, n)

	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dist[i][j] = 0
				que[end] = i*n + j
				end++
			} else {
				dist[i][j] = INF
			}
		}
	}
	res := INF

	for front < end {
		var h int
		for front < end {
			cur := que[front]
			front++
			x, y := cur/n, cur%n
			h = grid[x][y]
			if h == k {
				res = min(res, dist[x][y])
			} else {
				ps := points[h+1]
				for _, p := range ps {
					u, v := p.x, p.y
					dist[u][v] = min(dist[u][v], dist[x][y]+abs(u-x)+abs(v-y))
				}
			}
		}
		if h < k {
			ps := points[h+1]
			for _, p := range ps {
				u, v := p.x, p.y
				que[end] = u*n + v
				end++
			}
		}
	}

	return res
}

func groupByHeight(n, k int, grid [][]int) [][]Point {
	points := make([][]Point, k+1)

	for i := 0; i <= k; i++ {
		points[i] = make([]Point, 0, 3)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			h := grid[i][j]
			points[h] = append(points[h], Point{i, j})
		}
	}
	return points
}

func solve1(n int, k int, grid [][]int) int {
	points := groupByHeight(n, k, grid)

	items := make([][]*Item, n)
	pq := make(PriorityQueue, n*n)

	for i := 0; i < n; i++ {
		items[i] = make([]*Item, n)
		for j := 0; j < n; j++ {
			item := new(Item)
			item.value = i*n + j
			if grid[i][j] == 1 {
				item.priority = 0
			} else {
				item.priority = INF
			}
			pq[i*n+j] = item
			items[i][j] = item
		}
	}

	res := INF

	heap.Init(&pq)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		x, y := cur.value/n, cur.value%n
		h := grid[x][y]

		if h == k {
			res = min(res, items[x][y].priority)
		} else {
			ps := points[h+1]

			for _, p := range ps {
				u, v := p.x, p.y
				item := items[u][v]

				if item.priority > items[x][y].priority+abs(x-u)+abs(y-v) {
					pq.update(item, items[x][y].priority+abs(x-u)+abs(y-v))
				}
			}
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Point struct {
	x int
	y int
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
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	// item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
