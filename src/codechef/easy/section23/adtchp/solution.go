package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer
	n, m, q := readThreeNums(reader)
	Q := make([][]int, q)
	for i := 0; i < q; i++ {
		Q[i] = readNNums(reader, 5)
	}
	res := solve(n, m, Q)

	for i := 0; i < len(res); i++ {
		if res[i] {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

var dd = []int{-1, 0, 1, 0, -1}

func solve(n int, m int, Q [][]int) []bool {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
	}

	var q int
	for _, cur := range Q {
		t := cur[0]
		if t == 2 {
			q++
			continue
		}
		x1, y1, x2, y2 := cur[1]-1, cur[2]-1, cur[3]-1, cur[4]-1
		for i := x1; i <= x2; i++ {
			grid[i][y1] = 1
			grid[i][y2] = 1
		}
		for j := y1; j <= y2; j++ {
			grid[x1][j] = 1
			grid[x2][j] = 1
		}
	}

	set := NewDSU(n * m)

	fun := func(x, y int) {
		if grid[x][y] == 0 {
			if x > 0 && grid[x-1][y] == 0 {
				set.Union(x*m+y, (x-1)*m+y)
			}
			if x+1 < n && grid[x+1][y] == 0 {
				set.Union(x*m+y, (x+1)*m+y)
			}
			if y > 0 && grid[x][y-1] == 0 {
				set.Union(x*m+y, x*m+y-1)
			}
			if y+1 < m && grid[x][y+1] == 0 {
				set.Union(x*m+y, x*m+y+1)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fun(i, j)
		}
	}

	cancel := func(x1, y1, x2, y2 int) {
		for i := x1; i <= x2; i++ {
			grid[i][y1] = 0
			fun(i, y1)
			grid[i][y2] = 0
			fun(i, y2)
		}
		for j := y1; j <= y2; j++ {
			grid[x1][j] = 0
			fun(x1, j)
			grid[x2][j] = 0
			fun(x2, j)
		}
	}

	ans := make([]bool, q)

	for i := len(Q) - 1; i >= 0; i-- {
		cur := Q[i]
		x1, y1, x2, y2 := cur[1]-1, cur[2]-1, cur[3]-1, cur[4]-1

		if cur[0] == 1 {
			cancel(x1, y1, x2, y2)
			continue
		}
		a := x1*m + y1
		b := x2*m + y2
		q--
		ans[q] = set.Find(a) == set.Find(b)
	}

	return ans
}
