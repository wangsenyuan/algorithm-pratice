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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for x := 1; x <= tc; x++ {
		n, m := readTwoNums(scanner)
		G := make([]string, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			G[i] = scanner.Text()
		}
		queries := readNNums(scanner, m)
		res := solve(n, G, queries)
		fmt.Printf("Case #%d:\n", x)
		for i := 0; i < m; i++ {
			fmt.Println(res[i])
		}
	}
}

var dd = []int{-1, 0, 1, 0, -1}

const MAX_V = 4001
const HALF = 2000

func solve(n int, G []string, queries []int) []string {

	valid := func(x, y int) bool {
		return x >= 0 && x < n && y >= 0 && y < n
	}

	A := make([][][]*Item, n)
	for i := 0; i < n; i++ {
		A[i] = make([][]*Item, n)
		for j := 0; j < n; j++ {
			A[i][j] = make([]*Item, MAX_V)
			for v := -HALF; v <= HALF; v++ {
				item := new(Item)
				item.i = i
				item.j = j
				item.v = v
				// mock expr
				item.expr = ""
				A[i][j][v+HALF] = item
			}

		}
	}

	que := make(PriorityQueue, 0, n*n*MAX_V)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if G[i][j] >= '0' && G[i][j] <= '9' {
				x := int(G[i][j] - '0')
				item := A[i][j][x+HALF]
				item.expr = string(G[i][j])
				heap.Push(&que, item)
			}
		}
	}

	for que.Len() > 0 {
		head := heap.Pop(&que).(*Item)
		r, c, val := head.i, head.j, head.v
		for k := 0; k < 4; k++ {
			u, v := r+dd[k], c+dd[k+1]
			if valid(u, v) {
				op := G[u][v]

				for kk := 0; kk < 4; kk++ {
					a, b := u+dd[kk], v+dd[kk+1]
					if valid(a, b) {
						x := int(G[a][b] - '0')
						var y int
						if op == '+' {
							y = val + x
						} else {
							y = val - x
						}

						idx := y + HALF
						if idx < 0 || idx >= MAX_V {
							continue
						}

						tmp := head.expr + string(op) + string(G[a][b])
						xx := A[a][b][idx]
						if len(xx.expr) == 0 {
							xx.expr = tmp
							heap.Push(&que, xx)
						} else if len(xx.expr) > len(tmp) || (len(xx.expr) == len(tmp) && tmp < xx.expr) {
							que.update(xx, tmp)
						}
					}
				}
			}
		}
	}

	query := func(num int) string {
		var best string
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if G[i][j] >= '0' && G[i][j] <= '9' {
					tmp := A[i][j][num+HALF].expr
					if len(tmp) > 0 &&
						(len(best) == 0 ||
							len(tmp) < len(best) ||
							(len(tmp) == len(best) && tmp < best)) {
						best = tmp
					}
				}
			}
		}
		return best
	}

	res := make([]string, len(queries))

	for k := 0; k < len(queries); k++ {
		num := queries[k]
		res[k] = query(num)
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// An Item is something we manage in a priority queue.
type Item struct {
	i, j, v int
	expr    string // The value of the item; arbitrary.
	index   int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return len(pq[i].expr) < len(pq[j].expr) ||
		(len(pq[i].expr) == len(pq[j].expr) && pq[i].expr < pq[j].expr)
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
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, expr string) {
	item.expr = expr
	heap.Fix(pq, item.index)
}
