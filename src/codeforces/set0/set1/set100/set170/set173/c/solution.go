package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(grid [][]int) int {
	//
	n := len(grid)
	m := len(grid[0])

	sum := make([][]int, n)
	for i := 0; i < n; i++ {
		sum[i] = make([]int, m)
		for j := 0; j < m; j++ {
			sum[i][j] = grid[i][j]
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}

	get := func(a, b, c, d int) int {
		res := sum[c][d]
		if a > 0 {
			res -= sum[a-1][d]
		}
		if b > 0 {
			res -= sum[c][b-1]
		}
		if a > 0 && b > 0 {
			res += sum[a-1][b-1]
		}
		return res
	}

	kk := min(n, m)

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, 3)
			dp[i][j][0] = grid[i][j]
		}
	}
	res := math.MinInt32

	// 1 is the cell it-self
	// 3 need to calc
	for i := 0; i+3 <= n; i++ {
		for j := 0; j+3 <= m; j++ {
			tmp := get(i, j, i+2, j+2)
			tmp -= get(i+1, j, i+1, j+1)
			dp[i][j][1] = tmp
			res = max(res, tmp)
		}
	}
	// f[i,j,k] = x + f[i + 2, j + 2, k - 4]
	// 1， 3， 5， 7， 9
	// 0，1， 2，
	// k = 2 * i + 1
	// i = (k - 1) / 2

	for k := 5; k <= kk; k += 2 {
		x := ((k - 1) / 2) % 3
		// x
		y := (x + 1) % 3
		for i := 0; i+k <= n; i++ {
			for j := 0; j+k <= m; j++ {
				tmp := dp[i+2][j+2][y]
				tmp += get(i, j, i+k-1, j+k-1) - get(i+1, j+1, i+k-2, j+k-2)
				tmp -= grid[i+1][j]
				tmp += grid[i+2][j+1]
				dp[i][j][x] = tmp
				res = max(res, tmp)
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
