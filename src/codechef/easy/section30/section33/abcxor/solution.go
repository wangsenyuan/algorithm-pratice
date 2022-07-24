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
		res := solve(n, k)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

var twos = []int{3, 5, 6}
var ones = []int{1, 2, 4}

func solve(n int, k int) int64 {
	// (2 ^^ n - 1) * (2 ^^ n - 2)
	x := int64(1 << uint(n))

	return (x - 1) * (x - 2)
}

func solve1(n int, k int) int64 {
	// 如果k的第i位位0，则 a, b, c 为 0, 0, 0, 1, 1, 0
	//             1, 则 ....      0, 0, 1, 1, 1, 1
	// 不考虑distinct, 很简单，
	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, 8)
	}
	// a = b, b = c, & a = c
	dp[0][7] = 1
	for i := 0; i < n; i++ {
		x := (k >> uint(i)) & 1
		var nums []int
		if x == 0 {
			nums = twos
		} else {
			nums = ones
		}
		// 0, 0, 0, or 0, 1, 1
		for j := 0; j < 8; j++ {
			// a = b, b = c, a = c
			dp[i+1][j] += dp[i][j]
			for _, u := range nums {
				a := (u >> 2) & 1
				b := (u >> 1) & 1
				c := u & 1
				var nj int
				if (j>>2)&1 == 1 {
					// a == b
					if a == b {
						nj |= 1 << 2
					}
				}
				if (j>>1)&1 == 1 {
					// b == c
					if b == c {
						nj |= 1 << 1
					}
				}
				if j&1 == 1 {
					// a == c
					if a == c {
						nj |= 1
					}
				}
				dp[i+1][nj] += dp[i][j]
			}
		}
	}

	return dp[n][0]
}
