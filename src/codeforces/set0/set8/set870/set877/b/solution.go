package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)

	res := solve(s)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(s string) int {
	// n * n 肯定是可以的，有没有更好的方法?
	// dp[i] 表示到i为止a的数量
	// fp[i] 表示到i为止suf的a的数量
	// xp[i] 表示到i为止，以b结尾的，符合aaaabbb的最优值
	// 如果 s[i] = b, 那么 xp[i] = max(xp[i-1], dp[i-1]) + 1
	n := len(s)

	dp := make([]int, n)
	fp := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 {
			dp[i] = dp[i-1]
			fp[i] = fp[i-1]
		}
		if s[i] == 'a' {
			dp[i]++
		} else {
			if i == 0 {
				fp[i] = 1
			} else {
				fp[i] = max(fp[i-1], dp[i-1]) + 1
			}
		}
	}
	var best = max(dp[n-1], fp[n-1])
	var suf int

	for i := n - 1; i > 0; i-- {
		if s[i] == 'a' {
			suf++
		}
		best = max(best, suf+max(dp[i-1], fp[i-1]))
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
