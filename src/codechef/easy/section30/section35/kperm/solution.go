package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(n, k, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const MOD = 1000000007

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
}

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func reverse(a int) int {
	return pow(a, MOD-2)
}

func solve(n int, K int, P []int) int {
	cnt := make([]int, 2)
	cnt[0] = n / 2
	cnt[1] = (n + 1) / 2
	for i := 0; i < n; i++ {
		if P[i] > 0 {
			cnt[P[i]&1]--
		}
	}

	dp := make([][][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][][]int, K+1)
		for k := 0; k <= K; k++ {
			dp[i][k] = make([][]int, cnt[1]+1)
			for j := 0; j <= cnt[1]; j++ {
				dp[i][k][j] = make([]int, 2)
			}
		}
	}

	dp[0][0][0][0] = 1
	dp[0][0][0][1] = 1

	var zeros int
	for i := 1; i <= n; i++ {
		if P[i-1] == 0 {
			zeros++
		}
		for k := 1; k <= i && k <= K; k++ {
			for o := 0; o <= cnt[1]; o++ {
				if P[i-1] > 0 {
					par := P[i-1] & 1
					dp[i][k][o][1-par] = 0
					dp[i][k][o][par] = modAdd(dp[i-1][k][o][par], dp[i-1][k-1][o][par^1])
				} else {
					if o > 0 {
						dp[i][k][o][1] = modMul(modAdd(dp[i-1][k][o-1][1], dp[i-1][k-1][o-1][0]), cnt[1]-o+1)
					}
					if zeros-o <= cnt[0] {
						rem := cnt[0] + o - zeros + 1
						dp[i][k][o][0] = modMul(modAdd(dp[i-1][k-1][o][1], dp[i-1][k][o][0]), rem)
					}
				}
			}
		}
	}

	return modAdd(dp[n][K][cnt[1]][0], dp[n][K][cnt[1]][1])
}

func solve1(n int, K int, P []int) int {
	// 0 的位置时unknown yet
	cnt := make([]int, n+1)
	evens := n / 2
	odds := (n + 1) / 2
	for i, num := range P {
		cnt[i+1] = cnt[i]
		if num == 0 {
			cnt[i+1]++
		} else {
			if num&1 == 0 {
				evens--
			} else {
				odds--
			}
		}
	}

	dp := make([][][][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([][][]int, K+1)
		for k := 0; k <= K; k++ {
			dp[i][k] = make([][]int, evens+1)
			for j := 0; j <= evens; j++ {
				dp[i][k][j] = make([]int, 2)
				for p := 0; p < 2; p++ {
					dp[i][k][j][p] = -1
				}
			}
		}
	}

	dp[0][0][0][0] = 1
	dp[0][0][0][1] = 1

	var f func(pos int, k int, even int, par int) int

	f = func(pos int, k int, even int, par int) int {
		if k < 0 || even < 0 {
			return 0
		}

		if dp[pos][k][even][par] >= 0 {
			return dp[pos][k][even][par]
		}

		if pos == 0 {
			return 0
		}

		if P[pos-1] > 0 {
			if P[pos-1]&1 != par {
				// not same parity
				return 0
			}

			dp[pos][k][even][par] = modAdd(f(pos-1, k, even, par), f(pos-1, k-1, even, 1-par))
		} else {
			// 到目前为止使用的x + 剩余的odd = odds
			// x + evens - even = cnt[n] - cnt[pos]
			// x = cnt[n] - cnt[pos] - (evens - even)

			dp[pos][k][even][par] = 0

			if par == 0 {
				// need to place a even number here
				x := f(pos-1, k, even-1, par)
				y := f(pos-1, k-1, even-1, 1-par)
				dp[pos][k][even][par] = modMul(max(0, even), modAdd(x, y))
			} else {
				odd := odds - (cnt[n] - cnt[pos]) + (evens - even)
				x := f(pos-1, k, even, par)
				y := f(pos-1, k-1, even, 1-par)
				dp[pos][k][even][par] = modMul(max(0, odd), modAdd(x, y))
			}
		}

		return dp[pos][k][even][par]
	}

	total_ways := modAdd(f(n, K, evens, 0), f(n, K, evens, 1))

	return total_ways
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
