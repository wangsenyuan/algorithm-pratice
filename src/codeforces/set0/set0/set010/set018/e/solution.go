package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, _ := readTwoNums(reader)
	grid := make([]string, n)
	for i := 0; i < n; i++ {
		grid[i] = readString(reader)
	}
	cnt, res := solve(grid)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", cnt))
	for _, row := range res {
		buf.WriteString(row)
		buf.WriteByte('\n')
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

func solve(grid []string) (int, []string) {
	n := len(grid)
	m := len(grid[0])
	if m == 1 {
		return solve1(grid)
	}

	check := func(a int, b int, i int) int {
		var res int
		for j := 0; j < m; j++ {
			x := int(grid[i][j] - 'a')
			if j&1 == 0 && x != a || j&1 == 1 && x != b {
				res++
			}
		}
		return res
	}

	// dp[i][j] = 不是i*j时的最小值
	dp := make([][]int, 26)
	fp := make([][][]int, n)
	for i := 0; i < n; i++ {
		fp[i] = make([][]int, 26)
		for j := 0; j < 26; j++ {
			fp[i][j] = make([]int, 26)
		}
	}

	for i := 0; i < 26; i++ {
		dp[i] = make([]int, 26)
	}
	col := make([]int, 26)
	set := func() {
		for j := 0; j < 26; j++ {
			col[j] = inf
		}
	}

	for r := 0; r < n; r++ {
		for a := 0; a < 26; a++ {
			for b := 0; b < 26; b++ {
				fp[r][a][b] = inf
				if a != b {
					fp[r][a][b] = check(a, b, r) + dp[a][b]
				}
			}
		}

		set()
		for a := 0; a < 26; a++ {
			tmp := inf
			for b := 0; b < 26; b++ {
				dp[a][b] = tmp
				tmp = min(tmp, col[b])
				col[b] = min(tmp, fp[r][a][b])
			}
		}
		set()
		for a := 0; a < 26; a++ {
			tmp := inf
			for b := 25; b >= 0; b-- {
				dp[a][b] = min(dp[a][b], tmp)
				tmp = min(tmp, col[b])
				col[b] = min(tmp, fp[r][a][b])
			}
		}
		set()
		for a := 25; a >= 0; a-- {
			tmp := inf
			for b := 0; b < 26; b++ {
				dp[a][b] = min(dp[a][b], tmp)
				tmp = min(tmp, col[b])
				col[b] = min(tmp, fp[r][a][b])
			}
		}
		set()
		for a := 25; a >= 0; a-- {
			tmp := inf
			for b := 25; b >= 0; b-- {
				dp[a][b] = min(dp[a][b], tmp)
				tmp = min(tmp, col[b])
				col[b] = min(tmp, fp[r][a][b])
			}
		}
	}

	best := inf
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			best = min(best, fp[n-1][i][j])
		}
	}

	res := make([]string, n)

	var restore func(r int, a, b int, expect int)

	restore = func(r int, a, b int, expect int) {
		if r < 0 {
			return
		}
		for x := 0; x < 26; x++ {
			for y := 0; y < 26; y++ {
				if x != a && y != b && fp[r][x][y] == expect {
					cur := make([]byte, m)
					var cnt int
					for j := 0; j < m; j++ {
						if j&1 == 0 {
							cur[j] = byte(x + 'a')
						} else {
							cur[j] = byte(y + 'a')
						}
						if cur[j] != grid[r][j] {
							cnt++
						}
					}
					res[r] = string(cur)
					restore(r-1, x, y, expect-cnt)
					return
				}
			}
		}
	}

	restore(n-1, -1, -1, best)

	return best, res
}

func solve1(grid []string) (int, []string) {
	var change int
	n := len(grid)
	res := make([]string, n)
	for i := 0; i < n; {
		j := i
		for i < n && grid[i][0] == grid[j][0] {
			res[i] = grid[i]
			i++
		}
		tmp := "a"
		if (i-j)%2 == 1 {
			// 修改中间的
			if grid[j] == "a" {
				tmp = "b"
			}
		} else {
			// 偶数
			if grid[j] == "a" || (i < n && grid[i] == "a") {
				tmp = "b"
			}
			if grid[j] == "b" || (i < n && grid[i] == "b") {
				tmp = "c"
			}
		}
		for k := j + 1; k < i; k += 2 {
			res[k] = tmp
		}

		change += (i - j) / 2
	}

	return change, res
}
