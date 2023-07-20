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
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
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
	if len(bs) == 0 || bs[0] == '\n' {
		return readNum(reader)
	}
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

func solve(A []int) int {
	// dp[i][rem] = 到i为止，以i结尾的连续2个数，且sum为rem时的最小值
	// 处理i+1时，需要保证 (rem + A[i] + x) % 3 = 0
	// 同时，为了后续的处理，还必须知道在i处添加的数是多少
	// dp[i][a][b] = 到i为止时，在i处+a时，且(A[i-1] + ? + A[i] + a) % 3 = b时的最小值
	// dp[i+1][x][y] 就能够计算出来了

	dp := make([][][]int, 2)

	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, 3)
		for j := 0; j < 3; j++ {
			dp[i][j] = make([]int, 3)
			for k := 0; k < 3; k++ {
				dp[i][j][k] = inf
			}
		}
	}

	// for i = 1
	for a := 0; a < 3; a++ {
		for na := 0; na < 3; na++ {
			nb := (a + A[0] + A[1] + na) % 3
			dp[1][na][nb] = a + na
		}
	}

	n := len(A)
	for i := 2; i < n; i++ {
		it := i & 1
		for a := 0; a < 3; a++ {
			for b := 0; b < 3; b++ {
				dp[it][a][b] = inf
			}
		}
		for a := 0; a < 3; a++ {
			for b := 0; b < 3; b++ {
				for na := 0; na < 3; na++ {
					if (na+b+A[i])%3 == 0 {
						nb := (na + a + A[i-1] + A[i]) % 3
						dp[it][na][nb] = min(dp[it][na][nb], dp[it^1][a][b]+na)
					}
				}
			}
		}
	}
	res := inf

	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			res = min(res, dp[(n-1)&1][a][b])
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
