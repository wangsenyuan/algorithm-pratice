package main

import (
	"bufio"
	"bytes"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for x := 1; x <= tc; x++ {
		R, C := readTwoNums(scanner)
		puzzle := make([]string, R)
		for i := 0; i < R; i++ {
			scanner.Scan()
			puzzle[i] = scanner.Text()
		}
		res := solve(R, C, puzzle)
		fmt.Printf("Case #%d: %d\n", x, res)
	}
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(R, C int, puzzle []string) int {
	var cnt int
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if puzzle[i][j] == 'o' || puzzle[i][j] == 'w' {
				cnt++
			}
		}
	}

	if cnt == 1 {
		return solveOne(R, C, puzzle)
	}

	begin, end := make(State, cnt), make(State, cnt)
	var bi, ei int
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if puzzle[i][j] == 'x' || puzzle[i][j] == 'w' {
				end[ei] = []int{i, j}
				ei++
			}
			if puzzle[i][j] == 'o' || puzzle[i][j] == 'w' {
				begin[bi] = []int{i, j}
				bi++
			}
		}
	}

	isEmpty := func(x, y int, cur State) bool {
		if x < 0 || x >= R {
			return false
		}
		if y < 0 || y >= C {
			return false
		}
		if puzzle[x][y] == '#' {
			// a wall
			return false
		}

		for i := 0; i < cnt; i++ {
			a, b := cur[i][0], cur[i][1]
			if a == x && b == y {
				return false
			}
		}
		return true
	}

	pushBox := func(cur State) []State {
		res := make([]State, 0, 10)
		for i := 0; i < len(cur); i++ {
			x, y := cur[i][0], cur[i][1]
			for k := 0; k < 4; k++ {
				u, v := x+dd[k], y+dd[k+1]
				if isEmpty(u, v, cur) {
					a, b := x-dd[k], y-dd[k+1]
					if isEmpty(a, b, cur) {
						next := make(State, cnt)
						copy(next, cur)
						next[i] = []int{a, b}
						sort.Sort(next)
						res = append(res, next)
					}
				}
			}
		}
		return res
	}

	sq := make([]int, cnt)
	checked := make([]bool, cnt)
	isDangerous := func(cur State) bool {
		for i := 0; i < cnt; i++ {
			checked[i] = false
		}
		front, end := 0, 0
		sq[end] = 0

		checked[0] = true

		end++
		for front < end {
			idx := sq[front]
			front++
			x, y := cur[idx][0], cur[idx][1]

			for i := 0; i < cnt; i++ {
				if checked[i] {
					continue
				}
				a, b := cur[i][0], cur[i][1]
				if a == x && abs(y-b) == 1 || b == y && abs(x-a) == 1 {
					//connect
					checked[i] = true
					sq[end] = i
					end++
				}
			}
		}

		return end < cnt
	}

	sort.Sort(begin)
	sort.Sort(end)

	endKey := end.String()
	vis := make(map[string]bool)
	que := make(PriorityQueue, 0, 1000)
	heap.Push(&que, &Item{begin, 0})
	for que.Len() > 0 {
		cur := heap.Pop(&que).(*Item)
		curKey := cur.val.String()
		if curKey == endKey {
			return cur.dist
		}
		if vis[curKey] {
			continue
		}
		vis[curKey] = true

		nexts := pushBox(cur.val)

		for _, next := range nexts {
			if vis[next.String()] {
				continue
			}

			if isDangerous(next) {
				nn := pushBox(next)

				for _, n := range nn {
					if isDangerous(n) || vis[n.String()] {
						continue
					}
					heap.Push(&que, &Item{n, cur.dist + 2})
				}
			} else {
				heap.Push(&que, &Item{next, cur.dist + 1})
			}
		}
	}
	return -1
}

func solveOne(R, C int, puzzle []string) int {
	var a, b, c, d int
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if puzzle[i][j] == 'w' {
				return 0
			}
			if puzzle[i][j] == 'o' {
				a, b = i, j
			}
			if puzzle[i][j] == 'x' {
				c, d = i, j
			}
		}
	}

	que := make([]int, R*C)
	dist := make([]int, R*C)
	for i := 0; i < R*C; i++ {
		dist[i] = math.MaxInt32
	}

	dist[a*C+b] = 0

	front, end := 0, 0
	que[end] = a*C + b
	end++
	for front < end {
		cur := que[front]
		front++
		x, y := cur/C, cur%C
		if x == c && y == d {
			return dist[c*C+d]
		}

		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			p, q := x-dd[k], y-dd[k+1]
			if u < 0 || u >= R || v < 0 || v >= C || p < 0 || p >= R || q < 0 || q >= C {
				continue
			}
			if puzzle[u][v] != '#' && puzzle[p][q] != '#' && dist[u*C+v] > dist[cur]+1 {
				dist[u*C+v] = dist[cur] + 1
				que[end] = u*C + v
				end++
			}
		}
	}

	return -1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type State [][]int

func (s State) Len() int {
	return len(s)
}

func (s State) Less(i, j int) bool {
	return s[i][0] < s[j][0] || (s[i][0] == s[j][0] && s[i][1] < s[j][1])
}

func (s State) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s State) String() string {
	var buf bytes.Buffer
	for _, row := range s {
		buf.WriteString(fmt.Sprintf("{%d|%d}", row[0], row[1]))
	}
	return buf.String()
}

type Item struct {
	val  State
	dist int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].dist < pq[j].dist
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
