package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	mat := make([]string, n)
	for i := 0; i < n; i++ {
		mat[i] = readString(reader)
	}

	res := solve(mat, k)

	fmt.Println(res)
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

func solve(mat []string, k int) string {
	n := len(mat)
	buf := make([]byte, 2*n-1)

	if k >= 2*n-1 {
		for i := 0; i < len(buf); i++ {
			buf[i] = 'a'
		}
		return string(buf)
	}

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	var d int

	if k > 0 {

		fp := make([][]int, n)

		getBest := func(i, j int) int {
			if i > 0 && j > 0 {
				return min(fp[i-1][j], fp[i][j-1])
			}
			if i == 0 && j == 0 {
				return 0
			}
			if i == 0 {
				return fp[i][j-1]
			}
			return fp[i-1][j]
		}

		for i := 0; i < n; i++ {
			fp[i] = make([]int, n)
			for j := 0; j < n; j++ {
				if mat[i][j] == 'a' {
					fp[i][j] = getBest(i, j)
				} else {
					fp[i][j] = getBest(i, j) + 1
				}
				if fp[i][j] <= k {
					d = max(d, i+j)
				}
			}
		}
		d++
		for i := 0; i < d; i++ {
			buf[i] = 'a'
		}
		for i := 0; i < n; i++ {
			j := (d - 1) - i
			if j >= 0 && j < n && fp[i][j] == k {
				dp[i][j] = true
			}
		}
	} else {
		buf[0] = mat[0][0]
		dp[0][0] = true
		d++
	}

	for ; d < len(buf); d++ {
		var best byte = 'z'
		for i := 0; i < n; i++ {
			j := d - i
			if j >= 0 && j < n {
				if i > 0 && dp[i-1][j] || j > 0 && dp[i][j-1] {
					if mat[i][j] < best {
						best = mat[i][j]
					}
				}
			}
		}

		buf[d] = best

		for i := 0; i < n; i++ {
			j := d - i
			if j >= 0 && j < n && mat[i][j] == best {
				if i > 0 && dp[i-1][j] || j > 0 && dp[i][j-1] {
					dp[i][j] = true
				}
			}
		}
	}

	return string(buf)
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
