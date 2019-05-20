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
		m, n := readTwoNums(scanner)
		x, y := readTwoNums(scanner)
		grid := make([]string, m)
		for i := 0; i < m; i++ {
			scanner.Scan()
			grid[i] = scanner.Text()
		}
		res := solve(m, n, x, y, grid)

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if j < n-1 {
					fmt.Printf("%d ", res[i][j])
				} else {
					fmt.Printf("%d\n", res[i][j])
				}
			}
		}
	}
}

const INF = 1 << 20

func solve(m, n, x, y int, grid []string) [][]int {
	res := make([][][]int, m)

	for i := 0; i < m; i++ {
		res[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			res[i][j] = make([]int, 3)
			for k := 0; k < 3; k++ {
				res[i][j][k] = INF
			}
		}
	}
	res[x-1][y-1][0] = 0
	que := make([]int, m*n*3)
	var front int
	var tail int
	que[tail] = (x-1)*n*3 + (y-1)*3
	tail++

	for front < tail {
		cur := que[front]
		front++
		r := cur / (3 * n)
		c := (cur % (3 * n)) / 3
		state := (cur % (3 * n)) % 3

		if res[r][c][state] < 0 {
			// can't go any more
			break
		}

		if state == 0 {
			if c+2 < n && grid[r][c+1] == '1' && grid[r][c+2] == '1' {
				// right
				if res[r][c+1][1] > res[r][c][0]+1 {
					res[r][c+1][1] = res[r][c][0] + 1
					que[tail] = r*(3*n) + (c+1)*3 + 1
					tail++
				}
			}
			if r+2 < m && grid[r+1][c] == '1' && grid[r+2][c] == '1' {
				if res[r+1][c][2] > res[r][c][0]+1 {
					res[r+1][c][2] = res[r][c][0] + 1
					que[tail] = (r+1)*(3*n) + c*3 + 2
					tail++
				}
			}
			if c-2 >= 0 && grid[r][c-1] == '1' && grid[r][c-2] == '1' {
				if res[r][c-2][1] > res[r][c][0]+1 {
					res[r][c-2][1] = res[r][c][0] + 1
					que[tail] = r*(3*n) + (c-2)*3 + 1
					tail++
				}
			}

			if r-2 >= 0 && grid[r-2][c] == '1' && grid[r-1][c] == '1' {
				if res[r-2][c][2] > res[r][c][0]+1 {
					res[r-2][c][2] = res[r][c][0] + 1
					que[tail] = (r-2)*3*n + c*3 + 2
					tail++
				}
			}
		} else if state == 1 {
			// right
			if c+2 < n && grid[r][c+2] == '1' && res[r][c+2][0] > res[r][c][1]+1 {
				res[r][c+2][0] = res[r][c][1] + 1
				que[tail] = r*3*n + (c+2)*3
				tail++
			}
			// down
			if r+1 < m && grid[r+1][c] == '1' && grid[r+1][c+1] == '1' && res[r+1][c][1] > res[r][c][1]+1 {
				res[r+1][c][1] = res[r][c][1] + 1
				que[tail] = (r+1)*3*n + c*3 + 1
				tail++
			}

			//left
			if c-1 >= 0 && grid[r][c-1] == '1' && res[r][c-1][0] > res[r][c][1]+1 {
				res[r][c-1][0] = res[r][c][1] + 1
				que[tail] = r*3*n + (c-1)*3
				tail++
			}

			//up
			if r-1 >= 0 && grid[r-1][c] == '1' && grid[r-1][+1] == '1' && res[r-1][c][1] > res[r][c][1]+1 {
				res[r-1][c][1] = res[r][c][1] + 1
				que[tail] = (r-1)*3*n + c*3 + 1
				tail++
			}
		} else {
			// state == 2
			// right
			if c+1 < n && grid[r][c+1] == '1' && grid[r+1][c+1] == '1' && res[r][c+1][2] > res[r][c][2]+1 {
				res[r][c+1][2] = res[r][c][2] + 1
				que[tail] = r*3*n + (c+1)*3 + 2
				tail++
			}

			if r+2 < m && grid[r+2][c] == '1' && res[r+2][c][0] > res[r][c][2]+1 {
				res[r+2][c][0] = res[r][c][2] + 1
				que[tail] = (r+2)*3*n + c*3
				tail++
			}

			if c-1 >= 0 && grid[r][c-1] == '1' && grid[r+1][c-1] == '1' && res[r][c-1][2] > res[r][c][2]+1 {
				res[r][c-1][2] = res[r][c][2] + 1
				que[tail] = r*3*n + (c-1)*3 + 2
				tail++
			}
			if r-1 >= 0 && grid[r-1][c] == '1' && res[r-1][c][0] > res[r][c][2]+1 {
				res[r-1][c][0] = res[r][c][2] + 1
				que[tail] = (r-1)*3*n + c*3
				tail++
			}
		}
	}

	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ans[i][j] = -1
			if res[i][j][0] < INF {
				ans[i][j] = res[i][j][0]
			}
		}
	}

	return ans
}

