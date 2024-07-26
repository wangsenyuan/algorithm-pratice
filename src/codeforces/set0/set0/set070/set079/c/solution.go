package main

import (
	"bufio"
	"fmt"
	"index/suffixarray"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	n := readNum(reader)
	b := make([]string, n)
	for i := 0; i < n; i++ {
		b[i] = readString(reader)
	}

	res := solve(s, b)
	fmt.Printf("%d %d\n", res[0], res[1])
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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
func solve(s string, b []string) []int {
	n := len(s)

	dp := make([]int, n+2)
	for i := 0; i <= n; i++ {
		dp[i] = n + 1
	}
	fp := make([]bool, n)
	for _, x := range b {
		next := kmp(x)
		w := len(x)
		for i, j := 0, 0; i < n; i++ {
			for j > 0 && s[i] != x[j] {
				j = next[j-1]
			}
			if s[i] == x[j] {
				j++
			}
			if j == w {
				fp[i-w+1] = true
				j = next[j-1]
			}
		}
		pos := n + 1
		for i := n - 1; i >= 0; i-- {
			if fp[i] {
				pos = i
			}
			dp[i] = min(dp[i], dp[i+1], pos+w)
		}
		clear(fp)
	}

	ans := []int{0, 0}

	for i := 0; i < n; i++ {
		if dp[i]-i-1 > ans[0] {
			ans[0] = dp[i] - i - 1
			ans[1] = i
		}
	}

	return ans
}

func kmp(p string) []int {
	n := len(p)
	res := make([]int, n)
	for i := 1; i < n; i++ {
		j := res[i-1]
		for j > 0 && p[j] != p[i] {
			j = res[j-1]
		}
		if p[j] == p[i] {
			j++
		}
		res[i] = j
	}
	return res
}
func solve1(s string, b []string) []int {
	sa := suffixarray.New([]byte(s))

	n := len(s)

	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = n + 1
	}

	for _, x := range b {
		w := len(x)
		cur := sa.Lookup([]byte(x), -1)
		sort.Ints(cur)
		for i, j := n-1, len(cur)-1; i >= 0; i-- {
			dp[i] = min(dp[i], dp[i+1])
			for j >= 0 && i <= cur[j] {
				dp[i] = min(dp[i], cur[j]+w)
				j--
			}
		}
	}

	ans := []int{0, 0}

	for i := 0; i < n; i++ {
		if dp[i]-i-1 > ans[0] {
			ans[0] = dp[i] - i - 1
			ans[1] = i
		}
	}

	return ans
}
