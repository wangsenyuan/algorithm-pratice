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

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(s string) int {
	// 有效的状态
	// 居然还有0
	n := len(s)
	if n == 1 {
		if s[0] == '?' {
			// 0 or *
			return 2
		}
		if s[0] == '*' || s[0] == '0' {
			return 1
		}
		return 0
	}
	// dp[0...3] dp[0/1/2]表示s[i] 是数字时的状态
	// dp[1] 时需要知道它是否和前面的匹配了 增加一个状态dp[3]
	// dp[4] 表示s[i] 是炸弹时的状态
	dp := make([]int, 5)
	if s[0] == '?' {
		dp[0] = 1
		dp[1] = 1
		dp[4] = 1
	}
	if s[0] == '0' {
		dp[0] = 1
	}
	if s[0] == '1' {
		dp[1] = 1
	}
	if s[0] == '2' {
		return 0
	}
	if s[0] == '*' {
		dp[4] = 1
	}
	for i := 1; i < n; i++ {
		fp := make([]int, 5)
		if s[i] == '?' || s[i] == '0' {
			// 前面是0，或者是已完结的1
			fp[0] = add(dp[0], dp[3])
		}
		if s[i] == '?' || s[i] == '1' {
			// 前面是0，或者是已完结的1
			// fp[1] = dp[0] + dp[3]
			fp[1] = add(dp[0], dp[3])
			// 前面只能是炸弹
			fp[3] = dp[4]
		}
		if s[i] == '?' || s[i] == '2' {
			// 前面必须是炸弹
			fp[2] = dp[4]
		}
		if s[i] == '?' || s[i] == '*' {
			// 1* 2*, **
			// 前面是未完结的1，2，或者是炸弹
			fp[4] = add(add(dp[1], dp[2]), dp[4])
		}
		dp = fp
	}
	var ans int
	for i := 0; i < 5; i++ {
		if i == 2 || i == 1 {
			continue
		}
		ans = add(ans, dp[i])
	}
	return ans
}
