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
		m, n, k := readThreeNums(reader)
		A := make([][]int, m)
		for i := 0; i < m; i++ {
			A[i] = readNNums(reader, n)
		}
		D := make([][]byte, m)
		for i := 0; i < m; i++ {
			D[i], _ = reader.ReadBytes('\n')
		}
		res := solve(m, n, k, A, D)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const INF = 1 << 30

var dx = []int{-1, 0, 1, 0, -1, 0, 1, 0}
var dy = []int{0, 1, 0, -1, 0, 1, 0, -1}

func solve(m, n, k int, A [][]int, D [][]byte) int {
	pq := make(PQ, 0, m*n)
	table := make([][]*Item, m)
	for i := 0; i < m; i++ {
		table[i] = make([]*Item, n)
		for j := 0; j < n; j++ {
			table[i][j] = new(Item)
			table[i][j].value = i*n + j
		}
	}

	dir := func(i, j int) int {
		if D[i][j] == 'U' {
			return 0
		}
		if D[i][j] == 'R' {
			return 1
		}
		if D[i][j] == 'D' {
			return 2
		}
		return 3
	}

	check := func(x int) bool {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				table[i][j].priority = INF
				heap.Push(&pq, table[i][j])
			}
		}
		pq.update(table[0][0], 0)

		for pq.Len() > 0 {
			cur := heap.Pop(&pq).(*Item)
			cx, cy := cur.value/n, cur.value%n
			j := dir(cx, cy)
			for i := 0; i < 4; i++ {
				nx := cx + dx[j+i]
				ny := cy + dy[j+i]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && A[cx][cy]-i >= x {
					if table[nx][ny].priority > cur.priority+i {
						pq.update(table[nx][ny], cur.priority+i)
					}
				}
			}
		}

		return table[m-1][n-1].priority <= k
	}

	var left, right = 0, 1000000007

	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return max(0, right-1)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Item struct {
	value    int
	priority int
	index    int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(item interface{}) {
	cur := item.(*Item)
	cur.index = len(*pq)
	*pq = append(*pq, cur)
}

func (pq *PQ) Pop() interface{} {
	arr := *pq
	n := len(arr)
	res := arr[n-1]
	*pq = arr[:n-1]
	res.index = -1
	return res
}

func (pq *PQ) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
