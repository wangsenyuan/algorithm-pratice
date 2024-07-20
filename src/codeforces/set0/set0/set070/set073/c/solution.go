package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	first := readString(reader)
	var i int
	for first[i] != ' ' {
		i++
	}
	s := first[:i]
	var k int
	readInt([]byte(first), i+1, &k)
	n := readNum(reader)
	bonus := make([]string, n)
	for i := 0; i < n; i++ {
		bonus[i] = readString(reader)
	}
	res := solve(s, k, bonus)
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

const inf = 1 << 30

func solve(s string, k int, bonus []string) int {
	mat := make([][]int, 26)
	for i := 0; i < 26; i++ {
		mat[i] = make([]int, 26)
	}

	for _, cur := range bonus {
		a, b := int(cur[0]-'a'), int(cur[2]-'a')
		var c int
		readInt([]byte(cur), 4, &c)
		mat[a][b] = c
	}

	dp := make([][]int, k+2)
	fp := make([][]int, k+2)
	for i := 0; i <= k+1; i++ {
		dp[i] = make([]int, 26)
		fp[i] = make([]int, 26)
		for j := 0; j < 26; j++ {
			dp[i][j] = -inf
			fp[i][j] = -inf
		}
	}
	dp[0][int(s[0]-'a')] = 0
	for j := 0; j < 26; j++ {
		dp[1][j] = 0
	}
	n := len(s)

	for i := 1; i < n; i++ {
		c := int(s[i] - 'a')
		for x := 0; x <= k; x++ {
			for j := 0; j < 26; j++ {
				// 如果修改这里为u时，
				for u := 0; u < 26; u++ {
					fp[x+check(u != c)][u] = max(fp[x+check(u != c)][u], dp[x][j]+mat[j][u])
				}
			}
		}
		for x := 0; x <= k; x++ {
			for j := 0; j < 26; j++ {
				dp[x][j] = fp[x][j]
				fp[x][j] = -inf
			}
		}
	}
	var res int = -inf

	for x := 0; x <= k; x++ {
		for j := 0; j < 26; j++ {
			res = max(res, dp[x][j])
		}
	}

	return res
}

func check(b bool) int {
	if b {
		return 1
	}
	return 0
}
