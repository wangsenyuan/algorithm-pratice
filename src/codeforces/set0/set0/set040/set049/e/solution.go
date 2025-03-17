package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	s1 := readString(reader)
	s2 := readString(reader)
	n := readNum(reader)
	substitutions := make([]string, n)
	for i := range n {
		substitutions[i] = readString(reader)
	}
	return solve(s1, s2, substitutions)
}

type pair struct {
	first  int
	second int
}

func solve(s1 string, s2 string, substitutions []string) int {
	g := make([][]pair, 26)
	for _, cur := range substitutions {
		x := int(cur[0] - 'a')
		y := int(cur[3] - 'a')
		z := int(cur[4] - 'a')
		g[x] = append(g[x], pair{y, z})
	}

	for i := range 26 {
		g[i] = sortAndUnique(g[i])
	}

	get := func(s string) [][][]bool {
		n := len(s)
		dp := make([][][]bool, 26)
		for i := range 26 {
			dp[i] = make([][]bool, n)
			for j := range n {
				dp[i][j] = make([]bool, n)
			}
		}
		for i := range n {
			x := int(s[i] - 'a')
			dp[x][i][i] = true
		}
		// n * n * 26 * n * n = 1.6 * 1e8
		for r := range n {
			for l := r - 1; l >= 0; l-- {
				for x := range 26 {
					for _, cur := range g[x] {
						y, z := cur.first, cur.second
						for i := l; i < r; i++ {
							if dp[y][l][i] && dp[z][i+1][r] {
								dp[x][l][r] = true
								break
							}
						}
						if dp[x][l][r] {
							break
						}
					}
				}
			}
		}
		return dp
	}

	dp := get(s1)
	fp := get(s2)
	// l1, l2, r1, r2
	n := len(s1)
	m := len(s2)

	dist := make([][]int, n+1)
	for i := range n + 1 {
		dist[i] = make([]int, m+1)
		for j := range m + 1 {
			dist[i][j] = -1
		}
	}
	dist[0][0] = 0
	que := make([]int, (n+1)*(m+1))
	var head, tail int
	que[head] = 0
	head++
	for tail < head {
		i, j := que[tail]/m, que[tail]%m
		tail++
		if i == n && j == m {
			return dist[i][j]
		}
		// dist[i][j]
		if i < n && j < m && s1[i] == s2[j] && dist[i+1][j+1] < 0 {
			dist[i+1][j+1] = dist[i][j] + 1
			que[head] = (i+1)*m + (j + 1)
			head++
		}
		if i < n && j < m {
			for r1 := i; r1 < n; r1++ {
				for r2 := j; r2 < m; r2++ {
					for x := range 26 {
						if dp[x][i][r1] && fp[x][j][r2] && dist[r1+1][r2+1] < 0 {
							dist[r1+1][r2+1] = dist[i][j] + 1
							que[head] = (r1+1)*m + r2 + 1
							head++
						}
						if dist[r1+1][r2+1] > 0 {
							break
						}
					}
				}
			}
		}
	}

	return dist[n][m]
}

const inf = 1 << 60

func sortAndUnique(arr []pair) []pair {
	slices.SortFunc(arr, func(a, b pair) int {
		if a.first != b.first {
			return a.first - b.first
		}
		return a.second - b.second
	})
	var n int
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i] != arr[i-1] {
			arr[n] = arr[i-1]
			n++
		}
	}
	return arr[:n]
}
