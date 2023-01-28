package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k, A := readThreeNums(reader)
	S := make([][]int, n)

	for i := 0; i < n; i++ {
		S[i] = readNNums(reader, 2)
	}
	res := solve(n, k, A, S)
	fmt.Printf("%.9f\n", res)
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func solve(n int, k int, A int, senators [][]int) float64 {
	// 计算通过的可能性
	// 假设有两个人，每个人支持的可能性都是0.5, 那么至少有一个人通过的期望值是?
	// 两个人都通过 0.5 * 0.5, 一个人通过 0.5 * 0.5 + 0.5 * 0.5
	// 1 - (1 - p1) * (1 - p2)
	// 假设有n个人，至少m个人通过的期望值是？
	// dp[i][j] = dp[i-1][j-1] * pi + dp[i-1][j] * (1 - pi)

	X := make([]int, n)
	Y := make([]int, n)

	for i := 0; i < n; i++ {
		X[i] = senators[i][0]
		Y[i] = senators[i][1] / 10
	}

	var dfs1 func(i int, c int)

	var dfs2 func(i int, d int, s int, p float64) float64

	dfs2 = func(i int, d int, s int, p float64) float64 {
		if i == n {
			if d*2 > n {
				return p
			}
			return p * float64(A) / float64(A+s)
		}

		var res float64

		if Y[i] > 0 {
			res += dfs2(i+1, d+1, s, p*0.1*float64(Y[i]))
		}
		if Y[i] < 10 {
			res += dfs2(i+1, d, s+X[i], p*0.1*float64(10-Y[i]))
		}
		return res
	}

	var res float64

	dfs1 = func(i int, c int) {
		if i == n {
			res = math.Max(res, dfs2(0, 0, 0, 1))
			return
		}

		for x := 0; x <= c && Y[i]+x <= 10; x++ {
			Y[i] += x
			dfs1(i+1, c-x)
			Y[i] -= x
		}
	}

	dfs1(0, k)

	return res
}

type Pair struct {
	first  int
	second int
}