func solve1(m, n int, x int, y int, grid []string) [][]int {
	res := make([][]*Item, m)
	pq := make(PriorityQueue, 0, m*n)
	heap.Init(&pq)
	for i := 0; i < m; i++ {
		res[i] = make([]*Item, n)
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res[i][j] = new(Item)
				res[i][j].first = INF
				res[i][j].second = i*n + j
				heap.Push(&pq, res[i][j])
			}
		}
	}
	x--
	y--
	res[x][y].first = 0
	heap.Fix(&pq, res[x][y].index)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		if cur.first == INF {
			break
		}
		u, v := cur.second/n, cur.second%n

		// roll down up
		if u-2 >= 0 && grid[u-2][v] == '1' && grid[u-1][v] == '1' {
			rollup(u, grid, v, res, &pq, n)
		}

		if u+2 < m && grid[u+1][v] == '1' && grid[u+2][v] == '1' {
			rolldown(u, grid, v, res, &pq, m, n)
		}

		if v-2 >= 0 && grid[u][v-2] == '1' && grid[u][v-1] == '1' {
			rollleft(grid, u, v, res, &pq, m, n)
		}

		if v+2 < n && grid[u][v+2] == '1' && grid[u][v+1] == '1' {
			rollright(grid, u, v, res, &pq, m, n)
		}
	}

	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ans[i][j] = -1
			if res[i][j] != nil && res[i][j].first < INF {
				ans[i][j] = res[i][j].first
			}
		}
	}

	return ans
}

func rollright(grid []string, u int, v int, res [][]*Item, queues *PriorityQueue, m int, n int) {
	if v+3 < n && grid[u][v+3] == '1' && res[u][v+3].first > res[u][v].first+2 {
		res[u][v+3].first = res[u][v].first + 2
		heap.Fix(queues, res[u][v+3].index)
	}

	for x := u - 1; x >= 0; x-- {
		if grid[x][v+2] == '0' || grid[x][v+1] == '0' {
			break
		}

		if v+3 < n && grid[x][v+3] == '1' && res[x][v+3].first > res[u][v].first+u-x+2 {
			res[x][v+3].first = res[u][v].first + 2 + u - x
			heap.Fix(queues, res[x][v+3].index)
		}

		if grid[x][v] == '1' && res[x][v].first > res[u][v].first+u-x+2 {
			res[x][v].first = res[u][v].first + 2 + u - x
			heap.Fix(queues, res[x][v].index)
		}
	}

	for x := u + 1; x < m; x++ {
		if grid[x][v+2] == '0' || grid[x][v+1] == '0' {
			break
		}

		if v+3 < n && grid[x][v+3] == '1' && res[x][v+3].first > res[u][v].first+x-u+2 {
			res[x][v+3].first = res[u][v].first + 2 + x - u
			heap.Fix(queues, res[x][v+3].index)
		}

		if grid[x][v] == '1' && res[x][v].first > res[u][v].first+x-u+2 {
			res[x][v].first = res[u][v].first + 2 + x - u
			heap.Fix(queues, res[x][v].index)
		}
	}
}

