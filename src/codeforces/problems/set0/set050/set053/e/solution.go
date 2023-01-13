package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 2)
	}
	res := solve(n, k, E)

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

func solve(n int, k int, E [][]int) int64 {
	// n <= 10
	// dp[state][x][j]， 使用 i个节点，以x为根，且有j个 dead ends的计数
	// dp[state | y][y][j+1] += dp[state][x][j] y 作为新的root
	// dp[s1 | s2][x][j] = dp[s1][x][j1] + dp[s2][y][j2] 以x为root， 如果x是dead end，或者y为dead ent， j 和j1, j2 的关系无法确定
	//  所以，还需要一个一个信息，是哪些是dead end
	// dp[s1][s2] 是用s1，节点，其中s2为dead end的计数,
	// dp[s1 | x][s2 | x - y] += dp[s1][s2]  如果 x 连接 y, y 属于 s2，且size(s1 | x) <= k
	//  或者 dp[s1|x][s2|x] += dp[s1][s2] y 不属于s2
	conn := make([][]int, n)
	for i := 0; i < n; i++ {
		conn[i] = make([]int, n)
	}

	tot := 1 << n

	dp := make([][]int64, tot)
	for i := 0; i < tot; i++ {
		dp[i] = make([]int64, tot)
		for j := 0; j < tot; j++ {
			dp[i][j] = -1
		}
		// dp[1<<i][1<<i] = 1
	}

	dp[0][0] = 0

	for i := 0; i < n; i++ {
		dp[1<<i][1<<i] = 1
	}

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		conn[u][v]++
		conn[v][u]++
		dp[(1<<u)|(1<<v)][(1<<u)|(1<<v)] = 1
	}

	var dfs func(x int, y int) int64

	dfs = func(x int, y int) int64 {
		if x&y != y || digitCount(y) < 2 {
			return 0
		}

		if dp[x][y] >= 0 {
			return dp[x][y]
		}

		dp[x][y] = 0

		for u := 0; u < n; u++ {
			if (y>>u)&1 == 1 {
				for v := 0; v < n; v++ {
					if conn[u][v] > 0 && (y>>v)&1 == 0 && (x>>v)&1 == 1 {
						// v is also in x
						// 检查v是不是一个新的dead end
						dp[x][y] += dfs(x^(1<<u), y^(1<<u)^(1<<v))
						// it is a dead end
						dp[x][y] += dfs(x^(1<<u), y^(1<<u))
					}
				}
			}
		}

		dp[x][y] /= int64(digitCount(y))

		return dp[x][y]
	}

	var res int64

	for state := 1; state < tot; state++ {
		d := digitCount(state)
		if d == k {
			res += dfs(tot-1, state)
		}
	}

	return res
}

func digitCount(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}
