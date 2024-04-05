package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, p, s := readThreeNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	gain, first, second := solve(a, b, p, s)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", gain))
	for i := 0; i < len(first); i++ {
		buf.WriteString(fmt.Sprintf("%d ", first[i]))
	}
	buf.WriteByte('\n')
	for i := 0; i < len(second); i++ {
		buf.WriteString(fmt.Sprintf("%d ", second[i]))
	}
	buf.WriteByte('\n')

	fmt.Print(buf.String())
}

const inf = 1 << 60

func solve(a []int, b []int, p int, s int) (int, []int, []int) {
	n := len(a)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		x := a[id[i]] - b[id[i]]
		y := a[id[j]] - b[id[j]]
		return x > y
	})

	prev := make(IntHeap, 0, p+1)

	dp := make([]int, n)
	var sum int
	for j, i := range id {
		sum += a[i]
		heap.Push(&prev, a[i])
		if prev.Len() > p {
			sum -= heap.Pop(&prev).(int)
		}
		dp[j] = -inf
		if prev.Len() == p {
			dp[j] = sum
		}
	}

	suf := make(IntHeap, 0, s+1)

	sum = 0

	var best int
	var pos int

	for i := n - 1; i >= 0; i-- {
		sum += b[id[i]]
		heap.Push(&suf, b[id[i]])
		if suf.Len() >= s {
			if suf.Len() > s {
				sum -= heap.Pop(&suf).(int)
			}
			if i >= p && best < dp[i-1]+sum {
				best = dp[i-1] + sum
				pos = i
			}
		}
	}

	// 然后选出前面最大的a[i]
	sort.Slice(id[:pos], func(i, j int) bool {
		return a[id[i]] > a[id[j]]
	})
	first := id[:p]
	sort.Slice(id[pos:], func(i, j int) bool {
		return b[id[i+pos]] < b[id[j+pos]]
	})

	second := id[n-s:]

	for i := 0; i < p; i++ {
		first[i]++
	}

	for i := 0; i < s; i++ {
		second[i]++
	}

	return best, first, second
}
func solve1(a []int, b []int, p int, s int) (int, []int, []int) {
	n := len(a)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		x := a[id[i]] - b[id[i]]
		y := a[id[j]] - b[id[j]]
		return x > y
	})

	//前面选p个，后面选s个
	dp := make([][]int, n+1)
	dp[0] = make([]int, p+1)
	for j := 1; j <= p; j++ {
		dp[0][j] = -inf
	}
	for i, v := range id {
		dp[i+1] = make([]int, p+1)
		copy(dp[i+1], dp[i])

		for j := min(i+1, p); j > 0; j-- {
			dp[i+1][j] = max(dp[i+1][j], dp[i][j-1]+a[v])
		}
	}

	fp := make([][]int, n+1)
	fp[n] = make([]int, s+1)
	for j := 1; j <= s; j++ {
		fp[n][j] = -inf
	}

	for i := n - 1; i >= 0; i-- {
		fp[i] = make([]int, s+1)
		copy(fp[i], fp[i+1])
		v := id[i]
		for j := min(n-i, s); j > 0; j-- {
			fp[i][j] = max(fp[i][j], fp[i+1][j-1]+b[v])
		}
	}

	best := -1
	for i := p; i+s <= n; i++ {
		if best < 0 || dp[best][p]+fp[best][s] < dp[i][p]+fp[i][s] {
			best = i
		}
	}

	// 然后选出前面最大的a[i]
	sort.Slice(id[:best], func(i, j int) bool {
		return a[id[i]] > a[id[j]]
	})

	first := id[:p]
	sort.Slice(id[best:], func(i, j int) bool {
		return b[id[i+best]] < b[id[j+best]]
	})

	second := id[n-s:]

	for i := 0; i < p; i++ {
		first[i]++
	}

	for i := 0; i < s; i++ {
		second[i]++
	}

	return dp[best][p] + fp[best][s], first, second
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
