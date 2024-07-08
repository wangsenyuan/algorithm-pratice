package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, a, b := readThreeNums(reader)
		s := readString(reader)[:n]
		res := solve(s, a, b)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(res)
			buf.WriteByte('\n')
		}
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

type pair struct {
	first  int8
	second int8
}

func solve(nums string, a int, b int) string {
	n := len(nums)
	dp := make([][][][]bool, n+1)
	prev := make([][][][]pair, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][][]bool, n)
		prev[i] = make([][][]pair, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([][]bool, a)
			prev[i][j] = make([][]pair, a)
			for x := 0; x < a; x++ {
				dp[i][j][x] = make([]bool, b)
				prev[i][j][x] = make([]pair, b)
			}
		}
	}

	dp[0][0][0][0] = true

	for i := 1; i <= n; i++ {
		num := int(nums[i-1] - '0')
		for j := 0; j < i; j++ {
			// 前面j个用来作为red
			for ra := 0; ra < a; ra++ {
				for rb := 0; rb < b; rb++ {
					if dp[i-1][j][ra][rb] {
						// 如果当前作为red
						if j+1 < n {
							dp[i][j+1][(ra*10+num)%a][rb] = true
							prev[i][j+1][(ra*10+num)%a][rb] = pair{int8(ra), -1}
						}
						// 如果当前作为blue
						if i-j < n {
							dp[i][j][ra][(rb*10+num)%b] = true
							prev[i][j][ra][(rb*10+num)%b] = pair{-1, int8(rb)}
						}
					}
				}
			}
		}
	}
	// 0 个 red， n 个blue, invalid
	best := 0

	for j := 1; j < n; j++ {
		if dp[n][j][0][0] {
			// found a valid one
			k := n - j
			if abs(best-(n-best)) > abs(k-j) {
				best = j
			}
		}
	}

	if best == 0 {
		return ""
	}

	ans := make([]byte, n)
	ra, rb := int8(0), int8(0)
	for i := n; i > 0; i-- {
		cur := prev[i][best][ra][rb]
		if cur.first >= 0 {
			ans[i-1] = 'R'
			best--
			ra = cur.first
		} else {
			ans[i-1] = 'B'
			rb = cur.second
		}
	}

	return string(ans)
}

func abs(num int) int {
	return max(num, -num)
}
