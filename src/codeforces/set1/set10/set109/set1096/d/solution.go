package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	s := readString(reader)
	a := readNNums(reader, n)

	res := solve(s, a)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

const hard = "hard"

func solve(s string, a []int) int {
	dp := make([]int, 5)
	for i := 0; i < 5; i++ {
		dp[i] = inf
	}
	dp[0] = 0
	n := len(s)
	for i := 0; i < n; i++ {
		fp := make([]int, 5)
		for j := 0; j < 5; j++ {
			fp[j] = inf
		}
		for j := 0; j < 4; j++ {
			if s[i] == hard[j] {
				// 删除s[i] =>
				fp[j+1] = min(fp[j+1], dp[j])
			} else {
				fp[j] = min(fp[j], dp[j])
			}
			fp[j] = min(fp[j], dp[j]+a[i])
		}
		copy(dp, fp)
	}

	return min(dp[0], dp[1], dp[2], dp[3])
}

func solve1(s string, a []int) int {
	n := len(s)
	dp := make([]int, 4)
	for i := 0; i < 4; i++ {
		dp[i] = -inf
	}
	dp[0] = 0
	var sum int
	for i := 0; i < n; i++ {
		sum += a[i]
		if s[i] == 'h' {
			// 保留或者删除h
			// 删除h时，所有的都需要 + a[i]
			// 保留h时
			dp[1] = max(dp[1], dp[0]) + a[i]
			dp[2] += a[i]
			dp[3] += a[i]
			// 删除时，其他的不变
		} else if s[i] == 'a' {
			dp[2] = max(dp[2], dp[1]) + a[i]
			dp[0] += a[i]
			dp[3] += a[i]
		} else if s[i] == 'r' {
			dp[3] = max(dp[3], dp[2]) + a[i]
			dp[0] += a[i]
			dp[1] += a[i]
		} else if s[i] == 'd' {
			dp[0] += a[i]
			dp[1] += a[i]
			dp[2] += a[i]
		} else {
			dp[0] += a[i]
			dp[1] += a[i]
			dp[2] += a[i]
			dp[3] += a[i]
		}
	}

	return sum - max(dp[0], dp[1], dp[2], dp[3])
}
