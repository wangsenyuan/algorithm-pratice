package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		books := make([][]int, n)
		for i := 0; i < n; i++ {
			books[i] = readNNums(reader, 3)
		}

		fmt.Println(solve(n, k, books))
	}
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

func solve(n, k int, books [][]int) int64 {
	times := make([][]int, 4)
	sums := make([][]int64, 4)

	for i := 0; i < 4; i++ {
		times[i] = make([]int, 0, n)
		sums[i] = make([]int64, 0, n)
	}

	for _, book := range books {
		t, a, b := book[0], book[1], book[2]
		times[2*a+b] = append(times[2*a+b], t)
	}

	for i := 0; i < 4; i++ {
		sort.Ints(times[i])
		var cur int64
		sums[i] = append(sums[i], cur)
		for j := 0; j < len(times[i]); j++ {
			cur += int64(times[i][j])
			sums[i] = append(sums[i], cur)
		}
	}

	var best int64 = math.MaxInt64

	for cnt := 0; cnt < min(k+1, len(sums[3])); cnt++ {
		if k-cnt < len(sums[1]) && k-cnt < len(sums[2]) {
			tmp := sums[3][cnt] + sums[1][k-cnt] + sums[2][k-cnt]
			if tmp < best {
				best = tmp
			}
		}
	}

	if best == math.MaxInt64 {
		return -1
	}
	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve1(n, k int, books [][]int) int64 {

	sort.Sort(Books(books))

	var i int
	var k1, k2 int

	alice := make(IntHeap, 0, n)
	bob := make(IntHeap, 0, n)

	var tot int64
	for i < n && (k1 < k || k2 < k) {
		book := books[i]
		t, a, b := book[0], book[1], book[2]
		if a != 0 || b != 0 {
			tot += int64(t)
			if a == 1 {
				k1++
			}
			if b == 1 {
				k2++
			}
			if a == 0 {
				//only bob likes
				heap.Push(&bob, t)
			}
			if b == 0 {
				heap.Push(&alice, t)
			}
		}
		i++
	}

	for i < n && len(alice) > 0 && len(bob) > 0 {
		book := books[i]
		t, a, b := book[0], book[1], book[2]
		if a == 1 && b == 1 {
			tmp := t - alice[0] - bob[0]

			if tmp < 0 {
				//good to replace it, for bob & alice
				heap.Pop(&alice)
				heap.Pop(&bob)
				tot += int64(tmp)
			}
		}

		i++
	}

	if k1 < k || k2 < k {
		return -1
	}

	for alice.Len() > 0 && k1 > k {
		t := heap.Pop(&alice).(int)
		k1--
		tot -= int64(t)
	}

	for bob.Len() > 0 && k2 > k {
		t := heap.Pop(&bob).(int)
		k2--
		tot -= int64(t)
	}

	return tot
}

type Books [][]int

func (books Books) Len() int {
	return len(books)
}

func (books Books) Less(i, j int) bool {
	return books[i][0] < books[j][0]
}

func (books Books) Swap(i, j int) {
	books[i], books[j] = books[j], books[i]
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
