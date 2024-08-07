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
	rgs := getRegions(grid)
	n := len(grid)
	m := len(grid[0])
	R := make([]int, n+2)
	C := make([]int, m+2)
	RC := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		RC[i] = make([]int, m+2)
	}

	fr := make([]int, n)
	fc := make([]int, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '.' {
				fr[i]++
				fc[j]++
			}
		}
	}

	for _, rg := range rgs {
		t := max(0, rg.t-1)
		b := min(n-1, rg.b+1)
		l := max(0, rg.l-1)
		r := min(m-1, rg.r+1)
		R[t] += rg.sz
		R[b+1] -= rg.sz
		C[l] += rg.sz
		C[r+1] -= rg.sz
		RC[t][l] += rg.sz
		RC[b+1][l] -= rg.sz
		RC[t][r+1] -= rg.sz
		RC[b+1][r+1] += rg.sz
	}

	for i := 1; i <= n; i++ {
		R[i] += R[i-1]
	}

	for i := 1; i <= m; i++ {
		C[i] += C[i-1]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i > 0 {
				RC[i][j] += RC[i-1][j]
			}
			if j > 0 {
				RC[i][j] += RC[i][j-1]
			}
			if i > 0 && j > 0 {
				RC[i][j] -= RC[i-1][j-1]
			}
		}
	}

	var best int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tmp := R[i] + C[j] - RC[i][j] + fr[i] + fc[j]
			if grid[i][j] == '.' {
				tmp--
			}
			best = max(best, tmp)
		}
	}

	return best
}

func getRegions(grid []string) []*region {
	n := len(grid)
	m := len(grid[0])

	set := NewDSU(n * m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '#' {
				if i > 0 && grid[i-1][j] == '#' {
					set.Union((i-1)*m+j, i*m+j)
				}
				if j > 0 && grid[i][j-1] == '#' {
					set.Union(i*m+j-1, i*m+j)
				}
			}
		}
	}

	var regions []*region

	id := make([][]int, n)

	for i := 0; i < n; i++ {
		id[i] = make([]int, m)
		for j := 0; j < m; j++ {
			id[i][j] = -1
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '#' {
				r := set.Find(i*m + j)
				if id[r/m][r%m] == -1 {
					// first time
					id[r/m][r%m] = len(regions)
					regions = append(regions, &region{id[r/m][r%m], i, j, i, j, 1})
				} else {
					regions[id[r/m][r%m]].update(i, j)
				}
			}
		}
	}
	return regions
}

type pair struct {
	first  int
	second int
}

type region struct {
	id int
	t  int
	r  int
	b  int
	l  int
	sz int
}

func (this *region) update(i int, j int) {
	this.t = min(this.t, i)
	this.b = max(this.b, i)
	this.r = max(this.r, j)
	this.l = min(this.l, j)
	this.sz++
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
