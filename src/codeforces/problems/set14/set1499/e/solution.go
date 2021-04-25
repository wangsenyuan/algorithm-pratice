package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	x, _ := reader.ReadString('\n')
	y, _ := reader.ReadString('\n')

	res := solve(normalize(x), normalize(y))
	buf.WriteString(fmt.Sprintf("%d\n", res))
	fmt.Print(buf.String())
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

func normalize(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] < 'a' || s[i] > 'z' {
			return s[:i]
		}
	}
	return s
}

const MOD = 998244353

func solve(x, y string) int {
	m := len(x)
	n := len(y)
	dp := make([][][][]int64, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([][][]int64, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([][]int64, 2)
			for k := 0; k < 2; k++ {
				dp[i][j][k] = make([]int64, 4)
			}
		}
	}
	var res int64
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i < m {
				modAdd(&dp[i+1][j][0][1], 1)
			}
			if j < n {
				modAdd(&dp[i][j+1][1][2], 1)
			}
			for mask := 0; mask < 4; mask++ {
				if 0 < i && i < m && x[i-1] != x[i] {
					modAdd(&dp[i+1][j][0][mask|1], dp[i][j][0][mask])
				}
				if i < m && 0 < j && y[j-1] != x[i] {
					modAdd(&dp[i+1][j][0][mask|1], dp[i][j][1][mask])
				}

				if 0 < j && j < n && y[j-1] != y[j] {
					modAdd(&dp[i][j+1][1][mask|2], dp[i][j][1][mask])
				}

				if 0 < i && j < n && x[i-1] != y[j] {
					modAdd(&dp[i][j+1][1][mask|2], dp[i][j][0][mask])
				}
			}
			modAdd(&res, dp[i][j][0][3])
			modAdd(&res, dp[i][j][1][3])
		}
	}

	return int(res)
}

func modAdd(a *int64, b int64) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}
