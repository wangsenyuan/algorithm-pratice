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
		s := readString(reader)
		res := solve(s)
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

var nums = []int{1, 10, 100, 1000, 10000}

const inf = 1 << 60

func solve(s string) int {
	// ABCDE, 如果有一个E在很后面，那前面都是负值
	// 对于E来说，要改它的话，还要考虑把它改成D，C，B，A？
	// 而且只能改最后一个
	dp := make([][][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, 5)
		for j := 0; j < 5; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	n := len(s)
	var ptr int

	for i := n - 1; i >= 0; i-- {
		x := int(s[i] - 'A')
		for j := 0; j < 5; j++ {
			for k := 0; k < 2; k++ {
				dp[1^ptr][j][k] = -inf
			}
		}

		for j := 0; j < 5; j++ {
			for k := 0; k < 2; k++ {
				// dp[ptr][j][k]
				// if we don't change it
				if x < j {
					dp[1^ptr][j][k] = max(dp[1^ptr][j][k], -nums[x]+dp[ptr][j][k])
				} else {
					dp[1^ptr][x][k] = max(dp[1^ptr][x][k], nums[x]+dp[ptr][j][k])
				}
				if k == 0 {
					// have option to change it,
					for y := 0; y < 5; y++ {
						// change x to y
						sign := 1
						if y < j {
							sign = -1
						}
						dp[1^ptr][max(j, y)][1] = max(dp[1^ptr][max(j, y)][1], sign*nums[y]+dp[ptr][j][k])
					}
				}
			}
		}
		ptr ^= 1
	}

	var res = -inf
	for i := 0; i < 5; i++ {
		res = max(res, max(dp[ptr][i][0], dp[ptr][i][1]))
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
