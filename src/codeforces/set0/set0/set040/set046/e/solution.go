package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = readNNums(reader, m)
	}

	res := solve(grid)

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

const inf = 1 << 60

func solve(grid [][]int) int {
	// n := len(grid)
	m := len(grid[0])
	dp := make([]int, m)
	fp := make([]int, m)

	for i, row := range grid {
		if i%2 == 1 {
			s := sum(row)
			fp[m-1] = -inf
			for j := m - 2; j >= 0; j-- {
				s -= row[j+1]
				fp[j] = s + dp[j+1]
			}
			for j := 1; j < m; j++ {
				fp[j] = max(fp[j], fp[j-1])
			}
			copy(dp, fp)
		} else {
			s := row[0]
			if i == 0 {
				dp[0] = s
			} else {
				// 否则它是无效的
				dp[0] = -inf
			}
			for j := 1; j < m; j++ {
				s += row[j]
				dp[j] = s + fp[j-1]
			}
			for j := m - 2; j >= 0; j-- {
				dp[j] = max(dp[j], dp[j+1])
			}
		}
	}
	return slices.Max(dp)
}

func sum(arr []int) int {
	var res int
	for _, x := range arr {
		res += x
	}
	return res
}
