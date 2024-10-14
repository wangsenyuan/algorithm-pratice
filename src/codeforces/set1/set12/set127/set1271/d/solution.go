package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	castles := make([][]int, n)
	for i := 0; i < n; i++ {
		castles[i] = readNNums(reader, 3)
	}
	portals := make([][]int, m)
	for i := 0; i < m; i++ {
		portals[i] = readNNums(reader, 2)
	}

	res := solve(k, castles, portals)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

const inf = 1 << 60

func solve(k int, castles [][]int, portals [][]int) int {
	n := len(castles)

	sum := k
	for _, cur := range castles {
		if sum < cur[0] {
			return -1
		}
		sum += cur[1]
	}

	add_to := make([]int, n)
	for i := 0; i < n; i++ {
		add_to[i] = i
	}

	for _, cur := range portals {
		u, v := cur[0], cur[1]
		u--
		v--
		add_to[v] = max(add_to[v], u)
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, sum+1)
		for j := 0; j <= sum; j++ {
			dp[i][j] = -inf
		}
	}
	dp[0][k] = 0
	for i, cur := range castles {
		for j := cur[0]; j+cur[1] <= sum; j++ {
			dp[i+1][j+cur[1]] = max(dp[i+1][j+cur[1]], dp[i][j])
		}
		for j := 0; j <= i; j++ {
			if add_to[j] == i {
				// 这个地方，可以保证，每个j，只会迭代一次，所以整体还是 n * n 的复杂性
				for s := 1; s <= sum; s++ {
					dp[i+1][s-1] = max(dp[i+1][s-1], dp[i+1][s]+castles[j][2])
				}
			}
		}
	}

	return slices.Max(dp[n])
}

func solve1(k int, castles [][]int, portals [][]int) int {
	n := len(castles)

	sum := k
	for _, cur := range castles {
		if sum < cur[0] {
			return -1
		}
		sum += cur[1]
	}

	add_to := make([]int, n)
	for i := 0; i < n; i++ {
		add_to[i] = i
	}

	for _, cur := range portals {
		u, v := cur[0], cur[1]
		u--
		v--
		add_to[v] = max(add_to[v], u)
	}

	g := make([][]int, n)

	for i := 0; i < n; i++ {
		j := add_to[i]
		g[j] = append(g[j], castles[i][2])
	}

	pq := make(IntHeap, 0, n)

	sum = k
	for i := range n {
		sum += castles[i][1]
		var next int
		if i+1 < n {
			next = castles[i+1][0]
		}
		for j := 0; j < len(g[i]); j++ {
			heap.Push(&pq, g[i][j])
		}

		for sum-len(pq) < next {
			// 撤回士兵
			heap.Pop(&pq)
		}
	}

	var res int
	for pq.Len() > 0 {
		res += heap.Pop(&pq).(int)
	}

	return res
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
