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
		n, k := readTwoNums(reader)
		s := readString(reader)[:n]
		res := solve(s, k)
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

const inf = 1 << 30

func solve(s string, k int) int {
	n := len(s)

	pref := make([]int, n)
	for i := 0; i < n; i++ {
		pref[i] = int(s[i] - '0')
		if i > 0 {
			pref[i] += pref[i-1]
		}
	}

	// dp[i] = i点亮且满足条件时的最少操作次数
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 {
			// 把前面的都给turn off 掉
			dp[i] = pref[i-1]
		}
		dp[i] += check(s[i] == '0')
		// i是第一个被点亮的
		if i-k >= 0 {
			dp[i] = min(dp[i], check(s[i] == '0')+dp[i-k]+pref[i-1]-pref[i-k])
		}
	}

	// dp[n-1]肯定是大于0的，因为这里要求s[n-1]点亮了，但是最后一个，即使删除掉了，也是符合条件的
	ans := dp[n-1] - check(s[n-1] == '0')

	for i := 0; i < n-1; i++ {
		tmp := dp[i] + pref[n-1] - pref[i]
		ans = min(ans, tmp)
	}

	return ans
}

func check(b bool) int {
	if b {
		return 1
	}
	return 0
}
