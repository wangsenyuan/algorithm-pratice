package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	X := make([]int, n)
	W := make([]int, n)
	for i := 0; i < n; i++ {
		X[i], W[i] = readTwoNums(reader)
	}
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 2)
	}
	res := solve(X, W, Q)
	var buf bytes.Buffer

	for i := 0; i < m; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

const INF = 1 << 62

func solve(X []int, W []int, Q [][]int) []int64 {
	type Pair struct {
		first  int
		second int
	}
	n := len(X)
	queries := make([][]Pair, n+1)
	for i, cur := range Q {
		l, r := cur[0], cur[1]
		l--
		if len(queries[r]) == 0 {
			queries[r] = make([]Pair, 0, 1)
		}
		queries[r] = append(queries[r], Pair{l, i})
	}

	L := make([][]int, n+1)

	addEnd := func(r int, l int) {
		if L[r] == nil {
			L[r] = make([]int, 0, 1)
		}
		L[r] = append(L[r], l)
	}

	stack := make([]int, n)
	var p int
	for r := 0; r < n; r++ {
		for p > 0 && W[stack[p-1]] > W[r] {
			p--
		}
		if p > 0 {
			// W[x] < W[r] and x < r
			addEnd(r, stack[p-1])
		}
		stack[p] = r
		p++
	}

	p = 0

	for i := n - 1; i >= 0; i-- {
		for p > 0 && W[stack[p-1]] > W[i] {
			p--
		}
		if p > 0 {
			// W[x] < W[i] and x > i
			addEnd(stack[p-1], i)
		}
		stack[p] = i
		p++
	}

	arr := make([]int64, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = INF
	}

	update := func(i int, v int64) {
		for ; i > 0; i -= i & -i {
			arr[i] = min(arr[i], v)
		}
	}

	get := func(i int) int64 {
		var res int64 = INF
		for ; i <= n; i += i & -i {
			res = min(res, arr[i])
		}
		return res
	}
	ans := make([]int64, len(Q))
	for r := 1; r <= n; r++ {
		for _, l := range L[r-1] {
			val := int64(X[r-1]-X[l]) * int64(W[r-1]+W[l])
			update(l+1, val)
		}
		for _, cur := range queries[r] {
			ans[cur.second] = get(cur.first + 1)
		}
	}

	return ans
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
