package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	n := readNum(reader)
	Q := make([][]int, n)
	for i := 0; i < n; i++ {
		Q[i] = readNNums(reader, 2)
	}
	res := solve(Q)
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

const N = 200005

func solve(Q [][]int) []int64 {
	ask := make([][]Pair, N)

	add := func(i, l, r int) {
		if len(ask[l]) == 0 {
			ask[l] = make([]Pair, 0, 1)
		}
		ask[l] = append(ask[l], Pair{r, i})
	}
	ans := make([]int64, len(Q))
	for i, cur := range Q {
		l, r := cur[0], cur[1]
		add(i, l, r)
		ln := int64(r - l + 1)
		ans[i] = ln * (ln - 1) * (ln - 2) / 6
		ans[i] -= int64(max(r/6-(l-1)/3, 0))
		ans[i] -= int64(max(r/15-(l-1)/6, 0))
	}

	bit := NewBIT(N)

	cnt := make([]int64, N)

	for l := N - 1; l >= 1; l-- {
		for j := l << 1; j < N; j += l {
			bit.Add(j, cnt[j])
			cnt[j]++
		}
		for j := 0; j < len(ask[l]); j++ {
			r, i := ask[l][j].first, ask[l][j].second
			ans[i] -= bit.QueryRange(l, r)
		}
	}

	return ans
}

type BIT struct {
	arr []int64
	n   int
}

func NewBIT(n int) *BIT {
	arr := make([]int64, n+1)
	return &BIT{arr, n}
}

func (bit *BIT) Add(p int, v int64) {
	for p <= bit.n {
		bit.arr[p] += v
		p += p & -p
	}
}

func (bit *BIT) Query(p int) int64 {
	var res int64
	for p > 0 {
		res += bit.arr[p]
		p -= p & -p
	}
	return res
}

func (bit *BIT) QueryRange(l, r int) int64 {
	return bit.Query(r) - bit.Query(l-1)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}
