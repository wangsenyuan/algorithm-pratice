package main

import (
	"bufio"
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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, k := readTwoNums(scanner)
		segs := make([][]int, n)
		for i := 0; i < n; i++ {
			segs[i] = readNNums(scanner, 2)
		}
		fmt.Println(solve(n, k, segs))
	}
}

func solve(n int, k int, segs [][]int) int {
	ps := make(Pairs, n)

	for i := 0; i < n; i++ {
		ps[i] = Pair{segs[i][0], segs[i][1]}
	}
	sort.Sort(ps)
	pq := new(IntHeap)

	for i := 0; i < k-1; i++ {
		heap.Push(pq, ps[i].second)
	}
	var ans int

	for i := k - 1; i < n; i++ {
		heap.Push(pq, ps[i].second)
		ans = max(ans, heap.Pop(pq).(int)-ps[i].first)
	}

	return int(ans)
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

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first || (ps[i].first == ps[j].first && ps[i].second > ps[j].second)
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
