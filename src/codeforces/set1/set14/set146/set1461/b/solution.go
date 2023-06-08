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
		n, m := readTwoNums(reader)
		mat := make([]string, n)
		for i := 0; i < n; i++ {
			mat[i] = readString(reader)[:m]
		}
		res := solve(mat)

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

func solve(mat []string) int {
	// 如果 dp[r][i][j] 第r行，从 i..j (j - i + 1 is odd),
	// 存在一个高度h = （j - i + 2) / 2 的 tree
	// 则 dp[r+1][i-1][j+1] = dp[r][i][j] + (j + 1 - (i - 1) + 1) 都是*
	m := len(mat[0])
	dp := make([][]bool, m)
	fp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, m)
		fp[i] = make([]bool, m)
	}
	var res int
	for j := 0; j < m; j++ {
		if mat[0][j] == '*' {
			res++
			dp[j][j] = true
		}
	}

	for r := 1; r < len(mat); r++ {
		for i := 0; i < m; i++ {
			if mat[r][i] == '.' {
				continue
			}
			for j := i; j < m && mat[r][j] == '*'; j++ {
				if (j-i+1)%2 == 1 {
					if j == i {
						fp[i][j] = true
					} else if dp[i+1][j-1] {
						fp[i][j] = true
					}
				}
			}
		}
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				if fp[i][j] {
					res++
				}
				dp[i][j] = fp[i][j]
				fp[i][j] = false
			}
		}
	}

	return res
}
