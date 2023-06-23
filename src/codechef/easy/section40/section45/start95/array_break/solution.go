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
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		S := readString(reader)[:m]
		res := solve(A, S)
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

const MOD = 1_000_000_007

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

func inverse(a int) int {
	return pow(a, MOD-2)
}

func sub(a, b int) int {
	return add(a, MOD-b)
}

func div(a, b int) int {
	return mul(a, inverse(b))
}

func solve(A []int, S string) int {
	// dp[i][j][k] 表示在第k次操作后，A是数组i...j的概率
	// dp[i][a][k+1] 如果 S[k+1] = 'L' = (a - i + 1) / (b - a + 1) * dp[i][j][k]
	n := len(A)
	k := len(S)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][n-1] = 1

	ndp := make([][]int, n)

	for i := 0; i < n; i++ {
		ndp[i] = make([]int, n+1)
	}

	inv := make([]int, n+1)
	for i := 0; i <= n; i++ {
		inv[i] = inverse(i)
	}

	for i := 0; i < k; i++ {
		for l := 0; l < n; l++ {
			for r := l; r < n; r++ {
				if l == r {
					ndp[l][r] = add(ndp[l][r], dp[l][r])
					if S[i] == 'L' {
						ndp[l][r+1] = sub(ndp[l][r+1], dp[l][r])
					} else if l > 0 {
						ndp[l-1][r] = sub(ndp[l-1][r], dp[l][r])
					}
				} else {
					val := mul(dp[l][r], inv[r-l])
					if S[i] == 'L' {
						ndp[l][l] = add(ndp[l][l], val)
					} else {
						ndp[r][r] = add(ndp[r][r], val)
					}
					ndp[l][r] = sub(ndp[l][r], val)
				}
			}
		}
		if S[i] == 'L' {
			for l := 0; l < n; l++ {
				for r := 1; r < n; r++ {
					ndp[l][r] = add(ndp[l][r], ndp[l][r-1])
				}
			}
		} else {
			for r := 0; r < n; r++ {
				for l := n - 2; l >= 0; l-- {
					ndp[l][r] = add(ndp[l][r], ndp[l+1][r])
				}
			}
		}

		for l := 0; l < n; l++ {
			for r := 0; r <= n; r++ {
				dp[l][r] = ndp[l][r]
				ndp[l][r] = 0
			}
		}
	}
	var ans int

	for l := 0; l < n; l++ {
		var sum int
		for r := l; r < n; r++ {
			sum = add(sum, A[r])
			ans = add(ans, mul(sum, dp[l][r]))
		}
	}

	return ans
}
