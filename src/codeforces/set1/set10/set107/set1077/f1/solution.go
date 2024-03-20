package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k, x := readThreeNums(reader)

	a := readNNums(reader, n)

	res := solve(a, k, x)

	fmt.Println(res)
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

const oo = 1 << 60

func solve(a []int, k int, x int) int {
	n := len(a)

	dp := make([][][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, x+1)
		for j := 0; j <= x; j++ {
			dp[i][j] = make([]int, k+1)
			for l := 0; l <= k; l++ {
				dp[i][j][l] = -oo
			}
		}
	}

	dp[0][0][0] = 0
	var p int
	for i := 0; i < n; i++ {
		for j := 0; j <= x; j++ {
			for l := 0; l <= k; l++ {
				dp[p^1][j][l] = -oo
			}
		}
		for j := 0; j <= x && j <= i; j++ {
			for l := 0; l < k; l++ {
				// 如果不选择当前位置
				if l+1 < k {
					dp[p^1][j][l+1] = max(dp[p^1][j][l+1], dp[p][j][l])
				}
				// 如果选择a[i]
				if j+1 <= x {
					dp[p^1][j+1][0] = max(dp[p^1][j+1][0], dp[p][j][l]+a[i])
				}
			}
		}
		p ^= 1
	}

	res := -1

	for l := 0; l < k; l++ {
		res = max(res, dp[p][x][l])
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
