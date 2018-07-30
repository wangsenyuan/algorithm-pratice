package main

import (
	"container/heap"
	"fmt"
	"os"
	"bufio"
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
		N, M := readTwoNums(scanner)
		G := make([][]byte, N)
		for i := 0; i < N; i++ {
			scanner.Scan()
			G[i] = scanner.Bytes()
		}
		res := solve(N, M, G)
		fmt.Printf("Case #%d: %d\n", x, res)
	}
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(N int, M int, G [][]byte) int {
	cells := make([][]*Cell, N)
	for i := 0; i < N; i++ {
		cells[i] = make([]*Cell, M)
		for j := 0; j < M; j++ {
			cells[i][j] = &Cell{cost: -1, x: i, y: j}
		}
	}

	parent := make([]int, N*M)

	res := make([][]int, N)
	for i := 0; i < N; i++ {
		res[i] = make([]int, M)
		for j := 0; j < M; j++ {
			res[i][j] = -1
		}
	}

	vis := make([][]int, N)
	for i := 0; i < N; i++ {
		vis[i] = make([]int, M)
	}

	que := make([]int, N*M)
	front, end := 0, 0
	que[end] = 0
	end++

	pq := make(Cells, 0, N*M)

	bfs := func(flag int) {
		for pq.Len() > 0 {
			u := heap.Pop(&pq).(*Cell)
			if vis[u.x][u.y] == flag {
				continue
			}
			vis[u.x][u.y] = flag

			for k := 0; k < 4; k++ {
				a, b := u.x+dd[k], u.y+dd[k+1]
				if a >= 0 && a < N && b >= 0 && b < M && G[a][b] != '.' {
					if cells[a][b].cost < 0 || cells[a][b].cost > cells[u.x][u.y].cost+1 {
						cells[a][b].cost = cells[u.x][u.y].cost + 1
						heap.Push(&pq, cells[a][b])
						parent[a*M+b] = u.x*M + u.y
						if G[a][b] == 'T' {
							v := a*M + b
							for {
								if res[v/M][v%M] < 0 {
									res[v/M][v%M] = cells[v/M][v%M].cost
								}
								v = parent[v]
								if G[v/M][v%M] == 'T' {
									break
								}
							}
							que[end] = a*M + b
							end++
							return
						}
					}
				}
			}
		}
	}

	bfs2 := func(flag int) {
		for pq.Len() > 0 {
			u := heap.Pop(&pq).(*Cell)
			if vis[u.x][u.y] == flag {
				continue
			}
			vis[u.x][u.y] = flag
			for k := 0; k < 4; k++ {
				a, b := u.x+dd[k], u.y+dd[k+1]
				if a >= 0 && a < N && b >= 0 && b < M && G[a][b] != '.' {
					if cells[a][b].cost < 0 || cells[a][b].cost > cells[u.x][u.y].cost+1 {
						cells[a][b].cost = cells[u.x][u.y].cost + 1
						heap.Push(&pq, cells[a][b])
					}
				}
			}
		}
	}

	var flag int
	for front < end {
		head := que[front]
		front++
		x, y := head/M, head%M
		cell := cells[x][y]
		cell.cost = 0
		heap.Push(&pq, cell)
		flag++
		bfs(flag)
	}

	flag++
	bfs2(flag)

	var ans int
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if res[i][j] < 0 {
				res[i][j] = cells[i][j].cost
			}
			if G[i][j] != '.' {
				ans += res[i][j]
			}
		}
	}
	//fmt.Fprintf(os.Stderr, "[debug]cost result %v\n", res)
	return ans
}

type Cell struct {
	cost  int
	x     int
	y     int
	index int
}

type Cells []*Cell

func (cells Cells) Len() int {
	return len(cells)
}

func (cells Cells) Less(i, j int) bool {
	return cells[i].cost < cells[j].cost
}

func (cells Cells) Swap(i, j int) {
	cells[i], cells[j] = cells[j], cells[i]
	cells[i].index = i
	cells[j].index = j
}

func (cells *Cells) Push(item interface{}) {
	cell := item.(*Cell)
	cell.index = len(*cells)
	*cells = append(*cells, cell)
}

func (cells *Cells) Pop() interface{} {
	old := *cells
	n := len(old)
	last := old[n-1]
	*cells = old[:n-1]
	last.index = -1
	return last
}
