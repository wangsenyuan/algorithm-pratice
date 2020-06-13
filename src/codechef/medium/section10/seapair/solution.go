package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)

		X, D := make([]int, n), make([]int, n)

		for i := 0; i < n; i++ {
			X[i], D[i] = readTwoNums(reader)
		}

		fmt.Println(solve(n, k, X, D))
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

const MOD = 1000000007

var TT int

func init() {
	TT = pow(10, 100) + 1
	if TT >= MOD {
		TT -= MOD
	}
}

func pow(a, b int) int {
	A := int64(a)
	R := int64(1)

	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(R)
}

func solve(n, k int, X []int, D []int) int {
	k = n - k
	if k == 0 {
		return TT
	}

	events := make([]Pair, 2*n)

	for i := 0; i < n; i++ {
		a := int64(X[i]) * int64(D[i])
		b := max(0, a+int64(X[i]-1))
		events[2*i] = Pair{a, 0}
		events[2*i+1] = Pair{b, 1}
	}

	sort.Sort(Pairs(events))

	var cnt int
	var prev int64
	var res int64
	for i := 0; i < len(events); i++ {
		evt := events[i]
		if evt.second == 0 {
			cnt++
			if cnt == k {
				prev = evt.first
			}
		} else {
			cnt--
			if cnt == k-1 {
				res += evt.first - prev + 1
			}
		}
	}
	return int(res % MOD)
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

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
