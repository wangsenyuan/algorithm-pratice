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

var fs []bool

func init() {
	fs = make([]bool, 5010)
}

func solve(a []int) int {
	n := len(a)
	sum := make([][]Pair, n+1)
	vis := make([]bool, n+1)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		sum[i] = make([]Pair, 0, 1)
		dp[i] = make([]bool, 2*n+1)
	}

	for i := 0; i < n; i++ {
		if a[i] == 0 {
			sum[i+1] = append(sum[i+1], Pair{i + 1, 1})
		}
	}

	for i := 1; i <= n; i++ {
		var mex int
		for j := i; j <= n; j++ {
			vis[a[j-1]] = true
			for vis[mex] {
				mex++
			}
			if mex > a[i-1] && a[j-1] < a[i-1] {
				sum[j] = append(sum[j], Pair{i, mex})
				break
			}
		}
		copy(vis, fs)
		mex = 0
		for j := i; j > 0; j-- {
			vis[a[j-1]] = true
			for vis[mex] {
				mex++
			}
			if mex > a[i-1] && a[j-1] < a[i-1] {
				sum[i] = append(sum[i], Pair{j, mex})
				break
			}
		}
		copy(vis, fs)
	}

	dp[0][0] = true

	for i := 1; i <= n; i++ {
		copy(dp[i], dp[i-1])
		for _, j := range sum[i] {
			for z := 1; z <= 2*n; z++ {
				if z^j.second <= 2*n && dp[j.first-1][z^j.second] {
					dp[i][z] = true
				}
			}
		}
	}

	for x := 2 * n; x > 0; x-- {
		if dp[n][x] {
			return x
		}
	}

	return 0
}

type Pair struct {
	first  int
	second int
}
