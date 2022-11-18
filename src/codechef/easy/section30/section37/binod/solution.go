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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readNInt64s(reader, n)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 5)
		}
		res := solve(A, Q)
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d\n", ans))
		}
	}

	fmt.Print(buf.String())
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

const H = 60

func solve(A []int64, Q [][]int) []int64 {
	n := len(A)

	qs := make([]Query, len(Q))

	for i, cur := range Q {
		qs[i] = Query{i, cur[0], cur[1], cur[2], cur[3], cur[4]}
	}

	sort.Slice(qs, func(i, j int) bool {
		return qs[i].k < qs[j].k
	})
	ans := make([]int64, len(Q))
	cnt := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		cnt[i] = make([]int, 2)
	}

	query := func(q Query) int64 {
		a := int64(cnt[q.r1][0]-cnt[q.l1-1][0]) * int64(cnt[q.r2][1]-cnt[q.l2-1][1])
		b := int64(cnt[q.r1][1]-cnt[q.l1-1][1]) * int64(cnt[q.r2][0]-cnt[q.l2-1][0])
		return a + b
	}

	var j int
	for h := 0; h < H; h++ {

		for i := 0; i < n; i++ {
			copy(cnt[i+1], cnt[i])
			if (A[i]>>h)&1 == 1 {
				cnt[i+1][1]++
			} else {
				cnt[i+1][0]++
			}
		}

		for j < len(qs) && qs[j].k == h {
			ans[qs[j].id] = query(qs[j])
			j++
		}
	}

	return ans
}

type Query struct {
	id int
	k  int
	l1 int
	r1 int
	l2 int
	r2 int
}
