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
	A := readNNums(reader, n)

	queries := make([][]int, m)

	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 3)
	}

	res := solve(A, queries)

	var buf bytes.Buffer

	for _, vs := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", vs[0], vs[1]))
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

const mod = 180

func modAdd(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	if a < 0 {
		a += mod
	}
	return a
}

func solve(X []int, queries [][]int) [][]int {
	for i := 0; i < len(X); i++ {
		X[i] = ((X[i])%mod + mod) % mod
	}
	var res [][]int

	last := make([]int, mod)
	rst := make([]int, mod)
	for i := 0; i < mod; i++ {
		rst[i] = -1
	}
	for _, cur := range queries {
		if cur[0] == 1 {
			i := cur[1]
			x := cur[2]
			X[i-1] = (x%mod + mod) % mod
		} else {
			l, r := cur[1], cur[2]
			l--
			r--
			r = min(r, l+mod)
			copy(last, rst)
			last[0] = l
			var found bool
			var sum int
			for i := l; i <= r && !found; i++ {
				sum = modAdd(sum, X[i])
				if last[sum] >= 0 {
					found = true
					res = append(res, []int{last[sum], i})
				}
				last[sum] = i + 1
			}
			for v := 1; v <= mod/2 && !found; v++ {
				for x := 0; x < mod && !found; x++ {
					if last[x] >= 0 {
						y := modAdd(x, v)
						if last[y] >= 0 {
							found = true
							res = append(res, []int{min(last[x], last[y]), max(last[x], last[y]) - 1})
						}
					}
				}
			}
		}
	}

	for _, vs := range res {
		for i := 0; i < len(vs); i++ {
			vs[i]++
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
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
