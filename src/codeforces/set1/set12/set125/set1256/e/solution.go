package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	n := readNum(reader)
	a := readNNums(reader, n)

	profit, cnt, set := solve(a)

	buf.WriteString(fmt.Sprintf("%d %d\n", profit, cnt))
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", set[i]))
	}
	buf.WriteByte('\n')

	fmt.Println(buf.String())
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
	first  int
	second int
}

const inf = 1 << 60

func solve(a []int) (profit int, team int, set []int) {

	n := len(a)
	students := make([]pair, n)
	for i, x := range a {
		students[i] = pair{x, i}
	}

	slices.SortFunc(students, func(x, y pair) int {
		return x.first - y.first
	})
	// dp[i] = 以i最后一组的最后一个人的最优解
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			dp[i][j] = inf
		}
	}
	// dp[i][0] 表示，从i开始的分组
	// dp[i][1] 表示，从i的前面开始的分组，且i不是最后一个
	// dp[i][2] 表示以i结束的分组
	for i := 0; i < n; i++ {
		if i == 0 {
			// 只能从它开始
			dp[i][0] = -students[i].first
			continue
		}
		// 以i开始的状态
		dp[i][0] = dp[i-1][2] - students[i].first
		// i在中间的位置
		dp[i][1] = min(dp[i-1][0], dp[i-1][1])
		// 以i结尾的状态
		dp[i][2] = dp[i-1][1] + students[i].first
	}

	profit = dp[n-1][2]
	set = make([]int, n)

	res := profit
	ma := students[n-1].first
	var cnt int
	team = 1
	for i := n - 1; i >= 0; i-- {
		// 现在需要找到一个分组, dp[i-1][2] + ma - mn = res
		set[students[i].second] = team
		cnt++
		if cnt >= 3 && i > 0 && ma-students[i].first+dp[i-1][2] == res {
			cnt = 0
			ma = students[i-1].first
			res = dp[i-1][2]
			team++
		}
	}

	return
}
