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
		n, x := readTwoNums(reader)
		records := make([][]int, n)
		for i := 0; i < n; i++ {
			records[i] = readNNums(reader, 2)
		}
		res := solve(x, records)

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

func solve(x int, records [][]int) int {
	// m := len(records)
	// dp[i] 表示alice手头有
	// 这个算法逻辑上是成立的
	// 但是它的复杂性过高
	// 是因为每个节点，相当于存储了用或者不用的信息
	// 如果要获得h的happiness最少需要多少钱？
	// dp[i][j] = dp[i-1][j-h[i]] + c[i] 如果 dp[i][j] <= (i+1) * x
	//
	var sh int
	for _, rec := range records {
		sh += rec[1]
	}

	dp := make([]int, sh+1)
	ok := make([]bool, sh+1)
	for i := 0; i <= sh; i++ {
		dp[i] = inf
	}
	dp[0] = 0
	ok[0] = true

	for i, rec := range records {
		c, h := rec[0], rec[1]
		for j := sh; j >= h; j-- {
			if dp[j-h]+c <= i*x && ok[j-h] {
				dp[j] = min(dp[j], dp[j-h]+c)
				ok[j] = true
			}
		}
	}

	for j := sh; j > 0; j-- {
		if ok[j] {
			return j
		}
	}

	return 0
}
