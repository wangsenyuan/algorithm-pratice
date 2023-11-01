package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = readNNums(reader, n-i)
	}

	res := solve(a)

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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(a [][]int) int {

	n := len(a)

	if a[0][0] == 2 {
		return 0
	}

	check := func(cnt int, last int) bool {
		for i := 0; i < cnt; i++ {
			if a[i][cnt-1-i] == 0 {
				continue
			}
			if a[i][cnt-1-i] == 1 && last > i {
				return false
			}
			if a[i][cnt-1-i] == 2 && last <= i {
				return false
			}
		}
		return true
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	// dp[i][j] = s[i] != s[j]
	// 0 or 1
	dp[1][0] = 2

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			for _, k := range []int{j, i} {
				if check(i+1, k) {
					dp[i+1][k] = add(dp[i+1][k], dp[i][j])
				}
			}
		}
	}

	var ans int
	for j := 0; j < n; j++ {
		ans = add(ans, dp[n][j])
	}

	return ans
}
