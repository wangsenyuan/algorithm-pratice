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
		n, m := readTwoNums(reader)
		conditions := make([][]int, m)
		for i := 0; i < m; i++ {
			conditions[i] = readNNums(reader, 2)
		}
		res := solve(n, conditions)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for i := 0; i < n; i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
		}
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, conditions [][]int) []int {

	up := make([]int, n+1)
	for i := 1; i <= n; i++ {
		up[i] = n
	}
	for _, cond := range conditions {
		x, y := cond[0], cond[1]
		up[x] = min(up[x], y)
	}

	arrs := make([][]int, n)

	for i := 0; i < n; i++ {
		arrs[i] = make([]int, 0, 1)
	}

	for i := 1; i <= n; i++ {
		y := up[i]
		arrs[y-1] = append(arrs[y-1], i)
	}

	pq := make(IntHeap, 0, n)

	res := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		for _, x := range arrs[i] {
			heap.Push(&pq, x)
		}
		if len(pq) == 0 {
			return nil
		}
		x := heap.Pop(&pq).(int)
		res[i] = x
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type IntHeap []int

func (this IntHeap) Len() int {
	return len(this)
}

func (this IntHeap) Less(i, j int) bool {
	return this[i] > this[j]
}

func (this IntHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *IntHeap) Push(x interface{}) {
	*this = append(*this, x.(int))
}

func (this *IntHeap) Pop() interface{} {
	old := *this
	n := len(old)
	res := old[n-1]
	*this = old[:n-1]
	return res
}
