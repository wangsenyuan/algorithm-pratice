package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	moments := make([][][]int, m)
	for i := 0; i < m; i++ {
		k := readNum(reader)
		moments[i] = make([][]int, k)
		for j := 0; j < k; j++ {
			moments[i][j] = readNNums(reader, 2)
		}
	}
	k := readNum(reader)
	a := readNNums(reader, k)
	res := solve(n, moments, a)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

const oo = 1 << 60

func solve(n int, moments [][][]int, a []int) int {
	g := make([][]Pair, n)
	for i := 0; i < n; i++ {
		g[i] = make([]Pair, 0, 1)
	}

	for t, cur := range moments {
		for _, road := range cur {
			u, v := road[0]-1, road[1]-1
			g[u] = append(g[u], Pair{t, v})
			g[v] = append(g[v], Pair{t, u})
		}
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = oo
	}
	m := len(moments)
	pq := make([][]int, m)

	for i := 0; i < m; i++ {
		pq[i] = make([]int, 0, 1)
	}

	activate := func(u int) {
		for _, out := range g[u] {
			t, v := out.first, out.second
			pq[t] = append(pq[t], v)
		}
		g[u] = g[u][:0]
	}

	dp[0] = 0
	activate(0)
	upd := make([]int, n)
	for k, t := range a {
		t--
		var pos int
		for _, to := range pq[t] {
			if dp[to] > 1+k {
				dp[to] = 1 + k
				upd[pos] = to
				pos++
			}
		}
		pq[t] = make([]int, 0, 1)

		for i := 0; i < pos; i++ {
			activate(upd[i])
		}
	}

	if dp[n-1] == oo {
		return -1
	}
	return dp[n-1]
}

type Pair struct {
	first  int
	second int
}

// An Item is something we manage in a priority queue.
type Item struct {
	id       int
	priority int
	index    int
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
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}

func (pq *PriorityQueue) update(it *Item, p int) {
	it.priority = p
	heap.Fix(pq, it.index)
}
