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
	ss := make([]string, n)
	for i := 0; i < n; i++ {
		ss[i] = readString(reader)
	}
	res := solve(ss)
	fmt.Println(res)
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

func solve(ord []string) string {
	items := make(map[string]*Item)
	pq := make(PQ, 0, len(ord))

	for i, s := range ord {
		if it, ok := items[s]; ok {
			it.priority = i
			heap.Fix(&pq, it.index)
		} else {
			it = new(Item)
			it.name = s
			it.priority = i
			heap.Push(&pq, it)
			items[s] = it
		}
	}

	var buf bytes.Buffer

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		buf.WriteString(last(it.name))
	}

	return buf.String()
}

func last(s string) string {
	n := len(s)
	return s[n-2:]
}

type Item struct {
	name     string
	priority int
	index    int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	it := x.(*Item)
	it.index = len(*pq)
	*pq = append(*pq, it)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	res := old[n-1]
	res.index = -1
	*pq = old[:n-1]
	return res
}
