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
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const inf = 1 << 50

func solve(a []int) int {
	n := len(a)
	// dp[i]到i为止能够得到的最小操作数
	dp := make([]int, 3)
	dp[0] = 0
	// dp[i][1]表示前面剩余 <= 2 black cell的情况
	dp[1] = inf
	// dp[i][2] 表示前面剩余 <= 4个black cell的情况（但是前面两个被清理的情况
	dp[2] = inf

	for i := 1; i <= n; i++ {
		fp := make([]int, 3)
		fp[0] = dp[0] + 1
		fp[1] = inf
		fp[2] = inf
		if a[i-1] == 0 {
			fp[0] = dp[0]
		} else if a[i-1] <= 2 {
			// 使用操作1或者操作2
			fp[0] = min(dp[0], dp[1]) + 1
			// 保留前面两个
			fp[1] = dp[0]
			// 无效的状态
		} else if a[i-1] <= 4 {
			// 在尾部使用操作1, 但必须保证i-1的被清理完了
			// 所以要么前面的已经被处理完了，要么前面正好剩下后面两个，在当前全部处理掉
			fp[1] = min(dp[0], dp[2]) + 1
			// 在头部使用操作1，但是必须保证i-1的被清理了
			fp[2] = min(dp[0], dp[1]) + 1
		}
		// else 只使用操作2
		copy(dp, fp)
	}

	return dp[0]
}

func check(a bool) int {
	if a {
		return 1
	}
	return 0
}
