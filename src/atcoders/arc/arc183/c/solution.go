package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	conditions := make([][]int, m)
	for i := 0; i < m; i++ {
		conditions[i] = readNNums(reader, 3)
	}
	return solve(n, conditions)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func solve(n int, conditions [][]int) int {
	F := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(i, F[i-1])
	}

	I := make([]int, n+1)
	I[n] = pow(F[n], mod-2)
	for i := n - 1; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}

	nCr := func(n int, r int) int {
		if n < r || r < 0 {
			return 0
		}
		return mul(F[n], mul(I[r], I[n-r]))
	}

	type pair struct {
		first  int
		second int
	}
	ban := make([][]pair, n+1)
	for _, cur := range conditions {
		l, r, x := cur[0], cur[1], cur[2]
		ban[r] = append(ban[r], pair{l, x})
	}

	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}
	dp[n+1][n] = 1

	marked := make([]bool, n+2)

	for l := n; l > 0; l-- {
		// clear(marked)
		for x := 0; x <= n+1; x++ {
			marked[x] = false
		}
		dp[l][l-1] = 1
		for r := l; r <= n; r++ {
			for _, p := range ban[r] {
				if p.first >= l {
					marked[p.second] = true
				}
			}
			for x := l; x <= r; x++ {
				if !marked[x] {
					// 如果最大值在x这里
					dp[l][r] = add(dp[l][r], mul(nCr(r-l, x-l), mul(dp[l][x-1], dp[x+1][r])))
				}
			}
		}
	}

	return dp[1][n]
}
