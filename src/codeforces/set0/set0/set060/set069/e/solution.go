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
	n, k := readTwoNums(reader)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = readNum(reader)
	}
	res := solve(a, k)
	var buf bytes.Buffer

	for _, x := range res {
		if x < -inf {
			buf.WriteString("Nothing\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", x))
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

const inf = 1000000000

func solve(a []int, k int) []int {
	n := len(a)

	items := make(map[int]*Item)

	pq := make(Heap, 0, k)

	add := func(x int) {
		if _, ok := items[x]; !ok {
			items[x] = &Item{x, -1, 0}
		}
		it := items[x]
		it.freq++
		if it.freq == 1 {
			heap.Push(&pq, it)
		} else if it.freq == 2 {
			it.val = inf + 1
			heap.Fix(&pq, it.ind)
			heap.Pop(&pq)
			it.val = x
		}
		// else just ignore it
	}

	remove := func(x int) {
		it := items[x]
		it.freq--
		if it.freq == 1 {
			// 它原来 = 2, 不在heap中，现在要放进去
			heap.Push(&pq, it)
		} else if it.freq == 0 {
			// 它原来在heap中，现在要移除
			it.val = inf + 1
			heap.Fix(&pq, it.ind)
			heap.Pop(&pq)
			it.val = x
		}
	}
	var res []int

	for i := 0; i < n; i++ {
		add(a[i])
		if i >= k {
			remove(a[i-k])
		}
		if i >= k-1 {
			if len(pq) > 0 {
				res = append(res, pq[0].val)
			} else {
				res = append(res, -(inf + 1))
			}
		}
	}
	return res
}

type Item struct {
	val  int
	ind  int
	freq int
}

// An IntHeap is a min-heap of ints.
type Heap []*Item

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].ind = i
	h[j].ind = j
}

func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	it := x.(*Item)
	it.ind = len(*h)
	*h = append(*h, it)
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	x.ind = -1
	return x
}
