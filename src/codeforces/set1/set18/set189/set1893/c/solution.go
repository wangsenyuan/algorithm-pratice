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
		m := readNum(reader)
		L := make([]int, m)
		R := make([]int, m)
		a := make([][]int, m)
		c := make([][]int, m)
		var n int
		for i := 0; i < m; i++ {
			n, L[i], R[i] = readThreeNums(reader)
			a[i] = readNNums(reader, n)
			c[i] = readNNums(reader, n)
		}
		res := solve(L, R, a, c)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(L []int, R []int, a [][]int, c [][]int) int {
	sum := make([]int, 3)
	sumc := make([]int, len(L))
	suma := make(map[int]int)

	pos := make(map[int][]Pair)
	for i := 0; i < len(L); i++ {
		sum[0] += L[i]
		sum[1] += R[i]
		sum[2] += len(a[i])
		for j := 0; j < len(c[i]); j++ {
			suma[a[i][j]] += R[i]
			sumc[i] += c[i][j]
			if len(pos[a[i][j]]) == 0 {
				pos[a[i][j]] = make([]Pair, 0, 1)
			}
			pos[a[i][j]] = append(pos[a[i][j]], Pair{i, j})
		}
	}

	if sum[1]-sum[0] > sum[2] {
		return 0
	}
	var best = sum[1]

	process := func(sz int) {
		var must int
		x := sum[1] - suma[sz]
		for _, p := range pos[sz] {
			i, j := p.first, p.second
			other := sumc[i] - c[i][j]
			if other < L[i] {
				// 不够
				must += L[i] - other
				x += L[i]
			} else {
				x += min(other, R[i])
			}
		}
		tmp := must + max(0, sz-x)
		best = min(best, tmp)
	}

	for x := sum[0]; x <= sum[1]; x++ {
		process(x)
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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
