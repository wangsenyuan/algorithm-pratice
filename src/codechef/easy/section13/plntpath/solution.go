package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k, s := readThreeNums(reader)
	A := readNNums(reader, n)
	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = readNNums(reader, k)
	}
	res := solve(k, s, A, C)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
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

const INF = math.MaxInt64 >> 1

func solve(K, S int, A []int, C [][]int) []int64 {
	// K <= 1000
	D := make([][]int64, K)
	for i := 0; i < K; i++ {
		D[i] = make([]int64, K)
		for j := 0; j < K; j++ {
			D[i][j] = INF
		}
		D[i][i] = 0
	}
	n := len(A)
	for i := 0; i < n; i++ {
		x := A[i] - 1
		for j := 0; j < K; j++ {
			if C[i][j] >= 0 {
				D[x][j] = min(D[x][j], int64(C[i][j]))
			}
		}
	}

	nodes := make([]*Node, K)
	pq := make(PQ, K)
	for i := 0; i < K; i++ {
		tmp := new(Node)
		tmp.dist = INF
		tmp.index = i
		tmp.label = i
		nodes[i] = tmp
		pq[i] = tmp
	}

	S--

	nodes[A[S]-1].dist = 0
	heap.Init(&pq)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Node)
		for i := 0; i < K; i++ {
			if nodes[i].dist > cur.dist+D[cur.label][i] {
				pq.update(nodes[i], cur.dist+D[cur.label][i])
			}
		}
	}

	res := make([]int64, n)

	for i := 0; i < n; i++ {
		d := nodes[A[i]-1].dist
		if d < INF {
			res[i] = d
		} else {
			res[i] = -1
		}
	}
	return res
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type Node struct {
	dist  int64
	label int
	index int
}

type PQ []*Node

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(item interface{}) {
	cur := item.(*Node)
	cur.index = len(*pq)
	*pq = append(*pq, cur)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	ret := old[n-1]
	ret.index = -1
	*pq = old[:n-1]
	return ret
}

func (pq *PQ) update(node *Node, dist int64) {
	node.dist = dist
	heap.Fix(pq, node.index)
}
