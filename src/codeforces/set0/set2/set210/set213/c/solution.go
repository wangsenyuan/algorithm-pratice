package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = readNNums(reader, n)
	}

	res := solve(mat)

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

func solve(mat [][]int) int {
	n := len(mat)

	// dp[a][b] 表示在两个位置a,b时的最优值
	dp := make([][]int, n+1)
	fp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		fp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = -oo
			fp[i][j] = -oo
		}
	}

	get := func(a, b, c, d int) int {
		if a >= n || b >= n || c >= n || d >= n {
			// invalid
			return -oo
		}
		res := mat[a][b]
		if a != c || b != d {
			res += mat[c][d]
		}
		return res
	}

	dp[0][0] = mat[0][0]

	for step := 0; step < 2*n-2; {
		for a := max(0, step-n+1); a <= step && a < n; a++ {
			// a 是第一次的行号
			b := step - a
			// b = step - a < n => a > step - n
			for c := max(0, step-n+1); c <= step && c < n; c++ {
				d := step - c
				// 两次都横向移动
				fp[a][c] = max(fp[a][c], dp[a][c]+get(a, b+1, c, d+1))
				// 一次横向移动，一次垂直移动
				fp[a+1][c] = max(fp[a+1][c], dp[a][c]+get(a+1, b, c, d+1))
				fp[a][c+1] = max(fp[a][c+1], dp[a][c]+get(a, b+1, c+1, d))
				// 两次垂直移动
				fp[a+1][c+1] = max(fp[a+1][c+1], dp[a][c]+get(a+1, b, c+1, d))
			}
		}
		step++
		for a := max(0, step-n+1); a <= step && a < n; a++ {
			for b := max(0, step-n+1); b <= step && b < n; b++ {
				dp[a][b] = fp[a][b]
				fp[a][b] = -oo
			}
		}
	}

	return dp[n-1][n-1]
}
