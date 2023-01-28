package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	S := readString(reader)
	T := readString(reader)
	res := solve(S, T)
	fmt.Println(res)
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

const MOD = 1e9 + 7

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func solve(S, T string) int {
	// len(S) <= 5000, len(T) <= 5000
	// count of s[i...j] of subseq of T
	// f(i, j) = count of occurrence of s[i..j] as sub seq in T
	// f(i, j + 1 ) = f(i, j) + s[j + 1] occurrence in T after s[i...j] position in T
	// f(i, j) = result until S[i] and T[j]
	// s[i+1], 找到所有s[i]在T中出现的位置，j, f(i+1, j) += f(i, j - 1) + 1 (1 for a new start of s[i])
	n := len(S)
	m := len(T)
	// dp[i][j] 表示以s[i]结尾的子串在T[..j]中的出现次数
	// 如果 s[i] = t[j], 则 dp[i][j] = dp[i-1][j-1] + 1
	dp := make([]int, n+1)

	for j := 0; j < m; j++ {
		for i := n; i > 0; i-- {
			if S[i-1] == T[j] {
				dp[i] = modAdd(dp[i], dp[i-1])
				dp[i] = modAdd(dp[i], 1)
			}
		}
	}

	var res int

	for i := 1; i <= n; i++ {
		res = modAdd(res, dp[i])
	}

	return res
}
