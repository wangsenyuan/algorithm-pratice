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
		res := solve(m, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

const MOD = 1000000007

func solve(m int, A []int) int {
	n := len(A)
	dp := process(n, m)

	var ans int64 = 1

	vis := make([]int, n)

	for i := 0; i < n; i++ {
		if vis[i] == 0 {
			var cnt int
			var cur = i
			for vis[cur] == 0 {
				cnt++
				vis[cur] = 1
				cur = A[cur] - 1
			}

			if vis[cur] == 2 {
				ans = mul(ans, pow(m-1, cnt))
			} else {
				var cyc = 1
				for i := A[cur] - 1; i != cur; i = A[i] - 1 {
					cyc++
				}
				tmp := dp[cyc][0]
				ans = mul(ans, tmp)
				ans = mul(ans, pow(m-1, cnt-cyc))
			}
			cur = i
			for vis[cur] != 2 {
				vis[cur] = 2
				cur = A[cur] - 1
			}
		}
	}

	return int(ans)
}

func process(n int, m int) [][]int64 {
	dp := make([][]int64, n+1)
	//dp[i][0] = a[i] != a[1]
	//dp[i][1] = a[i] == a[1]
	dp[1] = make([]int64, 2)
	dp[1][0] = 0
	dp[1][1] = int64(m)
	for i := 2; i <= n; i++ {
		dp[i] = make([]int64, 2)
		// a[i] == a[1], and a[i] != a[i-1]
		dp[i][1] = dp[i-1][0]
		dp[i][0] = add(mul(int64(m-2), dp[i-1][0]), mul(int64(m-1), dp[i-1][1]))
	}
	return dp
}

func pow(a, b int) int64 {
	A := int64(a)
	var R int64 = 1
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return R
}

func add(a, b int64) int64 {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func sub(a, b int64) int64 {
	return add(a, MOD-b)
}

func mul(a, b int64) int64 {
	a *= b
	return a % MOD
}
