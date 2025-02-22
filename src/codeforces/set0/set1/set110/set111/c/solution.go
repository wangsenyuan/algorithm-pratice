package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	res := solve(n, m)
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

const INF = 1 << 20

func solve(n int, m int) int {
	if n < m {
		n, m = m, n
	}
	if n == 1 {
		return 0
	}
	M := 1 << m
	dp := make([][]int, M)
	for i := 0; i < M; i++ {
		dp[i] = make([]int, M)
		for j := 0; j < M; j++ {
			dp[i][j] = INF
		}
	}

	for i := 0; i < M; i++ {
		dp[0][i] = digitCount(i)
	}

	fp := make([][]int, M)
	for i := 0; i < M; i++ {
		fp[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fp[i][j] = INF
		}
	}

	check := func(state int) bool {
		for i := 0; i < m; i++ {
			if (state>>i)&1 == 0 {
				return true
			}
		}
		return false
	}

	for r := 1; r < n; r++ {
		for c := 0; c < M; c++ {
			cnt := digitCount(c)
			for b := 0; b < M; b++ {
				for a := 0; a < M; a++ {
					if !check(a | b | c | (b << 1) | (b >> 1)) {
						fp[b][c] = min(fp[b][c], dp[a][b]+cnt)
					}
				}
			}
		}
		for a := 0; a < M; a++ {
			for b := 0; b < M; b++ {
				dp[a][b] = fp[a][b]
				fp[a][b] = INF
			}
		}
	}

	res := INF
	for a := 0; a < M; a++ {
		for b := 0; b < M; b++ {
			if !check(a | b | (b >> 1) | (b << 1)) {
				res = min(res, dp[a][b])
			}
		}
	}

	return n*m - res
}

func digitCount(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve1(n int, m int) int {
	if n < m {
		n, m = m, n
	}
	// n * m <= 40 => m <= 6, 2 * m <=12, 3 * m <= 18
	// 需要哪些状态信息,
	// 假如对于(i, j), 要么它不动，从其他地方 （上、下、左、右）移入蜘蛛
	// 要么它移走, 它如果向上、左、右移走，没有遗留问题，但是如果要向下方移动
	// 移动下方的状态还没有计算， 下方的某几位状态是确定的
	M := pow(3, m)
	dp := make([]int, M)
	arr := make([]int, m)
	// 0 表示该位需要向（上、左、右）移动，1表示向下移动、2表示不移动
	for state := 0; state < M; state++ {
		splitAsArray(state, arr)
		var cnt int
		for i := 0; i < m && cnt >= 0; i++ {
			x := arr[i]
			if x == 0 {
				if i > 0 && arr[i-1] == 2 {
					cnt++
				} else if i+1 < m && arr[i+1] == 2 {
					cnt++
				} else {
					cnt = -1
				}
			} else if x == 1 {
				cnt++
			}
		}

		dp[state] = cnt
	}
	cur := make([]int, m)
	fp := make([]int, M)
	for r := 1; r < n; r++ {
		for state := 0; state < M; state++ {
			fp[state] = -1
		}

		for state := 0; state < M; state++ {
			if dp[state] < 0 {
				continue
			}
			splitAsArray(state, arr)

			for next := 0; next < M; next++ {
				splitAsArray(next, cur)
				valid := true
				for i := 0; i < m && valid; i++ {
					if arr[i] == 1 {
						// the cell up need to move down
						valid = cur[i] == 2
					}
				}
				if !valid {
					continue
				}
				var cnt int
				for i := 0; i < m && valid; i++ {
					if cur[i] == 0 {
						// need to go up, left or right
						cnt++
						valid = arr[i] == 2 || (i > 0 && cur[i-1] == 2) || (i+1 < m && cur[i+1] == 2)
					} else if cur[i] == 1 {
						cnt++
					}
				}
				if !valid {
					continue
				}
				if fp[next] < dp[state]+cnt {
					fp[next] = dp[state] + cnt
				}
			}
		}
		copy(dp, fp)
	}

	var res int
	for state := 0; state < M; state++ {
		if dp[state] < 0 {
			continue
		}
		splitAsArray(state, arr)
		valid := true
		for i := 0; i < m && valid; i++ {
			valid = arr[i] != 1
		}
		if valid {
			res = max(res, dp[state])
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func splitAsArray(num int, arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = num % 3
		num /= 3
	}
}

func pow(a int, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r *= a
		}
		a *= a
		b >>= 1
	}
	return r
}
