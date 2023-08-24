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
		res := solve(n)
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

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

const MAX_X = 40004

var dp [MAX_X]int

func init() {
	var palindromeNums []int

	for x := 1; x < MAX_X; x++ {
		if x == reverse(x) {
			palindromeNums = append(palindromeNums, x)
		}
	}

	// dp[i][x] = dp[i-1][x] + dp[i][x-p[i]]
	// for dp[0][x] = 1
	// dp[i][1] = 1
	//dp := make([]int, MAX_X)

	for i := 0; i < MAX_X; i++ {
		dp[i] = 1
	}

	n := len(palindromeNums)
	for i := 1; i < n; i++ {
		// 需要计算出 dp[i][x-p[i]],
		// dp[x+p[i]] += dp[i]
		for x := palindromeNums[i]; x < MAX_X; x++ {
			dp[x] = add(dp[x], dp[x-palindromeNums[i]])
		}
		//dp[1] = add(dp[1], 1)
	}
}

func reverse(num int) int {
	var res int
	for num > 0 {
		res = res*10 + num%10
		num /= 10
	}
	return res
}

func solve(n int) int {
	return dp[n]
}
