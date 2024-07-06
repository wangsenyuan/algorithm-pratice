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
		arr := readNNums(reader, 4)
		res := solve(arr)
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

const P = 201

var dp [P][P][P]int

func init() {
	dp[0][0][0] = 1
	for i := 0; i < P; i++ {
		for j := 0; j < P; j++ {
			for k := 0; k < P; k++ {
				var add int
				x := i & 1
				if j&1 == x && k&1 == x {
					add++
				}
				if i > 0 {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k]+add)
				}
				if j > 0 {
					dp[i][j][k] = max(dp[i][j][k], dp[i][j-1][k]+add)
				}
				if k > 0 {
					dp[i][j][k] = max(dp[i][j][k], dp[i][j][k-1]+add)
				}
			}
		}
	}
}

func solve(arr []int) int {
	a, b, c, d := arr[0], arr[1], arr[2], arr[3]
	// 将d降低到1的过程中
	// - 最后一次1
	// 减去d = 0的那次
	return (d+2)/2 + dp[a][b][c] - 2
}
