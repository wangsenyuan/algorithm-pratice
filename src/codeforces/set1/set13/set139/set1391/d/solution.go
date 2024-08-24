package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	grid := make([]string, n)

	for i := 0; i < n; i++ {
		grid[i] = readString(reader)[:m]
	}

	res := solve(grid)

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

func solve(grid []string) int {
	n := len(grid)
	m := len(grid[0])
	if n >= 4 && m >= 4 {
		return -1
	}
	if n == 1 || m == 1 {
		return 0
	}
	if n > m {
		grid = transpose(grid)
		n, m = m, n
	}
	// n <= 3
	if n == 2 {
		return solve2(grid)
	}
	return solve3(grid)
}

func transpose(grid []string) []string {
	n := len(grid)
	m := len(grid[0])
	buf := make([][]byte, m)
	for i := 0; i < m; i++ {
		buf[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			buf[j][i] = grid[i][j]
		}
	}
	res := make([]string, m)
	for i := 0; i < m; i++ {
		res[i] = string(buf[i])
	}
	return res
}

const inf = 1 << 30

func solve2(grid []string) int {
	n := len(grid[0])
	dp := make([]int, 4)
	fp := make([]int, 4)

	count := func(state int, i int) int {
		var res int
		for j := 0; j < 2; j++ {
			if (state>>j)&1 != int(grid[j][i]-'0') {
				res++
			}
		}
		return res
	}

	for state := 0; state < 4; state++ {
		dp[state] = count(state, 0)
	}

	for i := 1; i < n; i++ {
		for j := 0; j < 4; j++ {
			fp[j] = inf
		}
		for state := 0; state < 4; state++ {
			// 不能有偶数个1出来
			c1 := bits.OnesCount(uint(state))
			for cur := 0; cur < 4; cur++ {
				c2 := bits.OnesCount(uint(cur))
				if (c1+c2)%2 == 0 {
					// invalid
					continue
				}
				fp[cur] = min(fp[cur], dp[state]+count(cur, i))
			}
		}
		copy(dp, fp)
	}

	return slices.Min(dp)
}

func solve3(grid []string) int {
	// len(grid) = 3
	n := len(grid[0])

	count := func(state int, i int) int {
		var res int
		for j := 0; j < 3; j++ {
			if (state>>j)&1 != int(grid[j][i]-'0') {
				res++
			}
		}
		return res
	}

	dp := make([]int, 8)
	for state := 0; state < 8; state++ {
		dp[state] = count(state, 0)
	}

	fp := make([]int, 8)
	check := func(x int, y int) bool {
		cnt := bits.OnesCount(uint(x)) + bits.OnesCount(uint(y))
		return cnt%2 != 0
	}

	for i := 1; i < n; i++ {
		for j := 0; j < 8; j++ {
			fp[j] = inf
		}
		for state := 0; state < 8; state++ {
			for cur := 0; cur < 8; cur++ {
				if check(state&6, cur&6) && check(state&3, cur&3) {
					cnt := count(cur, i)
					fp[cur] = min(fp[cur], dp[state]+cnt)
				}
			}
		}
		copy(dp, fp)
	}

	return slices.Min(dp)
}
