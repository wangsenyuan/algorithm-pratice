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
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(n, k, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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
		if s[i] == '\n' {
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

func solve(n int, k int, A []int) int64 {
	// at most n * (n + 1) / 2

	// if a array all equals (including 1), then it substracts l * (l + 1) / 2
	// so we need to make sure l as small as possible
	var ans int64
	var pq PQ

	for i := 0; i < n; {
		j := i
		for i < n && A[i] == A[j] {
			i++
		}
		cnt := int64(i - j)
		ans += cnt * (cnt + 1) / 2
		node := new(Node)
		node.m = cnt
		node.k = 1
		node.priority = cost(cnt, 1) - cost(cnt, 0)
		heap.Push(&pq, node)
	}

	for pq.Len() > 0 && k > 0 {
		p := heap.Pop(&pq).(*Node)
		ans += p.priority
		if p.m >= p.k+1 {
			p.priority = cost(p.m, p.k+1) - cost(p.m, p.k)
			p.k++
			heap.Push(&pq, p)
		}
		k--
	}

	return ans
}

func cost(X int64, P int64) int64 {
	rem := (X - P) % (P + 1)
	val := (X - P) / (P + 1)
	return P + (P+1-rem)*val*(val+1)/2 + rem*(val+1)*(val+2)/2
}

type Node struct {
	m        int64
	k        int64
	priority int64
}

type PQ []*Node

func (this PQ) Len() int {
	return len(this)
}

func (this PQ) Less(i, j int) bool {
	return this[i].priority < this[j].priority
}

func (this PQ) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *PQ) Push(x interface{}) {
	i := x.(*Node)
	*this = append(*this, i)
}

func (this *PQ) Pop() interface{} {
	old := *this
	n := len(old)
	ret := old[n-1]
	*this = old[:n-1]
	return ret
}
