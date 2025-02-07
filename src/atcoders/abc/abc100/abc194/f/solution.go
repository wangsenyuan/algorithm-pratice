package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	i := strings.IndexByte(s, ' ')
	x := s[:i]
	k, _ := strconv.Atoi(s[i+1:])
	res := solve(k, x)
	fmt.Println(res)
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

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(num ...int) int {
	res := 1
	for _, x := range num {
		res = res * x % mod
	}
	return res
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

func inverse(a int) int {
	return pow(a, mod-2)
}

func div(a, b int) int {
	return mul(a, inverse(b))
}

func solve(k int, s string) int {
	n := len(s)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = -1
		}
	}

	const letters = "0123456789ABCDEF"

	id := func(x byte) int {
		return strings.IndexByte(letters, x)
	}

	var dfs func(i int, mask int, limit bool) int

	dfs = func(i int, mask int, limit bool) (res int) {
		x := bits.OnesCount(uint(mask))
		if x > k {
			return 0
		}
		if i == n {
			if x == k {
				return 1
			}
			return 0
		}
		if !limit && mask > 0 {
			if dp[i][x] >= 0 {
				return dp[i][x]
			}
			defer func() {
				dp[i][x] = res
			}()
		}
		if mask == 0 {
			res = dfs(i+1, 0, false)
		}
		up := 15
		if limit {
			up = id(s[i])
		}
		var d int
		if mask == 0 {
			d++
		}
		for d <= up {
			res = add(res, dfs(i+1, mask|(1<<d), limit && d == up))
			d++
		}
		return
	}

	return dfs(0, 0, true)
}
