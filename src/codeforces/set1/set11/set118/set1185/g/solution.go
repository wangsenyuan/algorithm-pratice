package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, x := readTwoNums(reader)
	songs := make([][]int, n)
	for i := 0; i < n; i++ {
		songs[i] = readNNums(reader, 2)
	}

	res := solve(x, songs)

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

const mod = 1e9 + 7

func add(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		res += a[i]
		if res >= mod {
			res -= mod
		}
	}
	return res
}

func mul(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		res *= a[i]
		res %= mod
	}
	return res
}

func solve(T int, songs [][]int) int {
	n := len(songs)

	f := make([][]int, n+2)
	// f[i][t] 表示选择i个0类歌曲，且总时长=t的情况
	for i := 0; i <= n+1; i++ {
		f[i] = make([]int, T+1)
	}

	f[0][0] = 1
	// g[i][j][t] 表示选择i个1类歌曲，j个2类歌曲，且总时长为t的情况
	g := make([][][]int, n+2)
	for i := 0; i <= n+1; i++ {
		g[i] = make([][]int, n+2)
		for j := 0; j <= n+1; j++ {
			g[i][j] = make([]int, T+1)
		}
	}
	g[0][0][0] = 1
	var cnt [3]int
	for _, cur := range songs {
		w, x := cur[0], cur[1]-1
		if x == 0 {
			for j := cnt[0]; j >= 0; j-- {
				for t := T; t >= w; t-- {
					f[j+1][t] = add(f[j+1][t], f[j][t-w])
				}
			}
		} else {
			var is [3]int
			is[x] = 1
			for j := cnt[1]; j >= 0; j-- {
				for k := cnt[2]; k >= 0; k-- {
					for t := T; t >= w; t-- {
						g[j+is[1]][k+is[2]][t] = add(g[j+is[1]][k+is[2]][t], g[j][k][t-w])
					}
				}
			}
		}

		cnt[x]++
	}

	fac := make([]int, n+1)
	fac[0] = 1
	for i := 1; i <= n; i++ {
		fac[i] = mul(i, fac[i-1])
	}

	dp := make([][][][3]int, cnt[0]+2)
	for i := 0; i <= cnt[0]+1; i++ {
		dp[i] = make([][][3]int, cnt[1]+2)
		for j := 0; j <= cnt[1]+1; j++ {
			dp[i][j] = make([][3]int, cnt[2]+2)
		}
	}

	dp[1][0][0][0] = 1
	dp[0][1][0][1] = 1
	dp[0][0][1][2] = 1

	var ans int

	for i, mat := range dp[:cnt[0]+1] {
		for j, row := range mat[:cnt[1]+1] {
			for k, comb := range row[:cnt[2]+1] {
				var sum int
				for t, fit := range f[i] {
					sum = add(sum, mul(fit, g[j][k][T-t]))
				}
				tmp := mul(fac[i], fac[j], fac[k], sum, add(comb[0], comb[1], comb[2]))
				ans = add(ans, tmp)
				dp[i+1][j][k][0] = add(dp[i+1][j][k][0], comb[1], comb[2])
				dp[i][j+1][k][1] = add(dp[i][j+1][k][1], comb[0], comb[2])
				dp[i][j][k+1][2] = add(dp[i][j][k+1][2], comb[0], comb[1])
			}
		}
	}

	return ans
}
