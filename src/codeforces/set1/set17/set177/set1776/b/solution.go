package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, h := readTwoNums(reader)
	x := readNNums(reader, n)
	res := solve(h, x)

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

const inf = 1 << 62

func solve(h int, x []int) int {
	n := len(x)
	// already sorted
	// i, i + 1, 如果要共享底座，它们要求的高度是
	// 假设它们要共享一个base, 2 * d - 1 >= dist
	//  如果是3个要共享呢？4个呢？
	// 假设 l...r段的要共享一个基座，需要的高度是多少？需要多少个基座
	//  是不是一定是，先移动到两个点，然后再往两边吗？
	// 假设知道 l...r需要的最少的基座
	// 怎么求l。。。r中间需要的最少的基座呢？它们肯定要从中点开始
	// 然后形成一个三角形，连接到两条边
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			dp[i][j] = inf
		}
		dp[i][i] = h
	}

	for r := 0; r < n; r++ {
		for l := r - 1; l >= 0; l-- {
			// dp[l][r]
			dist := x[r] - x[l]
			// 2 * d - 1 >= dist
			d := max(0, h+1-(dist+1)/2)

			for k := l; k < r; k++ {
				dp[l][r] = min(dp[l][r], dp[l][k]+dp[k+1][r]-d)
			}
		}
	}

	return dp[0][n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
