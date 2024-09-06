package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	first := readNNums(reader, 4)
	n, h, l, r := first[0], first[1], first[2], first[3]
	a := readNNums(reader, n)

	res := solve(a, h, l, r)

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

const inf = 1 << 30

func solve(a []int, h int, l int, r int) int {

	dp := make([]int, h)
	dp[0] = 0
	for i := 1; i < h; i++ {
		dp[i] = -inf
	}
	fp := make([]int, h)

	check := func(v int) int {
		v %= h
		if v >= l && v <= r {
			return 1
		}
		return 0
	}

	for _, num := range a {
		for i := 0; i < h; i++ {
			fp[i] = -inf
		}
		for i := 0; i < h; i++ {
			fp[(i+num)%h] = max(fp[(i+num)%h], dp[i]+check(i+num))
			fp[(i+num-1)%h] = max(fp[(i+num-1)%h], dp[i]+check(i+num-1))
		}

		copy(dp, fp)
	}

	return slices.Max(dp)
}

func solve1(a []int, h int, l int, r int) int {
	n := len(a)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, h)
		for j := 0; j < h; j++ {
			dp[i][j] = -1
		}
	}

	// 当到达i时，开始的时刻是j时的最大值， 不需要去重的
	var dfs func(i int, j int) int

	dfs = func(i int, j int) (ans int) {
		if i == n {
			return 0
		}

		if dp[i][j] >= 0 {
			return dp[i][j]
		}

		defer func() {
			dp[i][j] = ans
		}()

		x := dfs(i+1, (j+a[i])%h)
		if (j+a[i])%h >= l && (j+a[i])%h <= r {
			x++
		}

		y := dfs(i+1, (j+a[i]-1)%h)

		if (j+a[i]-1)%h >= l && (j+a[i]-1)%h <= r {
			y++
		}

		ans = max(x, y)
		return
	}

	ans := dfs(0, 0)

	return ans
}
