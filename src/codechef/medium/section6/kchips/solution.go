package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	V := readNNums(reader, n)
	q := readNum(reader)
	R := make([][]int, q)
	for i := 0; i < q; i++ {
		R[i] = readNNums(reader, 2)
	}
	res := solve(k, V, R)

	var buf bytes.Buffer

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

func solve(K int, V []int, RR [][]int) []int {
	K++
	n := len(V)
	sv := make([]Pair, n)
	for i := 0; i < n; i++ {
		sv[i] = Pair{V[i], i + 1}
	}

	sort.Sort(Pairs(sv))

	vv := make([]int, n+1)
	cnt := make([]int, n+1)
	var p int
	for i := 0; i < n; i++ {
		if i > 0 && sv[i].first > sv[i-1].first {
			p++
		}
		cnt[p]++
		vv[sv[i].second] = p
	}

	ind := make([][]int, p+1)
	for i := 0; i <= p; i++ {
		ind[i] = make([]int, 0, cnt[i])
	}

	Q := make([]Triple, len(RR))

	for i, cur := range RR {
		l, r := cur[0], cur[1]
		Q[i] = Triple{r + 1, l + 1, i}
	}

	sort.Sort(Triples(Q))

	arr := make([]int, n+1)

	update := func(p, v int) {
		for p <= n {
			arr[p] += v
			p += p & -p
		}
	}

	get := func(p int) int {
		var res int
		for p > 0 {
			res += arr[p]
			p -= p & -p
		}
		return res
	}
	ans := make([]int, len(Q))
	var tot int
	for i, j := 1, 0; i <= n; i++ {
		x := vv[i]
		if len(ind[x]) >= K {
			update(ind[x][len(ind[x])-K], -1)
		}
		ind[x] = append(ind[x], i)
		if len(ind[x]) == K {
			tot++
		}
		if len(ind[x]) >= K {
			update(ind[x][len(ind[x])-K], 1)
		}
		for j < len(Q) && Q[j].first == i {
			ans[Q[j].third] = tot - get(Q[j].second-1)
			j++
		}
	}

	return ans
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first < this[j].first
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type Triple struct {
	first, second, third int
}

type Triples []Triple

func (this Triples) Len() int {
	return len(this)
}

func (this Triples) Less(i, j int) bool {
	return this[i].first < this[j].first
}

func (this Triples) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func solve1(K int, V []int, RR [][]int) []int {
	n := len(V)
	ii := make(map[int]int)
	for i, j := 0, 0; i < n; i++ {
		v := V[i]
		if _, found := ii[v]; !found {
			ii[v] = j
			j++
		}
	}
	for i := 0; i < n; i++ {
		V[i] = ii[V[i]]
	}

	rr := make([]Round, len(RR))

	blockSize := int(math.Sqrt(float64(len(RR))))

	for i := 0; i < len(RR); i++ {
		rr[i] = Round{blockSize, RR[i][0], RR[i][1] + 1, i}
	}

	sort.Sort(Rounds(rr))

	cnt := make([]int, len(ii))
	var res int

	add := func(v int) {
		if cnt[v] == K {
			res++
		}
		cnt[v]++
	}

	sub := func(v int) {
		cnt[v]--
		if cnt[v] == K {
			res--
		}
	}

	var L, R int
	ans := make([]int, len(RR))

	for _, cur := range rr {
		l, r := cur.l, cur.r
		for R < r {
			add(V[R])
			R++
		}
		for R > r {
			R--
			sub(V[R])
		}
		for L > l {
			L--
			add(V[L])
		}
		for L < l {
			sub(V[L])
			L++
		}
		ans[cur.idx] = res
	}

	return ans
}

const B = 200

type Round struct {
	blockSize int
	l, r      int
	idx       int
}

func (this Round) Less(that Round) bool {
	if this.r/this.blockSize < that.r/this.blockSize {
		return true
	}

	if this.r/this.blockSize == that.r/this.blockSize {
		return this.l < that.l
	}

	return false
}

type Rounds []Round

func (this Rounds) Len() int {
	return len(this)
}

func (this Rounds) Less(i, j int) bool {
	return this[i].Less(this[j])
}

func (this Rounds) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
