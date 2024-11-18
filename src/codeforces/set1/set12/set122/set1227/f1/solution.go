package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	h := readNNums(reader, n)
	res := solve(n, k, h)
	fmt.Println(res)
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

func solve(n int, k int, h []int) int {
	dp := make([]int, 2*n+3)

	update := func(dp []int, i int, v int) {
		dp[i+n+1] = add(dp[i+n+1], v)
	}

	if h[0] == h[n-1] {
		// 任何答案，前后的得分都一样
		update(dp, 0, k)
	} else {
		// ans[n-1] = h[0]
		update(dp, 1, 1)
		update(dp, 0, k-2)
		// 如果ans[n-1] = h[n-1]
		update(dp, -1, 1)
	}

	get := func(dp []int, i int) int {
		return dp[i+n+1]
	}

	fp := make([]int, 2*n+3)

	for i := n - 2; i >= 0; i-- {
		// dp[j] = (k - 1) * dp[j] + dp[j-1]
		clear(fp)
		for j := n - i; j >= i-n; j-- {
			if h[i] == h[i+1] {
				// 不论答案，分数保持不变
				update(fp, j, mul(get(dp, j), k))
			} else {
				// 如果回答h[i+1], +1分
				update(fp, j, get(dp, j-1))
				// 回答不是h[i+1], h[i]， 分数不变
				update(fp, j, mul(k-2, get(dp, j)))
				// 如果回答h[i], 减一分
				update(fp, j, get(dp, j+1))
			}
		}
		copy(dp, fp)
	}

	var ans int

	for i := 1; i <= n; i++ {
		ans = add(ans, get(dp, i))
	}

	return ans
}
