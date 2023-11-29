package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		if len(res) == 0 {
			buf.WriteString("No\n")
			continue
		}
		buf.WriteString("Yes\n")
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

func solve(a []int) []int {
	sort.Ints(a)
	n := len(a)

	diff := a[n-1] - a[0]
	if diff == 0 {
		// no solution
		return nil
	}
	var res []int

	pos := make(IntHeap, 0, n)
	neg := make(IntHeap, 0, n)
	for i := 0; i < n; i++ {
		if abs(a[i]) >= diff {
			return nil
		}
		if a[i] == 0 {
			res = append(res, 0)
			continue
		}
		if a[i] > 0 {
			heap.Push(&pos, a[i])
		} else {
			heap.Push(&neg, a[i])
		}
	}
	var sum int
	for len(pos) > 0 || len(neg) > 0 {
		for sum <= 0 && len(pos) > 0 {
			// a pos sure exists
			sum += pos[0]
			res = append(res, pos[0])
			heap.Pop(&pos)
		}
		if abs(sum) >= diff {
			return nil
		}
		for sum > 0 && len(neg) > 0 {
			sum += neg[0]
			res = append(res, neg[0])
			heap.Pop(&neg)
		}

		if abs(sum) >= diff {
			return nil
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

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
