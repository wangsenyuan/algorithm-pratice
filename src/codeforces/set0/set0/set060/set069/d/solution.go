package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	nums := readNNums(reader, 4)
	x, y, n, d := nums[0], nums[1], nums[2], nums[3]

	vectors := make([][]int, n)

	for i := 0; i < n; i++ {
		vectors[i] = readNNums(reader, 2)
	}

	res := solve(x, y, d, vectors)

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

func solve(x int, y int, d int, vectors [][]int) string {

	dp := make([][]int, 2*d+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2*d+1)
	}

	var dfs func(x int, y int) bool

	checkDist := func(x int, y int) bool {
		return x*x+y*y <= d*d
	}

	dfs = func(x int, y int) bool {
		if dp[x+d][y+d] != 0 {
			return dp[x+d][y+d] == 1
		}

		dp[x+d][y+d] = -1

		for _, cur := range vectors {
			dx, dy := cur[0], cur[1]
			nx, ny := x+dx, y+dy
			if checkDist(nx, ny) && !dfs(nx, ny) {
				// 可以安全得到达一个失败点
				dp[x+d][y+d] = 1
				break
			}
		}

		return dp[x+d][y+d] == 1
	}

	res := dfs(x, y)
	if res {
		return "Anton"
	}
	return "Dasha"
}

func abs(num int) int {
	return max(num, -num)
}
