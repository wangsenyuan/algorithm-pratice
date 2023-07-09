package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, k := readThreeNums(reader)
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = readNNums(reader, m)
	}
	res := solve(mat, k)
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

func solve(mat [][]int, k int) int {
	dp := make([]int, k)
	for i := 1; i < k; i++ {
		dp[i] = math.MinInt32
	}

	dp[0] = 0

	m := len(mat[0])
	h := m / 2
	sp := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		sp[i] = make([]int, k)
	}

	process := func(row []int) {
		// sp[i] = 在 rem = i时，所能获得的最大sum
		for i := 0; i <= h; i++ {
			for j := 0; j < k; j++ {
				sp[i][j] = math.MinInt32
			}
		}

		sp[0][0] = 0

		for i := 0; i < m; i++ {
			for j := h; j >= 1; j-- {
				for r := 0; r < k; r++ {
					nr := (r + row[i]) % k
					sp[j][nr] = max(sp[j][nr], sp[j-1][r]+row[i])
					sp[j][r] = max(sp[j][r], sp[j-1][r])
				}
			}
		}
	}

	// 70 * 70 * 70 * 70

	fp := make([]int, k)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < k; j++ {
			fp[j] = math.MinInt32
		}
		process(mat[i])
		for a := 0; a < k; a++ {
			for j := 0; j <= h; j++ {
				for r := 0; r < k; r++ {
					nr := (a + r) % k
					fp[nr] = max(fp[nr], dp[a]+sp[j][r])
				}
			}
		}
		copy(dp, fp)
	}
	return dp[0]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
