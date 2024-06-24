package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(a, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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
func solve(a []int, h int) int {

	sort.Ints(a)

	f := func(ord []int) int {
		var ans int
		cur := h
		var j int
		for i, num := range a {
			for cur <= num && j < len(ord) {
				cur *= ord[j]
				j++
			}
			if cur <= num {
				break
			}
			cur += num / 2
			ans = i + 1
		}
		return ans
	}

	ans := max(f([]int{2, 2, 3}), f([]int{2, 3, 2}), f([]int{3, 2, 2}))

	return ans
}

func solve2(a []int, h int) int {
	sort.Ints(a)
	// 0, 1, 2
	// 0 for green, 1 for green, 2 for blue
	dp := make([]int, 8)
	dp[0] = h
	for state := 0; state < 8; state++ {
		if state&1 == 0 {
			// use green at the first time
			dp[state|1] = max(dp[state|1], 2*dp[state])
		}
		if state&2 == 0 {
			dp[state|2] = max(dp[state|2], 2*dp[state])
		}
		if state&4 == 0 {
			dp[state|4] = max(dp[state|4], 3*dp[state])
		}
	}

	var ans int

	for i, num := range a {
		for state := 0; state < 8; state++ {
			if dp[state] > num {
				dp[state] = max(dp[state], dp[state]+num/2)
				ans = i + 1
			}
		}

		for state := 0; state < 8; state++ {
			if (state & 1) == 0 {
				dp[state|1] = max(dp[state|1], dp[state]*2)
			}
			if (state & 2) == 0 {
				dp[state|2] = max(dp[state|2], dp[state]*2)
			}
			if (state & 4) == 0 {
				dp[state|4] = max(dp[state|4], dp[state]*3)
			}
		}
	}

	return ans
}

func solve1(a []int, h int) int {
	sort.Ints(a)
	// dp[i][g][b] = 吸收完i后，剩余g个green，b个blue后的power
	n := len(a)
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, 3)
		for j := 0; j < 3; j++ {
			dp[i][j] = make([]int, 2)
			for k := 0; k < 2; k++ {
				dp[i][j][k] = -1
			}
		}
	}
	dp[0][2][1] = h

	var ans int
	for i := 0; i <= n; i++ {
		if i > 0 {
			for j := 0; j < 3; j++ {
				for k := 0; k < 2; k++ {
					x := dp[i-1][j][k]
					if x > a[i-1] {
						dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k]+a[i-1]/2)
					}
				}
			}
		}
		for j := 2; j >= 0; j-- {
			for k := 1; k >= 0; k-- {
				if dp[i][j][k] > 0 {
					ans = i
				}
				if j > 0 {
					dp[i][j-1][k] = max(dp[i][j-1][k], 2*dp[i][j][k])
				}
				if k > 0 {
					dp[i][j][k-1] = max(dp[i][j][k-1], 3*dp[i][j][k])
				}
			}
		}
	}

	return ans
}
