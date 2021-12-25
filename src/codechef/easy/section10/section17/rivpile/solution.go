package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		line := readNNums(scanner, 3)
		N, M, W := line[0], line[1], line[2]
		piles := make([][]int, N)
		for i := 0; i < N; i++ {
			piles[i] = readNNums(scanner, 2)
		}
		disks := make([][]int, M)
		for i := 0; i < M; i++ {
			disks[i] = readNNums(scanner, 2)
		}
		res := solve(N, M, W, piles, disks)

		if res >= 0 {
			fmt.Println(res)
		} else {
			fmt.Println("impossible")
		}
	}
}

const INF = math.MaxInt64 - 1

func solve(N, M, W int, piles [][]int, disks [][]int) int64 {
	sort.Sort(PairSlice(disks))

	stacks := make([][]int, M)
	var p int

	for i := 0; i < M; i++ {
		for p > 0 && stacks[p-1][1] >= disks[i][1] {
			p--
		}
		stacks[p] = disks[i]
		p++
	}

	disks = stacks
	M = p

	edge := make([][][]int, N)
	for i := 0; i < N; i++ {
		edge[i] = make([][]int, N)
		for j := 0; j < N; j++ {
			edge[i][j] = make([]int, M)
			for k := 0; k < M; k++ {
				edge[i][j][k] = -1
			}
		}
	}

	for i := 0; i < N; i++ {
		x, y := piles[i][0], piles[i][1]
		for j := 0; j < N; j++ {
			u, v := piles[j][0], piles[j][1]
			dd := distance(x, y, u, v)

			b := M - 1
			for a := 0; a < M; a++ {
				if dd > float64(disks[a][0])+float64(disks[b][0]) {
					edge[i][j][a] = -1
					continue
				}
				for b > 0 && dd <= float64(disks[a][0])+float64(disks[b-1][0]) {
					b--
				}
				edge[i][j][a] = b
			}
		}
	}

	items := make([]*Item, N*M)
	pq := make(PriorityQueue, N*M)
	for i := 0; i < N; i++ {
		y := piles[i][1]
		for j := 0; j < M; j++ {
			r, c := disks[j][0], int64(disks[j][1])

			item := new(Item)
			item.value = i*M + j
			item.priority = INF

			if y <= r {
				item.priority = c
			}

			item.index = i*M + j
			items[i*M+j] = item
			pq[i*M+j] = item
		}
	}

	heap.Init(&pq)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		if cur.priority == INF {
			break
		}

		i, j := cur.value/M, cur.value%M

		y := piles[i][1]
		r := disks[j][0]

		if W-y <= r {
			return cur.priority
		}

		if j+1 < M {
			next := cur.priority + int64(disks[j+1][1]) - int64(disks[j][1])
			if items[i*M+j+1].index >= 0 && items[i*M+j+1].priority > next {
				pq.update(items[i*M+j+1], next)
			}
		}

		for k := 0; k < N; k++ {
			if k != i && edge[i][k][j] >= 0 {
				l := edge[i][k][j]
				next := cur.priority + int64(disks[l][1])
				if items[k*M+l].index >= 0 && items[k*M+l].priority > next {
					pq.update(items[k*M+l], next)
				}
			}
		}

	}

	return -1
}

type PairSlice [][]int

func (this PairSlice) Len() int {
	return len(this)
}

func (this PairSlice) Less(i, j int) bool {
	return this[i][0] < this[j][0]
}

func (this PairSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int   // The value of the item; arbitrary.
	priority int64 // The priority of the item in the queue.
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
func (pq *PriorityQueue) update(item *Item, priority int64) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

func solve1(N, M, W int, piles [][]int, disks [][]int) int64 {
	// N <= 250
	// M <= 250
	// dp[i][j] = the min cost at piles i, with disk j on it
	dp := make([][]int64, N+2)
	for i := 0; i <= N+1; i++ {
		dp[i] = make([]int64, M+1)
		for j := 0; j <= M; j++ {
			dp[i][j] = -1
		}
	}

	for i := 0; i < N; i++ {
		pile := piles[i]
		y := pile[1]

		for j := 0; j < M; j++ {
			disk := disks[j]
			r, c := disk[0], disk[1]
			if r >= y {
				dp[i][j] = int64(c)
			}
		}
	}

	vis := make([][]bool, N)
	for i := 0; i < N; i++ {
		vis[i] = make([]bool, M)
	}

	for {
		ui, uj := -1, -1

		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				if dp[i][j] < 0 || vis[i][j] {
					continue
				}
				if ui < 0 || dp[i][j] < dp[ui][uj] {
					ui = i
					uj = j
				}
			}
		}
		if ui < 0 {
			break
		}

		// current is ui, uj
		ux, uy := piles[ui][0], piles[ui][1]
		ur := disks[uj][0]
		if W-uy <= ur {
			// find solution
			return dp[ui][uj]
		}
		vis[ui][uj] = true
		// update dp

		for i := 0; i < N; i++ {
			pile := piles[i]
			x, y := pile[0], pile[1]
			dd := distance(ux, uy, x, y)
			for j := 0; j < M; j++ {
				if vis[i][j] {
					continue
				}
				disk := disks[j]
				r, c := disk[0], int64(disk[1])

				if dd <= float64(ur)+float64(r) {
					// can reach
					if dp[i][j] < 0 || dp[i][j] > dp[ui][uj]+c {
						dp[i][j] = dp[ui][uj] + c
					}
				}
			}
		}
	}

	return -1
}

func distance(a, b, c, d int) float64 {
	dx := float64(c) - float64(a)
	dy := float64(d) - float64(b)
	return math.Sqrt(dx*dx + dy*dy)
}
