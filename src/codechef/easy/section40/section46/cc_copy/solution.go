package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		s := readString(reader)
		res := solve(s)
		buf.WriteString(fmt.Sprintf("%s\n", res))
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
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
	if len(bs) == 0 || bs[0] == '\n' {
		return readNum(reader)
	}
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

const codechef = "codechef"

func solve(s string) string {
	cnt := make([]int, 3)
	for i := 0; i < 8; i++ {
		if s[i] == 'c' {
			cnt[0]++
		} else if s[i] == 'e' {
			cnt[1]++
		}

		if s[i] == s[0] {
			cnt[2]++
		}
	}

	check := func(s string) bool {
		for i := 0; i < 8; i++ {
			if codechef[i] == s[i] {
				return false
			}
		}
		return true
	}

	if cnt[0] >= 7 || cnt[1] >= 7 || cnt[2] == 8 && !check(s) {
		return "-1"
	}

	for !check(s) {
		buf := []byte(s)
		rand.Shuffle(8, func(i, j int) {
			buf[i], buf[j] = buf[j], buf[i]
		})
		s = string(buf)
	}
	return s
}

func solve2(s string) string {

	buf := []byte(s)

	var dfs func(i int, state int) bool

	dfs = func(i int, state int) bool {
		if i == 8 {
			return true
		}
		for j := 0; j < 8; j++ {
			if codechef[i] == s[j] || (state>>j)&1 == 1 {
				continue
			}
			buf[i] = s[j]
			tmp := dfs(i+1, state|(1<<j))
			if tmp {
				return true
			}
		}
		return false
	}

	res := dfs(0, 0)

	if res {
		return string(buf)
	}
	return "-1"
}

// const x = 8 * 7 * 6 * 5 * 4 * 3 * 2
func solve1(s string) string {
	// len(s) = 8
	// try all possible anagram of s could be possible
	n := len(s)
	N := 1 << n
	// 前i个是否可以被mask正确的处理
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = -1
		}
	}
	for j := 0; j < n; j++ {
		if codechef[0] != s[j] {
			dp[0][1<<j] = 0
		}
	}

	for i := 1; i < n; i++ {
		// try to put a new x at position i
		for state := 0; state < N; state++ {
			if dp[i-1][state] < 0 || digitCount(state) != i {
				continue
			}
			// dp[i-1][state] is valid
			for j := 0; j < n; j++ {
				if (state>>j)&1 == 1 || codechef[i] == s[j] {
					continue
				}
				dp[i][state|(1<<j)] = state
			}
		}
	}
	if dp[n-1][N-1] < 0 {
		return "-1"
	}
	cur := N - 1
	buf := make([]byte, n)

	for i := n - 1; i >= 0; i-- {
		prev := dp[i][cur]
		for j := 0; j < n; j++ {
			if (prev>>j)&1 == 0 && (cur>>j)&1 == 1 {
				buf[i] = s[j]
				break
			}
		}
		cur = prev
	}

	return string(buf)
}

func digitCount(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}
