package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	B := readNNums(reader, n)
	m := readNum(reader)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 2)
	}
	res := solve(n, B, Q)

	var buf bytes.Buffer

	for i := 0; i < m; i++ {
		buf.WriteString(fmt.Sprintf("%.1f\n", res[i]))
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

func solve(n int, B []int, Q [][]int) []float64 {

	set := func(arr []int, p int, v int, fn func(int, int) int) {
		p += n
		arr[p] = v
		for p > 1 {
			arr[p>>1] = fn(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(arr []int, l, r int, init_res int, fn func(int, int) int) int {
		res := init_res
		l += n
		r += n
		for l < r {
			if l&1 == 1 {
				res = fn(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = fn(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	dp := make([]int, 2*n)
	fp := make([]int, 2*n)

	for i := 0; i < 2*n; i++ {
		dp[i] = INF
		fp[i] = -INF
	}

	for i := 0; i < n; i++ {
		set(dp, i, B[i], min)
		set(fp, i, B[i], max)
	}

	res := make([]float64, len(Q))

	for i, cur := range Q {
		l, r := cur[0], cur[1]
		x := get(dp, l, r+1, INF, min)
		// first match stick burns outs
		var ans float64
		if l > 0 {
			ans = float64(get(fp, 0, l, 0, max)) + float64(x)
		}
		if r+1 < n {
			ans = math.Max(ans, float64(get(fp, r+1, n, 0, max)+x))
		}
		y := get(fp, l, r+1, 0, max)
		ans = math.Max(ans, float64(y-x)/2+float64(x))
		res[i] = ans
	}
	return res
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
