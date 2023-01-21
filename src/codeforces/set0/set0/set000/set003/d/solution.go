package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')

	var cnt int
	for i := 0; i < len(s); i++ {
		if s[i] == '?' {
			cnt++
		}
	}
	cost := make([][]int, cnt)
	for i := 0; i < cnt; i++ {
		cost[i] = readNNums(reader, 2)
	}
	res, ss := solve(s, cost)
	fmt.Println(res)
	if res < 0 {
		return
	}

	fmt.Println(ss)
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func solve(s string, cost [][]int) (int64, string) {
	ss := []byte(s)
	var open int
	pq := make(PriorityQueue, 0, len(s))
	var res int64
	for i, j := 0, 0; i < len(s); i++ {
		if s[i] != '(' && s[i] != ')' && s[i] != '?' {
			break
		}
		if s[i] == '(' {
			open++
		} else {
			open--
		}

		if s[i] == '?' {
			cur := cost[j]
			j++
			a, b := cur[0], cur[1]
			ss[i] = ')'
			heap.Push(&pq, &Item{b - a, i, -1})
			res += int64(b)
		}

		if open < 0 {
			if len(pq) == 0 {
				break
			}
			x := heap.Pop(&pq).(*Item)
			ss[x.second] = '('
			res += int64(-x.first)
			open += 2
		}
	}

	if open != 0 {
		return -1, ""
	}
	return res, string(ss)
}

type Item struct {
	first  int
	second int
	index  int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].first > pq[j].first
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
