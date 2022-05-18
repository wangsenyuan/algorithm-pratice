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
		n, M := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(A, M)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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
		if s[i] == '\n' {
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

const INF = 1 << 30

func solve(A []int, M int) int {
	j := 0

	n := len(A)

	for i := 1; i < n; i++ {
		if A[j] < A[i] {
			j = i
		}
	}

	var best int

	for i := 0; i < n; i++ {
		best = max(best, A[i]+A[j]+(A[i]-A[j]+M)%M)
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func solve1(A []int, M int) int {
	n := len(A)

	P := make([]Pair, n)

	for i := 0; i < n; i++ {
		P[i] = Pair{A[i], A[i] % M}
	}

	sort.Sort(Pairs(P))

	var best int
	var p int

	for i := 0; i < n; {
		j := i
		for i < n && P[j].second == P[i].second {
			i++
		}
		best = max(best, P[j].first+P[j].first)
		P[p] = P[j]
		p++
	}
	var peek int

	for i := 0; i < p; i++ {
		if i > 0 {
			tmp := P[i].first + P[i].second + peek + M
			best = max(best, tmp)
		}
		peek = max(peek, P[i].first-P[i].second)
	}

	peek = 0

	for i := p - 1; i >= 0; i-- {
		if i < p-1 {
			tmp := P[i].first + P[i].second + peek
			best = max(best, tmp)
		}
		peek = max(peek, P[i].first-P[i].second)
	}

	return best
}

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].second > this[j].second || this[i].second == this[j].second && this[i].first > this[j].first
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
