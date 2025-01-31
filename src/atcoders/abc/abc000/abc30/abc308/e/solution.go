package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	A := readNNums(reader, n)
	S := readString(reader)[:n]
	res := solve(A, S)
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

func solve(A []int, S string) int {
	n := len(A)
	// i, j, k s[i] = M, S[j] = 'E', S[k] = 'X'
	// 如果A[i]= 0, A[j] = 1, A[k] = 2, mex(i, j, k) = 3
	// dp[x][m] = 表示到目前位置x （按位表示的状态）的计数
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, 1<<3)
	}

	var ans int

	for i := 0; i < n; i++ {
		if S[i] == 'M' {
			dp[0][(1<<A[i])]++
		} else if S[i] == 'E' {
			for x := 0; x < 8; x++ {
				dp[1][x|(1<<A[i])] += dp[0][x]
			}
		} else {
			// S[i] = x
			// 0 not worth to add up
			for mex := 1; mex <= 3; mex++ {
				mask := (1 << mex) - 1
				for x := 0; x < 8; x++ {
					y := x | (1 << A[i])
					if (y>>mex)&1 == 1 || y&mask != mask {
						continue
					}
					ans += mex * dp[1][x]
				}
			}

		}
	}

	return ans
}
