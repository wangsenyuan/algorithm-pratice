package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer
	n, q := readTwoNums(reader)
	Q := make([]string, q)
	for i := 0; i < q; i++ {
		Q[i], _ = reader.ReadString('\n')
		Q[i] = Q[i][:len(Q[i])-1]
	}
	res := solve(n, Q)
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}
	fmt.Print(buf.String())
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

const MOD = 1000000007

func solve(C int, Q []string) []int {
	m := len(Q)
	// stack[0] 的top是最上面一层
	// stack[1] 的bottom是最上面一层，top是最下面一层
	stack := make([][]string, 2)
	sz := make([]int, 2)
	dp := make([][][][]int, 2)
	for i := 0; i < 2; i++ {
		stack[i] = make([]string, m)
		dp[i] = make([][][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([][]int, C)
			for k := 0; k < C; k++ {
				dp[i][j][k] = make([]int, C)
			}
		}
	}

	update := func(id int) {
		n := sz[id]
		if n == 1 {
			for i := 0; i < C; i++ {
				for j := 0; j < C; j++ {
					if i == j {
						dp[id][0][i][j] = 1
					}
				}
			}
			return
		}

		for i := 0; i < C; i++ {
			for j := 0; j < C; j++ {
				dp[id][n-1][i][j] = 0
			}
			for j := i - 1; j <= i+1; j++ {
				if j < 0 || j >= C || stack[id][n-1][i] != stack[id][n-2][j] {
					continue
				}
				for k := 0; k < C; k++ {
					dp[id][n-1][i][k] += dp[id][n-2][j][k]
					if dp[id][n-1][i][k] >= MOD {
						dp[id][n-1][i][k] -= MOD
					}
				}
			}
		}
	}
	res := make([]int, 0, m)
	for _, cur := range Q {
		if cur[0] == 'a' {
			stack[1][sz[1]] = cur[4:]
			sz[1]++
			update(1)
			continue
		}
		if cur[0] == 'r' {
			if sz[0] == 0 {
				for sz[1] > 0 {
					stack[0][sz[0]] = stack[1][sz[1]-1]
					sz[0]++
					sz[1]--
					update(0)
				}
			}
			sz[0]--
			continue
		}
		// query
		ss := strings.Split(cur, " ")
		c, _ := strconv.Atoi(ss[1])
		d, _ := strconv.Atoi(ss[2])
		c--
		d--
		if sz[1] == 0 {
			res = append(res, dp[0][sz[0]-1][c][d])
		} else if sz[0] == 0 {
			res = append(res, dp[1][sz[1]-1][d][c])
		} else {
			// both not empty
			var tmp int64
			for k := 0; k < C; k++ {
				for k2 := k - 1; k2 <= k+1; k2++ {
					if k < 0 || k >= C || stack[0][0][k] != stack[1][0][k2] {
						continue
					}
					tmp += int64(dp[0][sz[0]-1][c][k]) * int64(dp[1][sz[1]-1][d][k2]) % MOD
					tmp %= MOD
				}
			}
			res = append(res, int(tmp))
		}
	}

	return res
}
