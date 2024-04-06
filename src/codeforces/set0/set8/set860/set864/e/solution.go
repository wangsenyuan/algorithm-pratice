package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	items := make([][]int, n)
	for i := 0; i < n; i++ {
		items[i] = readNNums(reader, 3)
	}

	res, ord := solve(n, items)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n%d\n", res, len(ord)))

	for i := 0; i < len(ord); i++ {
		buf.WriteString(fmt.Sprintf("%d ", ord[i]))
	}
	buf.WriteByte('\n')
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

func solve(n int, items [][]int) (int, []int) {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return items[id[i]][1] < items[id[j]][1]
	})

	dp := make([]int, 2000)
	ans := make([][]int, 2000)

	for i := 0; i < n; i++ {
		x := id[i]
		for j := items[x][1] - 1; j >= items[x][0]; j-- {
			if dp[j] < dp[j-items[x][0]]+items[x][2] {
				dp[j] = dp[j-items[x][0]] + items[x][2]
				ans[j] = make([]int, len(ans[j-items[x][0]])+1)
				copy(ans[j], ans[j-items[x][0]])
				ans[j][len(ans[j])-1] = x + 1
			}
		}
	}

	var best int
	for j := 0; j < len(dp); j++ {
		if dp[j] > dp[best] {
			best = j
		}
	}

	return dp[best], ans[best]
}

func solve1(n int, items [][]int) (int, []int) {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return items[id[i]][1] < items[id[j]][1]
	})

	var D int
	for i := 0; i < n; i++ {
		D = max(D, items[i][1])
	}

	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, D+1)
	}

	dp[0][0] = 0

	for i := 0; i < n; i++ {
		x := id[i]

		for j := items[x][1] - 1; j >= 0; j-- {
			if j >= items[x][0] {
				dp[i+1][j] = max(dp[i][j], dp[i][j-items[x][0]]+items[x][2])
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}
	var k int
	var best int
	for i := D; i > 0; i-- {
		if best < dp[n][i] {
			best = dp[n][i]
			k = i
		}
	}

	var ord []int

	for i, j := n, k; i > 0; i-- {
		if dp[i][j] != dp[i-1][j] {
			x := id[i-1]
			ord = append(ord, x+1)
			j -= items[x][0]
		}

	}

	reverse(ord)

	return best, ord
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxOf(arr []int) int {
	res := arr[0]
	for _, num := range arr {
		if res < num {
			res = num
		}
	}
	return res
}
