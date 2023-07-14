package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	r, n := readTwoNums(reader)
	celebrities := make([][]int, n)
	for i := 0; i < n; i++ {
		celebrities[i] = readNNums(reader, 3)
	}
	res := solve(r, celebrities)
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

func solve(r int, celebrities [][]int) int {
	n := len(celebrities)
	mx := make([]int, n+1)

	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -n
	}
	// dp[0] is for (0, 0)
	dp[0] = 0

	canReach := func(i, j int) bool {
		dx := abs(celebrities[i][1] - celebrities[j][1])
		dy := abs(celebrities[i][2] - celebrities[j][2])
		return dx+dy <= celebrities[j][0]-celebrities[i][0]
	}

	for i := 1; i <= n; i++ {
		cur := celebrities[i-1]
		t, x, y := cur[0], cur[1], cur[2]
		// for those j (i - j) >= 2 * r
		// j <= i - 2r
		dp[i] = -n
		if x+y-2 <= t {
			dp[i] = 1
		}
		if i-2*r >= 0 {
			dp[i] = mx[i-2*r] + 1
		}

		for j := max(0, i-2*r) + 1; j < i; j++ {
			if canReach(j-1, i-1) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		mx[i] = max(mx[i-1], dp[i])
	}

	return mx[n]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
