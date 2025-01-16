package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	res := solve(n, k)
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

func solve(n int, kk int) int {
	dp := make([][][]int, 2)
	for i := range 2 {
		dp[i] = make([][]int, n+2)
		for j := range n + 2 {
			dp[i][j] = make([]int, n+2)
		}
	}
	dp[0][0][0] = 1

	for ii := range n {
		i := ii & 1
		ni := i ^ 1
		for j := range n + 1 {
			clear(dp[ni][j])
		}
		for j := range n + 1 {
			for k := range n + 1 {
				// 使用相同的颜色
				dp[ni][j+1][max(j+1, k)] = add(dp[ni][j+1][max(j+1, k)], dp[i][j][k])
				// 使用不同的颜色
				dp[ni][1][max(1, k)] = add(dp[ni][1][max(1, k)], dp[i][j][k])
			}
		}
	}

	cnt := make([]int, n+2)
	for i := range n + 1 {
		for j := range n + 1 {
			cnt[i] = add(cnt[i], dp[n&1][j][i])
		}
	}

	pr := make([]int, n+2)
	for i := range n + 1 {
		pr[i+1] = add(pr[i+1], pr[i])
		pr[i+1] = add(pr[i+1], cnt[i])
	}

	var ans int

	for i := 1; i <= n; i++ {
		// cnt[i]表示最长的一段连续相同颜色的长度为i时的计数
		// pr[?] 表示cnt[0...i]的sum
		// 当其中一条边长为i时，另一条边的最长为 (kk - 1) / i
		ans = add(ans, mul(pr[min(n+1, (kk-1)/i+1)], cnt[i]))
	}

	return mul(ans, pow(2, mod-2))
}
