package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		mat := make([]string, 7)
		for i := 0; i < 7; i++ {
			mat[i] = readString(reader)
		}

		res := solve(mat)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

const inf = 1 << 30

func solve(mat []string) int {
	var flag uint64
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			if mat[i][j] == 'B' {
				flag ^= (1 << (i*7 + j))
			}
		}
	}

	flip := func(v uint64, r int, c int) uint64 {
		return v ^ (1 << (r*7 + c))
	}

	get := func(v uint64, r int, c int) bool {
		return (v>>(r*7+c))&1 == 1
	}

	var odd []int
	var even []int

	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			if (i+j)%2 == 1 {
				odd = append(odd, i*7+j)
			} else {
				even = append(even, i*7+j)
			}
		}
	}

	valid := func(g uint64, pairty int) bool {
		for r := 1; r < 6; r++ {
			for c := 1; c < 6; c++ {
				if (r+c)%2 == pairty && get(g, r, c) {
					if get(g, r-1, c-1) && get(g, r-1, c+1) && get(g, r+1, c-1) && get(g, r+1, c+1) {
						return false
					}
				}
			}
		}

		return true
	}

	var check func(g uint64, flipsLeft int, idx int, vec []int, parity int) bool

	check = func(g uint64, flipsLeft int, idx int, vec []int, parity int) bool {
		if flipsLeft == 0 {
			return valid(g, parity)
		}

		if idx == len(vec) {
			return false
		}

		if check(g, flipsLeft, idx+1, vec, parity) {
			return true
		}

		r, c := vec[idx]/7, vec[idx]%7

		return check(flip(g, r, c), flipsLeft-1, idx+1, vec, parity)
	}

	var res int

	for i := 0; i <= 4; i++ {
		if check(flag, i, 0, odd, 1) {
			res += i
			break
		}
	}

	for i := 0; i <= 4; i++ {
		if check(flag, i, 0, even, 0) {
			res += i
			break
		}
	}

	return res
}

func solve1(mat []string) int {
	n := 16
	N := 1 << n
	dp := make([]int, N)

	get := func(r int, c int) int {
		if mat[r][c] == 'W' {
			return 0
		}
		return 1
	}

	for state := 0; state < N; state++ {
		for i := 0; i < n; i++ {
			r, c := i/7, i%7
			x := (state >> i) & 1
			y := get(r, c)

			if x != y {
				dp[state]++
			}
		}
	}

	// dp[state] 是在保证不违反规则下的最优值
	fp := make([]int, N)
	for pos := 16; pos < 49; pos++ {
		for state := 0; state < N; state++ {
			fp[state] = inf
		}
		r, c := pos/7, pos%7
		y := get(r, c)
		for state := 0; state < N; state++ {
			for x := 0; x < 2; x++ {
				// r >= 2
				if x == 1 {
					if c-2 >= 0 && state&1 == 1 && (state>>8)&1 == 1 && (state>>14)&1 == 1 && (state>>2)&1 == 1 {
						continue
					}
				}
				var cnt int
				if x != y {
					cnt++
				}
				next := (state >> 1) | (x << 15)
				fp[next] = min(fp[next], dp[state]+cnt)
			}
		}
		copy(dp, fp)
	}

	res := inf
	for state := 0; state < N; state++ {
		res = min(res, dp[state])
	}

	return res
}
