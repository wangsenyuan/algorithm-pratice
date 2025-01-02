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

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	s := readString(reader)[:n]
	return solve(s, k)
}

const inf = 1e12

func solve(s string, k int) int {
	n := len(s)
	lst := make([][]int, n)
	for i := 0; i < n; i++ {
		lst[i] = make([]int, 26)
		for j := 0; j < 26; j++ {
			lst[i][j] = -1
		}
		if i > 0 {
			copy(lst[i], lst[i-1])
		}
		lst[i][int(s[i]-'a')] = i
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n+1)
		dp[i][1] = 1
	}

	for j := 2; j <= n; j++ {
		for i := 1; i < n; i++ {
			for c := 0; c < 26; c++ {
				if lst[i-1][c] >= 0 {
					dp[i][j] = min(inf, dp[i][j]+dp[lst[i-1][c]][j-1])
				}
			}
		}
	}

	k--
	var res int

	for x := n - 1; x > 0 && k > 0; x-- {
		var cnt int
		for c := 0; c < 26; c++ {
			if lst[n-1][c] >= 0 {
				cnt += dp[lst[n-1][c]][x]
			}
		}
		j := min(k, cnt)
		res += j * (n - x)
		k -= j
	}

	if k > 0 {
		res += n
		k--
	}
	if k > 0 {
		return -1
	}
	return res
}

func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}
