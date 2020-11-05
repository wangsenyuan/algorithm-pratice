package main

import (
	"bufio"
	"bytes"
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
		E := make([][]int, n*(n-1)/2)
		for i := 0; i < len(E); i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E)
		if len(res) == 0 {
			buf.WriteString("-1\n")
			continue
		}
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))

		for i := 0; i < len(res); i++ {
			if i < len(res)-1 {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			} else {
				buf.WriteString(fmt.Sprintf("%d\n", res[i]))
			}
		}
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

func solve(n int, E [][]int) []int {
	if n&(n-1) != 0 {
		return nil
	}
	// has answer

	conn := make([][]int, n)

	for i := 0; i < n; i++ {
		conn[i] = make([]int, n)
	}

	for i := 0; i < len(E); i++ {
		u, v := E[i][0], E[i][1]
		u--
		v--
		conn[u][v] = i + 1
		conn[v][u] = -(i + 1)
	}
	use := make([]bool, len(E)+1)
	win := make([]int, n)
	for i := 0; i < n; i++ {
		win[i] = i
	}
	for num, d := n, 1; num > 1; num >>= 1 {
		for i := 0; i < n; i += 2 * d {
			a, b := win[i], win[i+d]
			if conn[a][b] < 0 {
				a, b = b, a
			}
			use[conn[a][b]] = true
			win[i] = a
		}
		d *= 2
	}

	res := make([]int, 0, len(E)-(n-1))
	for i := 1; i <= len(E); i++ {
		if !use[i] {
			res = append(res, i)
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

//
//// An Item is something we manage in a priority queue.
//type Item struct {
//	value    int // The value of the item; arbitrary.
//	priority int // The priority of the item in the queue.
//	// The index is needed by update and is maintained by the heap.Interface methods.
//	index int // The index of the item in the heap.
//}
//
//// A PriorityQueue implements heap.Interface and holds Items.
//type PriorityQueue []*Item
//
//func (pq PriorityQueue) Len() int { return len(pq) }
//
//func (pq PriorityQueue) Less(i, j int) bool {
//	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
//	return pq[i].priority < pq[j].priority
//}
//
//func (pq PriorityQueue) Swap(i, j int) {
//	pq[i], pq[j] = pq[j], pq[i]
//	pq[i].index = i
//	pq[j].index = j
//}
//
//func (pq *PriorityQueue) Push(x interface{}) {
//	n := len(*pq)
//	item := x.(*Item)
//	item.index = n
//	*pq = append(*pq, item)
//}
//
//func (pq *PriorityQueue) Pop() interface{} {
//	old := *pq
//	n := len(old)
//	item := old[n-1]
//	old[n-1] = nil  // avoid memory leak
//	item.index = -1 // for safety
//	*pq = old[0 : n-1]
//	return item
//}
