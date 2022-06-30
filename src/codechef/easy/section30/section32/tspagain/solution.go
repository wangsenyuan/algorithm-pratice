package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	// hint(105)
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for i := 0; i < tc; i++ {
		n, k := readTwoNums(reader)
		C := make([][]int, k)
		for j := 0; j < k; j++ {
			C[j] = readNNums(reader, 2)
		}
		res := solve(n, C)
		buf.WriteString(fmt.Sprintf("%d\n", res))
		if i+1 < tc {
			readString(reader)
		}
	}

	fmt.Print(buf.String())
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
		if s[i] == '\n' {
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

func solve1(n int, C [][]int) int {
	cost := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		cost[i] = make([]int, n+1)
	}
	for _, c := range C {
		cost[c[0]][c[1]]++
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if i+1 < n {
				cost[i][j] += cost[i+1][j]
			}
			if j-1 >= 0 {
				cost[i][j] += cost[i][j-1]
			}
			if i+1 < n && j-1 >= 0 {
				cost[i][j] -= cost[i+1][j-1]
			}
		}
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	// 以l, r为头的两条路径 (只能选择比 min(l, r) 小的节点)，且没有重复节点
	var fn func(l, r int) int

	fn = func(l, r int) int {
		if dp[l][r] >= 0 {
			return dp[l][r]
		}

		k := min(l, r) - 1

		if k == 0 {
			dp[l][r] = cost[0][l] + cost[r][0]
		} else {
			dp[l][r] = min(cost[k][l]+fn(k, r), fn(l, k)+cost[r][k])
		}
		return dp[l][r]
	}

	return fn(n-1, n-1)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve(n int, C [][]int) int {
	cost := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		cost[i] = make([]int, n+1)
	}
	for _, c := range C {
		cost[c[0]][c[1]]++
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if i+1 < n {
				cost[i][j] += cost[i+1][j]
			}
			if j-1 >= 0 {
				cost[i][j] += cost[i][j-1]
			}
			if i+1 < n && j-1 >= 0 {
				cost[i][j] -= cost[i+1][j-1]
			}
		}
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	// dp[i][j] 以 i，j为头的两条路径，且路径没有重复

	for l := 1; l < n; l++ {
		for r := 1; r < n; r++ {
			// dp[l][r]
			k := min(l, r) - 1
			if k == 0 {
				dp[l][r] = cost[0][l] + cost[r][0]
			} else {
				dp[l][r] = min(cost[k][l]+dp[k][r], dp[l][k]+cost[r][k])
			}
		}
	}

	return dp[n-1][n-1]
}
