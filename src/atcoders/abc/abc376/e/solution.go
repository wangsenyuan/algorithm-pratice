package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for tc > 0 {
		tc--
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) int64 {
	n, k := readTwoNums(reader)

	a := readNNums(reader, n)
	b := readNNums(reader, n)
	return solve(k, a, b)
}

type pair struct {
	first  int
	second int
}

func solve(k int, a []int, b []int) int64 {
	n := len(a)
	arr := make([]pair, n)

	for i := 0; i < n; i++ {
		arr[i] = pair{a[i], b[i]}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].first < arr[j].first
	})

	pq := make(IntHeap, 0, n)
	var ans int64 = math.MaxInt64
	var sum int64
	for i := 0; i < n; i++ {
		if pq.Len() == k {
			sum -= int64(heap.Pop(&pq).(int))
		}
		// pq.Len() <= k - 1
		if pq.Len() == k-1 {
			x := int64(arr[i].first)
			y := int64(arr[i].second) + sum
			ans = min(ans, x*y)
		}
		heap.Push(&pq, arr[i].second)
		sum += int64(arr[i].second)
	}

	return ans
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
