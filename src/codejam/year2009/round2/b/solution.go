package main

import (
	"container/heap"
	"math"
	"bufio"
	"os"
	"fmt"
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
		firstLine := readNNums(scanner, 3)
		R, C, F := firstLine[0], firstLine[1], firstLine[2]
		G := make([]string, R)
		for i := 0; i < R; i++ {
			scanner.Scan()
			G[i] = scanner.Text()
		}
		res := solve(R, C, F, G)
		if res < 0 {
			fmt.Printf("Case #%d: No\n", x)
		} else {
			fmt.Printf("Case #%d: Yes %d\n", x, res)
		}
	}
}

func solve(R, C, F int, G []string) int {
	items := make([][][][]*Item, R)
	que := make(PriorityQueue, 0, R*C*C*C)
	for i := 0; i < R; i++ {
		items[i] = make([][][]*Item, C)
		for j := 0; j < C; j++ {
			items[i][j] = make([][]*Item, C)
			for l := 0; l < C; l++ {
				items[i][j][l] = make([]*Item, C)
				for r := 0; r < C; r++ {
					item := new(Item)
					item.x = i
					item.y = j
					item.l = l
					item.r = r
					item.cost = math.MaxInt32
					items[i][j][l][r] = item
					heap.Push(&que, item)
				}
			}
		}
	}

	game := make([][]byte, R)
	for i := 0; i < R; i++ {
		game[i] = []byte(G[i])
	}

	bak := make([]byte, C)

	fall := func(x, y int) int {
		for x < R-1 && game[x+1][y] == '.' {
			x++
		}
		return x
	}

	que.update(items[0][0][0][0], 0)
	for que.Len() > 0 {
		cur := heap.Pop(&que).(*Item)
		x, y, l, r, c := cur.x, cur.y, cur.l, cur.r, cur.cost
		if c == math.MaxInt32 {
			return -1
		}

		if x == R-1 {
			return c
		}

		for i := 0; i < C; i++ {
			bak[i] = '.'
			if i >= y-l && i <= y+r {
				//current row, it should be empty
				game[x][i], bak[i] = bak[i], game[x][i]
			}
		}

		for dy := -1; dy <= 1; dy += 2 {
			yy := y + dy
			if yy < 0 || yy >= C || game[x][yy] == '#' {
				continue
			}
			tx := fall(x, yy)
			if tx-x > F {
				continue
			}
			tl, tr := 0, 0
			if tx == x {
				// reach another empty cell in the same row
				tl, tr = max(0, l+dy), max(0, r-dy)
			}
			// else fall through (x, yy) to (tx, yy)

			if items[tx][yy][tl][tr].cost > c {
				que.update(items[tx][yy][tl][tr], c)
			}
		}

		for dy := -1; x < R-1 && dy <= 1; dy += 2 {
			for d := 1; true; d++ {
				//dig next row
				yy := y + dy*d
				if yy < 0 || yy >= C || game[x][yy] == '#' {
					break
				}
				if game[x+1][yy] == '.' {
					//will fall
					break
				}
				//tc := c + d
				//after dug tc bricks, can reach (x + 1, y + dy)
				tx := fall(x+1, y+dy)

				if tx-x > F {
					break
				}

				var tl, tr int
				if tx == x+1 {
					if dy < 0 {
						tl, tr = d-1, 0
					} else {
						tl, tr = 0, d-1
					}
				}
				if items[tx][y+dy][tl][tr].cost > c+d {
					que.update(items[tx][y+dy][tl][tr], c+d)
				}
			}
		}
		for i := 0; i < C; i++ {
			if i >= y-l && i <= y+r {
				//swap back
				game[x][i], bak[i] = bak[i], game[x][i]
			}
		}
	}
	return -1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// An Item is something we manage in a priority queue.
type Item struct {
	x, y, l, r int // The value of the item; arbitrary.
	cost       int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].cost < pq[j].cost
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

func (pq *PriorityQueue) update(item *Item, cost int) {
	item.cost = cost
	heap.Fix(pq, item.index)
}
