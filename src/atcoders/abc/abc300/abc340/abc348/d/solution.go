package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) bool {
	n, m := readTwoNums(reader)
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = readString(reader)[:m]
	}
	k := readNum(reader)
	E := make([][]int, k)
	for i := 0; i < k; i++ {
		E[i] = readNNums(reader, 3)
	}
	return solve(a, E)
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(a []string, E [][]int) bool {
	n := len(a)
	m := len(a[0])

	val := make([][]int, n)

	start := []int{0, 0}
	dest := []int{0, 0}

	for i := 0; i < n; i++ {
		val[i] = make([]int, m)
		for j := 0; j < m; j++ {
			val[i][j] = -1
			if a[i][j] == 'T' {
				dest = []int{i, j}
			}
			if a[i][j] == 'S' {
				start = []int{i, j}
			}
		}
	}
	for _, cur := range E {
		a, b, c := cur[0], cur[1], cur[2]
		a--
		b--
		val[a][b] = max(val[a][b], c)
	}

	if val[start[0]][start[1]] <= 0 {
		return false
	}

	// 300 * 200 * 200 = 12 * 1e6 = 1e7 可行的
	var D [][]int
	var first int
	second := -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if val[i][j] > 0 && a[i][j] != '#' {
				D = append(D, []int{i, j})
				if i == start[0] && j == start[1] {
					first = len(D) - 1
				}
				if i == dest[0] && j == dest[1] {
					second = i
				}
			}
		}
	}

	if second < 0 {
		D = append(D, dest)
		second = len(D) - 1
	}

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
	}
	que := make([]int, n*m)
	// len(D) <= len(E) <= 300
	bfs := func(sx int, sy int) []bool {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				dist[i][j] = -1
			}
		}
		dist[sx][sy] = 0
		var head, tail int
		que[head] = sx*m + sy
		head++
		for tail < head {
			r, c := que[tail]/m, que[tail]%m
			tail++
			for i := 0; i < 4; i++ {
				x, y := r+dd[i], c+dd[i+1]
				if x >= 0 && x < n && y >= 0 && y < m && a[x][y] != '#' && dist[x][y] < 0 {
					dist[x][y] = dist[r][c] + 1
					que[head] = x*m + y
					head++
				}
			}
		}
		res := make([]bool, len(D))
		for i, cur := range D {
			r, c := cur[0], cur[1]
			if dist[r][c] < 0 || dist[r][c] > val[sx][sy] {
				continue
			}
			res[i] = true
		}
		return res
	}

	dists := make([][]bool, len(D))
	for i, cur := range D {
		dists[i] = bfs(cur[0], cur[1])
	}

	var head, tail int
	ok := make([]bool, len(D))
	ok[first] = true
	que[head] = first
	head++

	for tail < head {
		u := que[tail]
		tail++
		for i := 0; i < len(D); i++ {
			if dists[u][i] && !ok[i] {
				que[head] = i
				head++
				ok[i] = true
			}
		}
	}

	return ok[second]
}

func solve1(a []string, E [][]int) bool {
	n := len(a)
	m := len(a[0])

	val := make([][]int, n)

	start := []int{0, 0}
	dest := []int{0, 0}

	for i := 0; i < n; i++ {
		val[i] = make([]int, m)
		for j := 0; j < m; j++ {
			val[i][j] = -1
			if a[i][j] == 'T' {
				dest = []int{i, j}
			}
			if a[i][j] == 'S' {
				start = []int{i, j}
			}
		}
	}
	for _, cur := range E {
		a, b, c := cur[0], cur[1], cur[2]
		a--
		b--
		val[a][b] = max(val[a][b], c)
	}

	if val[start[0]][start[1]] <= 0 {
		return false
	}

	// 300 * 200 * 200 = 12 * 1e6 = 1e7 可行的
	var D [][]int
	var first int
	second := -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if val[i][j] > 0 && a[i][j] != '#' {
				D = append(D, []int{i, j})
				if i == start[0] && j == start[1] {
					first = len(D) - 1
				}
				if i == dest[0] && j == dest[1] {
					second = i
				}
			}
		}
	}

	if second < 0 {
		D = append(D, dest)
		second = len(D) - 1
	}

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
	}
	que := make([]int, n*m)
	// len(D) <= len(E) <= 300
	bfs := func(sx int, sy int) []int {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				dist[i][j] = -1
			}
		}
		dist[sx][sy] = 0
		var head, tail int
		que[head] = sx*m + sy
		head++
		for tail < head {
			r, c := que[tail]/m, que[tail]%m
			tail++
			for i := 0; i < 4; i++ {
				x, y := r+dd[i], c+dd[i+1]
				if x >= 0 && x < n && y >= 0 && y < m && a[x][y] != '#' && dist[x][y] < 0 {
					dist[x][y] = dist[r][c] + 1
					que[head] = x*m + y
					head++
				}
			}
		}
		res := make([]int, len(D))
		for i, cur := range D {
			r, c := cur[0], cur[1]
			res[i] = dist[r][c]
		}
		return res
	}

	dists := make([][]int, len(D))
	for i, cur := range D {
		dists[i] = bfs(cur[0], cur[1])
	}

	// 现在计算每个医疗点能够到达的最大值
	active := make(PriorityQueue, len(D))
	items := make([]*Item, len(D))
	for i := range D {
		it := new(Item)
		it.id = i
		it.priority = -1
		it.index = i
		items[i] = it
		active[i] = it
	}
	heap.Init(&active)
	active.update(items[first], val[D[first][0]][D[first][1]])

	for active.Len() > 0 {
		it := heap.Pop(&active).(*Item)
		u := it.id
		for i := 0; i < len(D); i++ {
			d := dists[u][i]
			if d < 0 || it.priority-d < 0 || items[i].index == -1 {
				continue
			}
			// can reach
			r, c := D[i][0], D[i][1]
			val[r][c] = max(val[r][c], it.priority-d)
			if items[i].priority < val[r][c] {
				active.update(items[i], val[r][c])
			}
		}
	}
	return val[dest[0]][dest[1]] >= 0
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Item struct {
	id       int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	it := x.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	it := old[n-1]
	*pq = old[:n-1]
	it.index = -1
	return it
}

func (pq *PriorityQueue) update(it *Item, v int) {
	it.priority = v
	heap.Fix(pq, it.index)
}
