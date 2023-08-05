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

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		nums := readNNums(reader, n)
		res := solve(nums)
		buf.WriteString(fmt.Sprintf("%d ", len(res)))
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func solve(A []int) []int {
	// 这个题目要抽象一下
	// 如果 gcd(A[i], A[i+1]) = 1, 那么相当于在 i 和 i + 2之间连了一条线
	// 也就是如果能到达i, 下一个到达的位置就是i+2
	// 否则就是在i和i+1之间有一条线
	// 所以最后剩下的肯定是一个环（或者单点）
	// 如果到达了节点i，
	n := len(A)

	next := make([][]int, n)

	for i := 0; i < n; i++ {
		next[i] = []int{(i - 1 + n) % n, (i + 1) % n}
	}
	pq := make(PQ, 0, n)

	for i := 0; i < n; i++ {
		if gcd(A[i], A[(i+1)%n]) == 1 {
			heap.Push(&pq, Pair{0, i})
			i++
		}
	}

	var ans []int
	vis := make([]bool, n)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(Pair)
		u := cur.second
		if vis[u] {
			continue
		}
		v := next[u][1]
		vis[v] = true
		ans = append(ans, v+1)
		next[u][1] = next[v][1]
		next[next[v][1]][0] = u
		if gcd(A[u], A[next[u][1]]) == 1 {
			heap.Push(&pq, Pair{cur.first + 1, u})
		}
	}

	return ans
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type Pair struct {
	first  int
	second int
}

type PQ []Pair

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].first < pq[j].first || pq[i].first == pq[j].first && pq[i].second < pq[j].second
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Pair))
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}
