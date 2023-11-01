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

	t := readNum(reader)

	for t > 0 {
		t--
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		q := make([][]int, m)
		for i := 0; i < m; i++ {
			q[i] = readNNums(reader, 2)
		}
		res := solve(a, q)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		}
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

const N = 1_000_00*2 + 10

var lpf [N]int

func init() {
	var primes []int
	for x := 2; x < N; x++ {
		if lpf[x] == 0 {
			lpf[x] = x
			primes = append(primes, x)
		}
		for j := 0; j < len(primes); j++ {
			if primes[j]*x >= N {
				break
			}
			lpf[primes[j]*x] = primes[j]
			if x%primes[j] == 0 {
				break
			}
		}
	}
}

func solve(a []int, queries [][]int) []int {
	n := len(a)

	items := make(map[int][]*Item)

	pq := make(PQ, 0, n)

	// 假设n有两个影子 a, b 且 a % b = 0, 那么就没有必要处理b

	find := func(d int) {
		cur := make([]*Item, d)
		// d is not prime
		for i := 0; i < d; i++ {
			var sum int
			for j := i; j < n; j += d {
				sum += a[j]
			}
			it := new(Item)
			it.id = d*n + i
			it.value = sum * d
			cur[i] = it
			heap.Push(&pq, it)
		}

		items[d] = cur
	}

	var sum int
	for i := 0; i < n; i++ {
		sum += a[i]
	}

	for m := n; m > 1; {
		d := lpf[m]
		x := n / d
		if _, ok := items[x]; !ok && x > 1 {
			find(x)
		}
		m /= d
	}

	get := func() int {
		if len(pq) > 0 {
			return max(sum, pq[0].value)
		}
		return sum
	}

	ans := make([]int, 1+len(queries))
	ans[0] = get()

	for i, cur := range queries {
		p, x := cur[0]-1, cur[1]
		diff := x - a[p]

		ok := make(map[int]bool)

		for m := n; m > 1; {
			d := lpf[m]
			c := n / d
			r := p % c
			if r < len(items[c]) && !ok[items[c][r].id] {
				pq.update(items[c][r], items[c][r].value+diff*c)
				ok[items[c][r].id] = true
			}
			m /= d
		}
		sum += diff

		ans[i+1] = get()

		a[p] += diff
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Item struct {
	id    int
	value int
	index int
}

type PQ []*Item

func (h PQ) Len() int           { return len(h) }
func (h PQ) Less(i, j int) bool { return h[i].value > h[j].value }
func (h PQ) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *PQ) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	it := x.(*Item)
	it.index = len(*h)
	*h = append(*h, it)
}

func (h *PQ) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *PQ) update(it *Item, v int) {
	it.value = v
	heap.Fix(h, it.index)
}
