package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	x := readNNums(reader, n)
	holes := make([][]int, m)
	for i := 0; i < m; i++ {
		holes[i] = readNNums(reader, 2)
	}

	res := solve(x, holes)

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

const inf = 1 << 60

type Pair struct {
	first  int
	second int
}

func solve(x []int, holes [][]int) int {
	sort.Ints(x)
	n := len(x)
	m := len(holes)

	sort.Slice(holes, func(i, j int) bool {
		return holes[i][0] < holes[j][0]
	})

	// dp[i][j] = 前j个老鼠进入前i个hole的最优答案
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = inf
	}
	dp[0] = 0
	que := make([]Pair, n+1)

	for i := 1; i <= m; i++ {
		var front, end int

		que[front] = Pair{0, 0}
		front++

		var sum int

		for j := 0; j < n; j++ {
			for end < front && j+1-que[end].second > holes[i-1][1] {
				end++
			}
			sum += abs(x[j] - holes[i-1][0])
			cur := dp[j+1] - sum
			dp[j+1] = min(dp[j+1], que[end].first+sum)

			for end < front && que[front-1].first >= cur {
				front--
			}
			que[front] = Pair{cur, j + 1}
			front++
		}
	}

	if dp[n] >= inf {
		return -1
	}

	return dp[n]
}

func abs(num int) int {
	return max(num, -num)
}
