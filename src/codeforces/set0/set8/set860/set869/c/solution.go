package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a, b, c := readThreeNums(reader)

	res := solve(a, b, c)

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

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(red int, blue int, purple int) int {
	// red-blue
	dp := make([]int, blue+1)
	// red-purple
	fp := make([]int, purple+1)

	dp[0] = 1
	fp[0] = 1
	for i := 1; i <= red; i++ {
		for j := min(i, blue); j > 0; j-- {
			// 当前红色节点不连接
			//dp[i][j] = dp[i-1][j]
			dp[j] = add(dp[j], mul(dp[j-1], blue-(j-1)))
		}
		for j := min(i, purple); j > 0; j-- {
			fp[j] = add(fp[j], mul(fp[j-1], purple-(j-1)))
		}
	}

	gp := make([]int, purple+1)
	gp[0] = 1
	for i := 1; i <= blue; i++ {
		for j := min(i, purple); j > 0; j-- {
			gp[j] = add(gp[j], mul(gp[j-1], purple-(j-1)))
		}
	}

	res := []int{0, 0, 0}

	for i := 0; i <= min(red, blue); i++ {
		res[0] = add(res[0], dp[i])
	}

	for i := 0; i <= min(red, purple); i++ {
		res[1] = add(res[1], fp[i])
	}

	for i := 0; i <= min(blue, purple); i++ {
		res[2] = add(res[2], gp[i])
	}

	return mul(res[0], mul(res[1], res[2]))
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
