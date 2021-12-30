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
		n, a, b := readThreeNums(reader)
		I := make([][]int64, n)
		for i := 0; i < n; i++ {
			var l, r uint64
			s, _ := reader.ReadBytes('\n')
			pos := readUint64(s, 0, &l)
			readUint64(s, pos+1, &r)
			I[i] = []int64{int64(l), int64(r)}
		}
		res := solve(n, a, b, I)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Println(buf.String())
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

const INF = 1000000010

func solve(n int, a, b int, I [][]int64) int {
	A := int64(a)
	B := int64(b)

	P := make([]Pair, 0, 2*n+1)

	for _, cur := range I {
		P = append(P, Pair{cur[0], 0})
		P = append(P, Pair{cur[1], 1})
	}
	P = append(P, Pair{1e18, 1})

	sort.Sort(Pairs(P))
	var joy int
	var best int

	var prev int64 = -1
	for _, p := range P {

		y := p.first
		x := search(B-A+1+1, func(x int64) bool {
			return A*x+x*(x-1)/2 > y
		})
		x--
		R := B*x - x*(x-1)/2
		if R >= prev && best < joy {
			best = joy
		}
		if p.second == 0 {
			joy++
		} else {
			joy--
		}
		prev = p.first
	}
	return best
}

func search(n int64, fn func(int64) bool) int64 {
	var l, r int64 = 0, n
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

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type Pair struct {
	first  int64
	second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first || ps[i].first == ps[j].first && ps[i].second < ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
