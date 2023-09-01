package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		m := readNum(reader)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(A, Q)
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

const inf = 1 << 30

func solve(A []int, Q [][]int) []int {
	n := len(A)

	arr := make([]Pair, n*2)

	for i := n; i < 2*n; i++ {
		arr[i] = Pair{i - n, A[i-n]}
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = min_pair(arr[i*2], arr[i*2+1])
	}

	set := func(p int, v Pair) {
		p += n
		arr[p] = v
		for p > 1 {
			arr[p>>1] = min_pair(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l int, r int) Pair {
		res := Pair{-1, inf}
		for i, j := l+n, r+n; i < j; i, j = i>>1, j>>1 {
			if i&1 == 1 {
				res = min_pair(res, arr[i])
				i++
			}
			if j&1 == 1 {
				j--
				res = min_pair(res, arr[j])
			}
		}
		return res
	}

	buf := make([]Pair, 31)

	ans := make([]int, len(Q))

	check := func(m int) int {
		best := inf - 1
		for i := 0; i < m; i++ {
			for j := i + 1; j < m; j++ {
				best = min(best, buf[i].second|buf[j].second)
			}
		}
		return best
	}

	for i, cur := range Q {
		l, r := cur[0], cur[1]
		l--
		var j int
		for j < len(buf) {
			x := get(l, r)
			if x.second == inf {
				break
			}
			buf[j] = x
			j++
			set(x.first, Pair{x.first, inf})
		}

		ans[i] = check(j)

		for j > 0 {
			x := buf[j-1]
			set(x.first, x)
			j--
		}
	}

	return ans
}

type Pair struct {
	first  int
	second int
}

func min_pair(a, b Pair) Pair {
	if a.second < b.second {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
