package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nums := readNNums(reader, 4)
	res := solve(nums[0], nums[1], nums[2], nums[3])
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

const mod = 100000000

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(n1 int, n2 int, k1 int, k2 int) int {
	dp := make([][][][]int, n1+1)
	for i := range n1 + 1 {
		dp[i] = make([][][]int, n2+1)
		for j := range n2 + 1 {
			dp[i][j] = make([][]int, 2)
			for u := range 2 {
				dp[i][j][u] = make([]int, max(k1, k2)+1)
			}
		}
	}
	dp[1][0][0][1] = 1
	dp[0][1][1][1] = 1

	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			if i+1 <= n1 {
				// 再安排一个footman
				// 之前都是horseman
				for x := 0; x <= k2; x++ {
					dp[i+1][j][0][1] = add(dp[i+1][j][0][1], dp[i][j][1][x])
				}
				for x := k1; x > 0; x-- {
					dp[i+1][j][0][x] = add(dp[i+1][j][0][x], dp[i][j][0][x-1])
				}
			}
			if j+1 <= n2 {
				for x := 0; x <= k1; x++ {
					dp[i][j+1][1][1] = add(dp[i][j+1][1][1], dp[i][j][0][x])
				}
				for x := k2; x > 0; x-- {
					dp[i][j+1][1][x] = add(dp[i][j+1][1][x], dp[i][j][1][x-1])
				}
			}
		}
	}
	var ans int
	for x := 0; x <= k1; x++ {
		ans = add(ans, dp[n1][n2][0][x])
	}
	for x := 0; x <= k2; x++ {
		ans = add(ans, dp[n1][n2][1][x])
	}

	return ans
}
