package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)

	fmt.Println(solve(s))
}

const inf = 1 << 30

func solve(s string) int {
	n := len(s)

	m := n * n

	dp := make([][][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int, m+1)
		}
	}
	dp[0][0][0] = 0

	var p int
	var cnt0 int
	for i := 0; i < n; i++ {
		for j := 0; j <= n; j++ {
			for k := 0; k <= m; k++ {
				dp[1^p][j][k] = inf
			}
		}
		for j := 0; j <= i; j++ {
			for k := 0; k <= j*(i-j); k++ {
				// 如果s[i] = 0, 则不需要交换就可以到达状态j+1
				// 否则必须要有一次交换才可以
				dp[1^p][j+1][k] = min(dp[1^p][j+1][k], dp[p][j][k]+getValue(s[i] != '0'))

				dp[1^p][j][k+j] = min(dp[1^p][j][k+j], dp[p][j][k]+getValue(s[i] == '0'))
			}
		}
		if s[i] == '0' {
			cnt0++
		}
		p ^= 1
	}

	res := dp[p][cnt0][cnt0*(n-cnt0)/2]
	return res / 2
}

func getValue(b bool) int {
	if b {
		return 1
	}
	return 0
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
