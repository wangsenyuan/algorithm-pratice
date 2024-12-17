package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

const RGB = "RGB"

const inf = 1 << 30

func solve(s string, k int) int {
	n := len(s)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 3)
	}

	best := inf

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			// dp[i][j]
			nj := (j + 1) % 3
			if s[i] != RGB[nj] {
				dp[i+1][nj] = dp[i][j] + 1
			} else {
				dp[i+1][nj] = dp[i][j]
			}
		}
		if i >= k-1 {
			i1 := i - k + 1
			for j := 0; j < 3; j++ {
				nj := (j + k) % 3
				tmp := dp[i+1][nj] - dp[i1][j]
				best = min(best, tmp)
			}
		}
	}

	return best
}
