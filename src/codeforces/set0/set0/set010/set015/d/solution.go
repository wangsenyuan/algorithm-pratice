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

	first_line := readNNums(reader, 4)
	n, m, a, b := first_line[0], first_line[1], first_line[2], first_line[3]
	H := make([][]int, n)
	for i := 0; i < n; i++ {
		H[i] = readNNums(reader, m)
	}

	pos, ans := solve(n, m, a, b, H)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(pos)))

	for i := 0; i < len(pos); i++ {
		buf.WriteString(fmt.Sprintf("%d %d %d\n", pos[i][0], pos[i][1], ans[i]))
		if buf.Len() > 10000 {
			fmt.Print(buf.String())
			buf.Reset()
		}
	}
	fmt.Print(buf.String())
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

const INF = 2 * 1e9

func solve(n int, m int, a int, b int, H [][]int) ([][]int, []int64) {

	var pos [][]int
	var ans []int64

	if a == 1 && b == 1 {
		for i := 1; i <= n; i++ {
			for j := 1; j <= m; j++ {
				pos = append(pos, []int{i, j})
				ans = append(ans, 0)
			}
		}
		return pos, ans
	}
	sum := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		sum[i] = make([]int64, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + int64(H[i-1][j-1])
		}
	}

	mnv := make([][]int, n+1)
	mn := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		mnv[i] = make([]int, m+1)
		mn[i] = make([]int, m+1)
	}

	q := NewDeque(max(m, n))
	for i := 1; i <= n; i++ {
		q.Reset()
		for j := 1; j <= m; j++ {
			for q.Size() > 0 && H[i-1][q.End()-1] > H[i-1][j-1] {
				q.PopEnd()
			}
			q.PushEnd(j)
			for q.Front() <= j-b {
				q.PopFront()
			}
			mnv[i][j] = H[i-1][q.Front()-1]
		}
	}

	for j := 1; j <= m; j++ {
		q.Reset()
		for i := 1; i <= n; i++ {
			for q.Size() > 0 && mnv[q.End()][j] > mnv[i][j] {
				q.PopEnd()
			}
			q.PushEnd(i)
			for q.Front() <= i-a {
				q.PopFront()
			}
			mn[i][j] = mnv[q.Front()][j]
		}
	}

	pq := make(PQ, 0, m*n-(n*b)-(m*a)+a*b)

	used := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		used[i] = make([]bool, m+1)
	}

	for i := a; i <= n; i++ {
		for j := b; j <= m; j++ {
			area := sum[i][j] - sum[i-a][j] - sum[i][j-b] + sum[i-a][j-b]
			area -= int64(a) * int64(b) * int64(mn[i][j])
			cell := Cell{area, i, j}
			heap.Push(&pq, cell)
		}
	}

	for pq.Len() > 0 {
		tmp := heap.Pop(&pq).(Cell)
		x := tmp.i
		y := tmp.j
		if used[x][y] || used[x-a+1][y] || used[x][y-b+1] || used[x-a+1][y-b+1] {
			continue
		}
		ans = append(ans, tmp.area)
		pos = append(pos, []int{x - a + 1, y - b + 1})
		for r := x - a + 1; r <= x; r++ {
			for c := y - b + 1; c <= y; c++ {
				used[r][c] = true
			}
		}
	}

	return pos, ans
}

type Cell struct {
	area int64
	i    int
	j    int
}

type PQ []Cell

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].area < pq[j].area || pq[i].area == pq[j].area && pq[i].i < pq[j].i || pq[i].area == pq[j].area && pq[i].i == pq[j].i && pq[i].j < pq[j].j
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Cell))
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	ret := old[n-1]
	*pq = old[:n-1]
	return ret
}

type Deque struct {
	arr   []int
	sz    int
	front int
	end   int
}

func NewDeque(sz int) *Deque {
	sz++
	arr := make([]int, sz)

	return &Deque{arr, sz, 0, 0}
}

func (q *Deque) Reset() {
	q.end = 0
	q.front = 0
}

func (q *Deque) Size() int {
	return (q.end - q.front + q.sz) % q.sz
}

func (q *Deque) PushFront(v int) {
	q.front = (q.front - 1 + q.sz) % q.sz
	q.arr[q.front] = v
}

func (q *Deque) Front() int {
	return q.arr[q.front]
}

func (q *Deque) PopFront() int {
	ret := q.arr[q.front]
	q.front = (q.front + 1) % q.sz
	return ret
}

func (q *Deque) PushEnd(v int) {
	q.arr[q.end] = v
	q.end = (q.end + 1) % q.sz
}

func (q *Deque) End() int {
	return q.arr[(q.end-1+q.sz)%q.sz]
}

func (q *Deque) PopEnd() int {
	ret := q.arr[(q.end-1+q.sz)%q.sz]
	q.end = (q.end - 1 + q.sz) % q.sz
	return ret
}

type Number interface {
	int | int64
}

func max[T Number](a, b T) T {
	if a >= b {
		return a
	}
	return b
}
