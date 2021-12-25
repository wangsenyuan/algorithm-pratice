package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		fmt.Println(solve(n, A))
	}
}

const N_INF = -(1 << 60)

func solve(n int, A []int) int64 {
	dp := make([][][]int64, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int64, 31)
		for j := 0; j < 31; j++ {
			dp[i][j] = make([]int64, 31)
			for k := 0; k < 31; k++ {
				dp[i][j][k] = N_INF
			}
		}
	}

	for i := 0; i < n; i++ {
		x := A[i]
		c := digitCount(x)
		for j := 1; j < 31; j++ {
			dp[0][c][j] = max(dp[0][c][j], (int64(x)<<uint64(j))-int64(x))
		}
		for j := 1; j < 31; j++ {
			dp[1][j][c] = max(dp[1][j][c], int64(x)-(int64(x)<<uint64(j)))
		}
	}

	var ans int64

	for i := 1; i < 31; i++ {
		for j := 1; j < 31; j++ {
			ans = max(ans, dp[0][i][j]+dp[1][i][j])
		}
	}

	return ans
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func digitCount(num int) int {
	var res int
	for num > 0 {
		res++
		num >>= 1
	}
	return res
}
