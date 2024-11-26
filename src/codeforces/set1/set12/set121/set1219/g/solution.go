package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	a := make([][]int, n)
	for i := range n {
		a[i] = readNNums(reader, m)
	}
	return solve(a)
}

func solve(a [][]int) int {
	ta := transpose(a)

	res := max(solve1(a), solve1(ta))

	if len(a) > 4 && len(ta) > 4 {
		if len(a) >= len(ta) {
			// 迭代行
			res = max(res, solve2(a))
		} else {
			res = max(res, solve2(ta))
		}
	}

	return res
}

func transpose(a [][]int) [][]int {
	n := len(a)
	m := len(a[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res[j][i] = a[i][j]
		}
	}
	return res
}

func solve1(a [][]int) int {
	// 选择4行，或者 3列一行
	n := len(a)
	m := len(a[0])

	row := make([]int, n)
	col := make([]int, m)
	var sum int

	arr := make([]int, 4)
	add := func(num int) {
		for i := 0; i < 4; i++ {
			if num >= arr[i] {
				num, arr[i] = arr[i], num
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			row[i] += a[i][j]
			col[j] += a[i][j]
			sum += a[i][j]
		}
		add(row[i])
	}
	if n <= 4 || m <= 4 {
		return sum
	}
	// 选择4行
	res := arr[0] + arr[1] + arr[2] + arr[3]

	for i := 0; i < n; i++ {
		clear(arr)
		for j := 0; j < m; j++ {
			add(col[j] - a[i][j])
		}
		res = max(res, arr[0]+arr[1]+arr[2]+row[i])
	}

	return res
}

func solve2(a [][]int) int {
	n := len(a)
	m := len(a[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
	}

	row := make([]int, n)
	col := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			col[j] += a[i][j]
			row[i] += a[i][j]
		}
	}

	var ans int

	for i := 0; i < n; i++ {
		if i > 0 {
			var best int
			for l := 0; l < m; l++ {
				for r := l + 1; r < m; r++ {
					tmp := dp[l][r] + col[l] - a[i][l] + col[r] - a[i][r]
					best = max(best, tmp)
				}
			}
			ans = max(ans, best+row[i])
		}
		for l := 0; l < m; l++ {
			for r := l + 1; r < m; r++ {
				dp[l][r] = max(dp[l][r], row[i]-a[i][l]-a[i][r])
			}
		}
	}

	return ans
}
