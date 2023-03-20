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

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		blocks := make([][]int, n)
		for i := 0; i < n; i++ {
			blocks[i] = readNNums(reader, 4)
		}
		res := solve(k, blocks)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

const INF = 1 << 30

func solve(k int, blocks [][]int) int {

	vedge := make([][]int, k+1)
	hedge := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		vedge[i] = make([]int, k+1)
		hedge[i] = make([]int, k+1)
	}

	for _, blk := range blocks {
		for i := 0; i < len(blk); i++ {
			blk[i]--
		}
		r1, c1, r2, c2 := blk[0], blk[1], blk[2], blk[3]
		if r1 == r2 {
			if c1 > c2 {
				c1, c2 = c2, c1
			}
			vedge[r1][c2]++
			vedge[k-r1-1][k-c2]++
		} else {
			if r1 > r2 {
				r1, r2 = r2, r1
			}
			hedge[r2][c1]++
			hedge[k-r2][k-c1-1]++
		}
	}

	dist := make([][]*Item, k+1)
	pq := make(PQ, (k+1)*(k+1))

	for i := 0; i <= k; i++ {
		dist[i] = make([]*Item, k+1)
		for j := 0; j <= k; j++ {
			it := new(Item)
			it.x = i
			it.y = j
			it.dist = INF
			it.index = i*(k+1) + j
			dist[i][j] = it
			pq[it.index] = it
		}
	}

	heap.Init(&pq)

	pq.update(dist[0][0], 0)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		r, c, d := cur.x, cur.y, cur.dist
		if r > 0 && d+vedge[r-1][c] < dist[r-1][c].dist {
			pq.update(dist[r-1][c], d+vedge[r-1][c])
		}
		if r < k && d+vedge[r][c] < dist[r+1][c].dist {
			pq.update(dist[r+1][c], d+vedge[r][c])
		}
		if c > 0 && d+hedge[r][c-1] < dist[r][c-1].dist {
			pq.update(dist[r][c-1], d+hedge[r][c-1])
		}
		if c < k && d+hedge[r][c] < dist[r][c+1].dist {
			pq.update(dist[r][c+1], d+hedge[r][c])
		}
	}

	return len(blocks) - dist[k/2][k/2].dist
}

type Item struct {
	x     int
	y     int
	dist  int
	index int
}

type PQ []*Item

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

func (pq *PQ) Push(x interface{}) {
	i := x.(*Item)
	i.index = len(*pq)
	*pq = append(*pq, i)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	rt := old[n-1]
	rt.index = -1
	*pq = old[:n-1]
	return rt
}

func (pq *PQ) update(i *Item, d int) {
	i.dist = d
	heap.Fix(pq, i.index)
}
