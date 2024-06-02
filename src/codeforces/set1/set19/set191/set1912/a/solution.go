package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	x, n := readTwoNums(reader)
	lists := make([][]int, n)

	for i := 0; i < n; i++ {
		s, _ := reader.ReadBytes('\n')
		var m int
		pos := readInt(s, 0, &m) + 1
		lists[i] = make([]int, m)
		for j := 0; j < m; j++ {
			pos = readInt(s, pos, &lists[i][j]) + 1
		}
	}

	res := solve(x, lists)

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

func solve(x int, lists [][]int) int {

	// 只关心第一次拿到正数收益的情况
	getNextValue := func(i int, pos int) *Item {
		var sum int
		mn := 1 << 60
		for pos < len(lists[i]) {
			sum += lists[i][pos]
			mn = min(mn, sum)
			if sum >= 0 {
				it := new(Item)
				it.id = i
				it.pos = pos + 1
				it.expect = -mn
				it.gain = sum
				return it
			}
			pos++
		}
		return nil
	}

	n := len(lists)

	pq := make(PQ, 0, n)

	for i := 0; i < n; i++ {
		it := getNextValue(i, 0)
		if it != nil {
			heap.Push(&pq, it)
		}
	}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)
		if it.expect > x {
			// can't proceed any more
			break
		}
		x += it.gain
		it = getNextValue(it.id, it.pos)
		if it != nil {
			heap.Push(&pq, it)
		}
	}

	return x
}

type Item struct {
	id     int
	pos    int
	expect int
	gain   int
	index  int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].expect < pq[j].expect || pq[i].expect == pq[j].expect && pq[i].gain > pq[j].gain
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
