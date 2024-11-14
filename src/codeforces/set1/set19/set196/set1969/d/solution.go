package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		a := readNNums(reader, n)
		b := readNNums(reader, n)
		res := solve(k, a, b)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(k int, a []int, b []int) int {
	type pair struct {
		first  int
		second int
	}

	n := len(a)

	arr := make([]pair, n)

	for i := 0; i < n; i++ {
		arr[i] = pair{b[i], a[i]}
	}
	// 按照b降序排列
	slices.SortFunc(arr, func(x, y pair) int {
		return y.first - x.first
	})
	// suf[i] 表示在i...n的最大收益
	suf := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] + max(arr[i].first-arr[i].second, 0)
	}
	if k == 0 {
		return suf[0]
	}
	// 前缀的最小cost
	cost := make(IntHeap, 0, k+1)
	var sum int
	var res int
	for i := 0; i < n; i++ {
		heap.Push(&cost, arr[i].second)
		sum += arr[i].second
		if cost.Len() > k {
			// 把最贵的舍弃掉
			sum -= heap.Pop(&cost).(int)
		}

		if cost.Len() == k {
			res = max(res, suf[i+1]-sum)
		}
	}

	return res
}

func solve1(k int, a []int, b []int) int {
	n := len(a)
	arr := make([]Pair, n)
	var profit int
	for i := 0; i < n; i++ {
		arr[i] = Pair{a[i], b[i]}
		profit += max(b[i]-a[i], 0)
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].second > arr[j].second
	})

	var best int

	if k == 0 {
		best = profit
	}

	var f int

	pq := make(IntHeap, 0, n)

	for _, cur := range arr {
		profit -= max(0, cur.second-cur.first)
		heap.Push(&pq, cur.first)
		f += cur.first
		if pq.Len() > k {
			f -= heap.Pop(&pq).(int)
		}
		if pq.Len() == k {
			best = max(profit-f, best)
		}
	}

	return best

}

type Pair struct {
	first  int
	second int
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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
