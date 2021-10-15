package main

import (
	"bufio"
	"bytes"
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
		n, q := readTwoNums(reader)
		C := make([][]int, n)
		for i := 0; i < n; i++ {
			C[i] = readNNums(reader, 3)
		}
		Q := make([]int, q)
		for i := 0; i < q; i++ {
			Q[i] = readNum(reader)
		}
		res := solve(n, C, Q)

		for i := 0; i < q; i++ {
			buf.WriteString(fmt.Sprintf("%d\n", res[i]))
		}
	}

	fmt.Print(buf.String())
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

func solve(N int, C [][]int, Q []int) []int {
	items := make([]Item, N)
	var n0 int

	for i := 0; i < N; i++ {
		items[i] = Item{C[i][0], C[i][1], C[i][2]}
		if items[i].hasNoMetals() {
			n0++
		}
	}
	sort.Sort(Items(items[1:]))

	first := items[0]

	num := make([]int, N)
	var steps int
	for i := 1; i < N; i++ {
		if first.Less(items[i]) {
			steps += up(&first, &items[i])
			if first.Less(items[i]) {
				steps += down(&first, &items[i])
			}
		}
		num[i] = steps
	}

	ans := make([]int, len(Q))

	for i, k := range Q {
		if first.hasNoMetals() {
			ans[i] = N - n0 + 1
			continue
		}

		j := search(N-1, func(i int) bool {
			return num[i+1] > k
		})
		j++
		if j == N {
			// can get to rank 1
			ans[i] = 1
			continue
		}

		p1 := items[0]
		p2 := items[j]
		k -= num[j-1]
		upSteps(&p1, num[j-1])
		downSteps(&p2, k)
		if p2.Less(p1) || p1.Eq(p2) {
			ans[i] = N - j
		} else {
			ans[i] = N - j + 1
		}
	}
	return ans
}

func search(n int, fn func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if fn(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func upSteps(p *Item, n int) {
	tmp := min(n, p.bronze)
	p.gold += tmp
	p.bronze -= tmp
	n -= tmp
	tmp = min(n, p.silver)
	p.gold += tmp
	p.silver -= tmp
}

func downSteps(p *Item, n int) {
	tmp := min(p.gold, n)
	p.gold -= tmp
	p.bronze += tmp
	n -= tmp
	tmp = min(p.silver, n)
	p.silver -= tmp
	p.bronze += tmp
}

func up(p1, p2 *Item) int {
	if p1.Eq(*p2) || p2.Less(*p1) {
		return 0
	}
	// p2.gold > p1.gold
	stp := min(p1.bronze, p2.gold-p1.gold)
	p1.bronze -= stp
	p1.gold += stp
	tmp := min(p1.silver, p2.gold-p1.gold)
	p1.silver -= tmp
	p1.gold += tmp
	stp += tmp
	if p1.Less(*p2) {
		if p1.bronze > 0 {
			p1.bronze--
			p1.gold++
		} else if p1.silver > 0 {
			p1.silver--
			p1.gold++
		} else {
			stp--
		}
		stp++
	}
	return stp
}

func down(p1, p2 *Item) int {
	if p1.Eq(*p2) {
		return 0
	}
	return p2.gold - p1.gold + 1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Item struct {
	gold, silver, bronze int
}

func (this Item) Eq(that Item) bool {
	return this.gold == that.gold && this.silver == that.silver && this.bronze == that.bronze
}

func (this Item) Less(that Item) bool {
	if this.gold != that.gold {
		return this.gold < that.gold
	}

	if this.silver != that.silver {
		return this.silver < that.silver
	}

	if this.bronze != that.bronze {
		return this.bronze < that.bronze
	}

	return false
}

func (this Item) hasNoMetals() bool {
	return this.gold == 0 && this.silver == 0 && this.bronze == 0
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	return items[i].Less(items[j])
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}