func rollleft(grid []string, u int, v int, res [][]*Item, queues *PriorityQueue, m int, n int) {
	if v-3 >= 0 && grid[u][v-3] == '1' && res[u][v-3].first > res[u][v].first+2 {
		res[u][v-3].first = res[u][v].first + 2
		heap.Fix(queues, res[u][v-3].index)
	}

	for x := u - 1; x >= 0; x-- {
		if grid[x][v-2] == '0' || grid[x][v-1] == '0' {
			break
		}

		if v-3 >= 0 && grid[x][v-3] == '1' && res[x][v-3].first > res[u][v].first+u-x+2 {
			res[x][v-3].first = res[u][v].first + 2 + u - x
			heap.Fix(queues, res[x][v-3].index)
		}

		if grid[x][v] == '1' && res[x][v].first > res[u][v].first+u-x+2 {
			res[x][v].first = res[u][v].first + 2 + u - x
			heap.Fix(queues, res[x][v].index)
		}
	}

	for x := u + 1; x < m; x++ {
		if grid[x][v-2] == '0' || grid[x][v-1] == '0' {
			break
		}

		if v-3 >= 0 && grid[x][v-3] == '1' && res[x][v-3].first > res[u][v].first+x-u+2 {
			res[x][v-3].first = res[u][v].first + 2 + x - u
			heap.Fix(queues, res[x][v-3].index)
		}

		if grid[x][v] == '1' && res[x][v].first > res[u][v].first+x-u+2 {
			res[x][v].first = res[u][v].first + 2 + x - u
			heap.Fix(queues, res[x][v].index)
		}
	}
}

func rolldown(u int, grid []string, v int, res [][]*Item, pq *PriorityQueue, m, n int) {
	// can lay down
	if u+3 < m && grid[u+3][v] == '1' && res[u+3][v].first > res[u][v].first+2 {
		res[u+3][v].first = res[u][v].first + 2
		heap.Fix(pq, res[u+3][v].index)
	}
	for y := v - 1; y >= 0; y-- {
		if grid[u+1][y] == '0' || grid[u+2][v] == '0' {
			break
		}
		if u+3 < m && grid[u+3][y] == '1' && res[u+3][y].first > res[u][v].first+v-y+2 {
			res[u+3][y].first = res[u][v].first + v - y + 2
			heap.Fix(pq, res[u+3][y].index)
		}

		if grid[u][y] == '1' && res[u][y].first > res[u][v].first+v-y+2 {
			res[u][y].first = res[u][v].first + v - y + 2
			heap.Fix(pq, res[u][y].index)
		}
	}
	for y := v + 1; y < n; y++ {
		if grid[u+2][y] == '0' || grid[u+1][v] == '0' {
			break
		}
		if u+3 >= 0 && grid[u+3][y] == '1' && res[u+3][y].first > res[u][v].first+y-v+2 {
			res[u+3][y].first = res[u][v].first + y - v + 2
			heap.Fix(pq, res[u+3][y].index)
		}

		if grid[u][y] == '1' && res[u][y].first > res[u][v].first+y-v+2 {
			res[u][y].first = res[u][v].first + y - v + 2
			heap.Fix(pq, res[u][y].index)
		}
	}
}

func rollup(u int, grid []string, v int, res [][]*Item, pq *PriorityQueue, n int) {
	// can lay down
	if u-3 >= 0 && grid[u-3][v] == '1' && res[u-3][v].first > res[u][v].first+2 {
		res[u-3][v].first = res[u][v].first + 2
		heap.Fix(pq, res[u-3][v].index)
	}
	for y := v - 1; y >= 0; y-- {
		if grid[u-2][y] == '0' || grid[u-1][v] == '0' {
			break
		}
		if u-3 >= 0 && grid[u-3][y] == '1' && res[u-3][y].first > res[u][v].first+v-y+2 {
			res[u-3][y].first = res[u][v].first + v - y + 2
			heap.Fix(pq, res[u-3][y].index)
		}

		if grid[u][y] == '1' && res[u][y].first > res[u][v].first+v-y+2 {
			res[u][y].first = res[u][v].first + v - y + 2
			heap.Fix(pq, res[u][y].index)
		}
	}
	for y := v + 1; y < n; y++ {
		if grid[u-2][y] == '0' || grid[u-1][v] == '0' {
			break
		}
		if u-3 >= 0 && grid[u-3][y] == '1' && res[u-3][y].first > res[u][v].first+y-v+2 {
			res[u-3][y].first = res[u][v].first + y - v + 2
			heap.Fix(pq, res[u-3][y].index)
		}

		if grid[u][y] == '1' && res[u][y].first > res[u][v].first+y-v+2 {
			res[u][y].first = res[u][v].first + y - v + 2
			heap.Fix(pq, res[u][y].index)
		}
	}
}

type Item struct {
	first  int
	second int
	index  int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].first < pq[j].first
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(item interface{}) {
	n := len(*pq)
	x := item.(*Item)
	x.index = n
	*pq = append(*pq, x)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	res := (*pq)[n-1]
	res.index = -1
	*pq = (*pq)[:n-1]
	return res
}
