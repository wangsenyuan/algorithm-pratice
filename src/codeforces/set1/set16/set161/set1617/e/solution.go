package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := 1

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)

		buf.WriteString(fmt.Sprintf("%d %d %d\n", res[0], res[1], res[2]))
	}
	fmt.Print(buf.String())
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(A []int) []int {
	// x, y, 它们之间的步骤等于通过
	// x0, x1, x2, .. xn-1
	// x0 = x, xn-1 = y
	// 且 x0 + x1 = 2 ** k
	// 首先同一个数字，在这条path上不会出现两次
	// 那么意味这k不会重复，是否可以在greedy的同时，在21的的时间内计算出呢?
	// 而且路径是对称的
	//sort.Ints(A)
	// 有点像lca的形式，假设x, 到达某个数字p 为a， y到达某个数字为b， 那么x和y就是a + b
	// 这时候为了找到最大的最小值，只需要找到这棵树的直径
	// n个叶子节点，那么最多有 3 * n 个内部节点
	// 所有的叶子节点，都变成 2 ** k - a
	n := len(A)

	dist := make(map[int][]Pair)

	pq := make(IntHeap, 0, n)

	for i, v := range A {
		heap.Push(&pq, v)
		dist[v] = make([]Pair, 2)
		dist[v][0] = Pair{0, i + 1}
	}

	for pq.Len() > 0 {
		od := heap.Pop(&pq).(int)
		if pq.Len() > 0 && pq[0] == od {
			continue
		}
		if od != 1 {
			nx := get(od)
			heap.Push(&pq, nx)
			second := dist[od][0].AddFirst(1)
			if dist[nx] == nil {
				dist[nx] = make([]Pair, 2)
			}
			if dist[nx][0].first < second.first {
				dist[nx][1] = dist[nx][0]
				dist[nx][0] = second
			} else if dist[nx][1].first < second.first {
				dist[nx][1] = second
			}
		}
	}

	var ra, rb int
	rc := -1
	for _, tg := range dist {
		if tg[0].second == 0 || tg[1].second == 0 {
			continue
		}
		if rc < tg[0].first+tg[1].first {
			rc = tg[0].first + tg[1].first
			ra = tg[0].second
			rb = tg[1].second
		}
	}
	return []int{ra, rb, rc}
}

func get(x int) int {
	for i := 0; ; i++ {
		if 1<<i >= x {
			return (1 << i) - x
		}
	}
	return x
}

type Pair struct {
	first  int
	second int
}

func (this Pair) AddFirst(v int) Pair {
	return Pair{this.first + v, this.second}
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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
