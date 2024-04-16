package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)

	res := solve(n, k)

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

const mod = 998244353

func add(arr ...int) int {
	var res int
	for i := 0; i < len(arr); i++ {
		res += arr[i]
		if res >= mod {
			res -= mod
		}
	}
	return res
}

func solve(n int, k int) int {
	if k == 1 {
		// 0 or 1
		return 2
	}
	dp := make([][]int, 4)
	fp := make([][]int, 4)
	for i := 0; i < 4; i++ {
		dp[i] = make([]int, k+1)
		fp[i] = make([]int, k+1)
	}
	dp[0][1] = 1
	dp[1][2] = 1
	dp[2][2] = 1
	dp[3][1] = 1
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			// 和前面的有相同的颜色，所以不产生新的块
			fp[0][j] = add(dp[0][j], dp[1][j], dp[2][j], dp[3][j-1])
			fp[1][j] = add(dp[0][j-1], dp[1][j], dp[3][j-1])
			fp[2][j] = add(dp[0][j-1], dp[2][j], dp[3][j-1])
			if j >= 2 {
				fp[1][j] = add(fp[1][j], dp[2][j-2])
				fp[2][j] = add(fp[2][j], dp[1][j-2])
			}
			fp[3][j] = add(dp[0][j-1], dp[1][j], dp[2][j], dp[3][j])
		}
		for x := 0; x < 4; x++ {
			for j := 1; j <= k; j++ {
				dp[x][j] = fp[x][j]
				fp[x][j] = 0
			}
		}
	}

	var res int
	for x := 0; x < 4; x++ {
		res = add(res, dp[x][k])
	}

	return res
}
