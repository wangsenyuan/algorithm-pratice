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
		a := readNNums(reader, n)
		res := solve(a)
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

func solve(a []int) int {
	n := len(a)
	dp := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		dp[i] = make([]int, n+2)
		for j := i + 1; j <= n+1; j++ {
			dp[i][j] = inf
		}
	}

	for ln := 1; ln <= n; ln++ {
		for l := 0; l+ln <= n; l++ {
			if a[l]%2 != (l+1)%2 || a[l] > l+1 {
				continue
			}
			v := (l + 1 - a[l]) / 2
			r := l + ln
			for m := l + 1; m < r; m += 2 {
				if dp[l+1][m] <= v {
					w := max(v, dp[m+1][r]-(m-l+1)/2)
					dp[l][r] = min(dp[l][r], w)
				}
			}
		}
	}

	fp := make([]int, n+1)

	for r := 1; r <= n; r++ {
		fp[r] = fp[r-1]
		for l := 0; l < r; l++ {
			if dp[l][r] <= fp[l] {
				fp[r] = max(fp[r], fp[l]+(r-l)/2)
			}
		}
	}

	return fp[n]
}
