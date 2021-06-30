package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	res := solve(n, k)
	var buf bytes.Buffer

	for i := 0; i < k; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	fmt.Println(buf.String())
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

const MOD = 1000000007

const K = 405

var I []int64
var F []int64
var IV []int64

func init() {
	F = make([]int64, K)
	F[0] = 1
	for i := 1; i < K; i++ {
		F[i] = (int64(i) * F[i-1]) % MOD
	}
	I = make([]int64, K)
	I[K-1] = int64(pow(int(F[K-1]), MOD-2))
	for i := K - 2; i >= 0; i-- {
		I[i] = (int64(i+1) * I[i+1]) % MOD
	}

	IV = make([]int64, K)

	IV[1] = 1
	for i := 2; i < K; i++ {
		IV[i] = IV[MOD%i] * (MOD - MOD/int64(i)) % MOD
	}
}

func pow(a, b int) int {
	A := int64(a)
	R := int64(1)
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(R)
}

func nCr(n, r int) int {
	res := F[n]
	res = (res * I[r]) % MOD
	res = (res * I[n-r]) % MOD
	return int(res)
}

func big_nCr(n, r int) int {
	var res int64 = 1
	for i := n - r + 1; i <= n; i++ {
		res = (res * int64(i)) % MOD
	}

	for i := 1; i <= r; i++ {
		res = (res * IV[i]) % MOD
	}
	return int(res)
}

func solve(n, k int) []int {
	// dp[i][j] = length n array, j swaps to make it sorted
	dp := make([][]int64, 2*k+1)
	dp[0] = make([]int64, 2*k+1)
	for i := 1; i <= 2*k; i++ {
		dp[i] = make([]int64, 2*k+1)
		dp[i][0] = 1
		for j := 1; j <= 2*k; j++ {
			dp[i][j] = (dp[i-1][j] + int64(i-1)*dp[i-1][j-1]%MOD) % MOD
		}
	}
	res := make([]int, k)
	ans := []int64{1, 0}

	for j := 1; j <= k; j++ {
		for i := 1; i <= min(n, 2*j); i++ {
			var cnt int64
			for f := 0; f <= i; f++ {
				if f%2 == 1 {
					cnt = (cnt - (int64(nCr(i, f))*dp[i-f][j])%MOD + MOD) % MOD
				} else {
					cnt = (cnt + int64(nCr(i, f))*dp[i-f][j]) % MOD
				}
			}
			ans[j%2] = (ans[j%2] + (int64(big_nCr(n, i))*cnt)%MOD) % MOD
		}
		res[j-1] = int(ans[j%2])
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
