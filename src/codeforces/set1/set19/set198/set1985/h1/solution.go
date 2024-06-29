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
		n, m := readTwoNums(reader)
		grid := make([]string, n)
		for i := 0; i < n; i++ {
			grid[i] = readString(reader)[:m]
		}
		res := solve(grid)
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
	// 先计算不添加时的最大值
	var res int

	res = max(res, f(grid))
	res = max(res, f(transpose(grid)))

	return res
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

func f(grid []string) int {
	// 从上到下
	n := len(grid)
	m := len(grid[0])

	set := NewDSU(n * m)
	dp := make([][]Pair, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]Pair, m)
		for j := 0; j < m; j++ {
			dp[i][j] = Pair{0, -1}
		}
		for j := 0; j < m; j++ {
			// 不考虑下面的情况
			if grid[i][j] == '#' {
				if i > 0 && grid[i-1][j] == '#' {
					set.Union(i*m+j, (i-1)*m+j)
				}
				if j > 0 && grid[i][j-1] == '#' {
					set.Union(i*m+j, i*m+j-1)
				}
			}
		}
		for j := 0; j < m; j++ {
			if grid[i][j] == '#' {
				p := set.Find(i*m + j)
				dp[i][j] = Pair{set.cnt[p], p}
			}
		}
	}

	set.Reset()
	// from bottom to top
	fp := make([]Pair, m)
	for j := 0; j < m; j++ {
		fp[j] = Pair{0, -1}
	}
	var ans int
	vis := make(map[int]bool)
	for i := n - 1; i >= 0; i-- {

		// 如果这一行被改变
		clear(vis)
		tmp := m
		for j := 0; j < m; j++ {
			if i > 0 && !vis[dp[i-1][j].second] {
				tmp += dp[i-1][j].first
				vis[dp[i-1][j].second] = true
			}
		}
		clear(vis)
		for j := 0; j < m; j++ {
			if !vis[fp[j].second] {
				tmp += fp[j].first
				vis[fp[j].second] = true
			}
		}

		ans = max(ans, tmp)

		for j := 0; j < m; j++ {
			if grid[i][j] == '#' {
				if i+1 < n && grid[i+1][j] == '#' {
					set.Union(i*m+j, (i+1)*m+j)
				}
				if j+1 < m && grid[i][j+1] == '#' {
					set.Union(i*m+j, i*m+j+1)
				}
			}
		}
		for j := 0; j < m; j++ {
			if grid[i][j] == '#' {
				p := set.Find(i*m + j)
				fp[j] = Pair{set.cnt[p], p}
			} else {
				fp[j] = Pair{0, -1}
			}
		}
	}

	return ans
}

type Pair struct {
	first  int
	second int
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}

func (this *DSU) Reset() {
	for i := 0; i < len(this.arr); i++ {
		this.arr[i] = i
		this.cnt[i] = 1
	}
}
