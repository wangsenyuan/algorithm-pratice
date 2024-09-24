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

		readString(reader)
		grid := make([]string, 3)
		grid[0] = readString(reader)
		grid[1] = readString(reader)
		grid[2] = readString(reader)

		res := solve(grid)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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
	first  int
	second int
}

func solve(grid []string) bool {
	n := len(grid[0])
	dp := make([][]bool, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < 3; i++ {
		if grid[i][0] == 's' {
			dp[i][0] = true
		}
	}

	trains := make([][]*pair, 3)

	for i := 0; i < 3; i++ {
		for j := 0; j < n; {
			if grid[i][j] == '.' || grid[i][j] == 's' {
				j++
				continue
			}
			l := j
			// 不同的train， 连在一起，算一格，貌似也没有关系
			for j < n && grid[i][j] != '.' {
				j++
			}
			trains[i] = append(trains[i], &pair{l, j - 1})
		}
	}

	check2 := func(i int, j int) bool {
		if j > (n+2)/3+1 {
			return true
		}
		// 在当前时刻，没有火车头处在位置 (i, j)
		for _, t := range trains[i] {
			if t.first == j {
				return false
			}
		}
		return true
	}

	check := func(i int, j int) bool {
		if j <= (n+2)/3+1 {
			// 超过时间n/3后，最右边的火车尾巴也移动到了外面，不需要检查了
			// 必须知道i行，当前位置后2秒内是否安全
			for _, t := range trains[i] {
				l, r := t.first, t.second
				if l-2 <= j && j <= r {
					return false
				}
			}
		}

		for _, d := range []int{-1, 0, 1} {
			if i+d >= 0 && i+d < 3 && dp[i+d][j-1] && check2(i+d, j) {
				return true
			}
		}

		return false
	}

	for j := 1; j < n; j++ {
		// for row 0
		for i := 0; i < 3; i++ {
			dp[i][j] = check(i, j)
			// move trains
		}
		for i := 0; i < 3; i++ {
			for _, t := range trains[i] {
				t.first -= 2
				t.second -= 2
			}
		}
	}
	return dp[0][n-1] || dp[1][n-1] || dp[2][n-1]
}
