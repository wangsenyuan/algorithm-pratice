package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = readNNums(reader, m)
	}
	res := solve(mat)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(mat [][]int) int {
	// n * n * n * n TLE

	m := len(mat)
	n := len(mat[0])
	sum := make([]int, n)
	prev := make([]int, n+1)
	mv := make([]int, n)

	stack := make([]int, n)
	L := make([]int, n)

	var best int
	// a,b是矩形的上、下边界
	for a := 0; a < m; a++ {
		// row[i] = 第i列，从a到b行间的和
		// min[i] ..................最小值
		for i := 0; i < n; i++ {
			sum[i] = 0
			mv[i] = 301
		}
		for b := a; b < m; b++ {
			for i := 0; i < n; i++ {
				sum[i] += mat[b][i]

				prev[i+1] = prev[i] + sum[i]

				mv[i] = min(mv[i], mat[b][i])
			}
			// 如果以 mat[b][i]为最小值，结果是多少
			// 这时需要知道i的左右边界
			var it int
			for i := 0; i < n; i++ {
				for it > 0 && mv[stack[it-1]] >= mv[i] {
					it--
				}
				L[i] = -1
				if it > 0 {
					L[i] = stack[it-1]
				}
				stack[it] = i
				it++
			}
			it = 0
			for i := n - 1; i >= 0; i-- {
				for it > 0 && mv[stack[it-1]] >= mv[i] {
					it--
				}
				R := n
				if it > 0 {
					R = stack[it-1]
				}
				tmp := prev[R] - prev[L[i]+1]

				best = max(best, tmp*mv[i])

				stack[it] = i
				it++
			}
		}
	}

	return best
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
