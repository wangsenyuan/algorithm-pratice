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

	n := readNum(reader)
	Q := make([][]int, n)
	for i := 0; i < n; i++ {
		var tp int
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &tp)
		if tp == 1 {
			var m int
			readInt(s, pos+1, &m)
			Q[i] = []int{1, m}
		} else {
			Q[i] = []int{tp}
		}
	}
	res := solve(Q)
	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}

	buf.WriteByte('\n')

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
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

type Customer struct {
	id    int
	money int
	index int
}

type PQ []*Customer

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].money > pq[j].money || pq[i].money == pq[j].money && pq[i].id < pq[j].id
}

func (pq *PQ) Push(x interface{}) {
	c := x.(*Customer)
	c.index = len(*pq)
	*pq = append(*pq, c)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	rt := old[n-1]
	*pq = old[:n-1]
	rt.index = -1
	return rt
}

func (pq *PQ) update(c *Customer, money int) {
	c.money = money
	heap.Fix(pq, c.index)
}

func solve(Q [][]int) []int {
	var res []int

	id := 1

	pq := make(PQ, 0, len(Q))

	que := make([]*Customer, len(Q))
	var front, end int

	for _, cur := range Q {
		if cur[0] == 1 {
			cus := &Customer{id, cur[1], -1}
			heap.Push(&pq, cus)
			que[end] = cus
			end++
			id++
		} else if cur[0] == 2 {
			for front < end && que[front].index < 0 {
				front++
			}
			cus := que[front]
			pq.update(cus, 1<<30)
			heap.Pop(&pq)
			res = append(res, cus.id)
			front++
		} else {
			cus := heap.Pop(&pq).(*Customer)
			res = append(res, cus.id)
		}
	}

	return res
}
