package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		A := readNNums(scanner, n)
		for i := 0; i < n-1; i++ {
			scanner.Scan()
		}
		fmt.Println(solve(n, A, nil))
	}
}

func solve(n int, A []int, edges [][]int) int {
	if n == 1 {
		return 0
	}
	var cnt [3]int
	for i := 0; i < n; i++ {
		cnt[A[i]]++
	}
	if cnt[0] == n || cnt[1] == n || cnt[2] == n {
		return 0
	}

	if cnt[1] == 0 {
		return 2
	}

	if cnt[0] == 0 || cnt[2] == 0 {
		return 1
	}

	conn := make([][]int, n)
	for i := 0; i < n; i++ {
		conn[i] = make([]int, 0, 10)
	}

	for _, edge := range edges {
		a, b := edge[0]-1, edge[1]-1
		conn[a] = append(conn[a], b)
		conn[b] = append(conn[b], a)
	}
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, cnt[0]+1)
		for j := 0; j <= cnt[0]; j++ {
			dp[i][j] = make([]int, 3)
		}
	}
	tmp := make([]int, cnt[0]+1)
	var dfs func(p, u int)

	dfs = func(p, u int) {
		for _, v := range conn[u] {
			if p == v {
				continue
			}
			dfs(u, v)
		}
		for i := 0; i <= cnt[0]; i++ {
			dp[u][i][0] = 1e9 + 7
			dp[u][i][1] = 1e9 + 7
			dp[u][i][2] = 1e9 + 7
		}

		dp[u][1][0] = 0
		dp[u][0][1] = 1
		dp[u][0][2] = 0

		for _, v := range conn[u] {
			if p == v {
				continue
			}
			for i := 0; i <= cnt[0]; i++ {
				tmp[i] = 1e9 + 7
			}
			for i := 0; i <= cnt[0]; i++ {
				for j := 0; j <= i; j++ {
					tmp[i] = min(tmp[i], dp[u][j][0]+min(dp[v][i-j][0], dp[v][i-j][1]))
				}
			}
			for i := 0; i <= cnt[0]; i++ {
				dp[u][i][0] = tmp[i]
				tmp[i] = 1e9 + 7
			}
			for i := 0; i <= cnt[0]; i++ {
				for j := 0; j <= i; j++ {
					tmp[i] = min(tmp[i], dp[u][j][2]+min(dp[v][i-j][2], dp[v][i-j][1]))
				}
			}
			for i := 0; i <= cnt[0]; i++ {
				dp[u][i][2] = tmp[i]
				tmp[i] = 1e9 + 7
			}

			for i := 0; i <= cnt[0]; i++ {
				for j := 0; j <= i; j++ {
					tmp[i] = min(tmp[i], dp[u][j][1]+min3(dp[v][i-j][0], dp[v][i-j][1], dp[v][i-j][2]))
				}
			}
			for i := 0; i <= cnt[0]; i++ {
				dp[u][i][1] = tmp[i]
				tmp[i] = 1e9 + 7
			}
		}
	}

	dfs(-1, 0)

	for c := 0; c < 3; c++ {
		if dp[0][cnt[0]][c] <= cnt[1] {
			return 1
		}
	}
	return 2
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}
