package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		segs := make([][]int, n)
		for i := 0; i < n; i++ {
			segs[i] = readNNums(reader, 3)
		}
		res := solve(segs)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(segs [][]int) int {
	// seg[i] = (ci, li, ri)
	n := len(segs)
	ps := make([]Point, 2*n)
	for i := 0; i < n; i++ {
		ps[i*2] = Point{segs[i][1], i, 0}
		ps[i*2+1] = Point{segs[i][2], i, 1}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].pos < ps[j].pos || ps[i].pos == ps[j].pos && ps[i].close < ps[j].close
	})

	dsu := NewDSU(n)

	sets := make([]map[int]int, 2)
	for i := 0; i < 2; i++ {
		sets[i] = make(map[int]int)
	}
	vis := make([]int, n)

	for _, p := range ps {
		id := p.id
		vis[id]++
		c := segs[id][0]
		if vis[id] == 1 {
			// merge it with other side
			x := -1
			xv := -1
			for k, v := range sets[1^c] {
				if xv < v || xv == v && x < k {
					x = k
					xv = v
				}
				dsu.Union(id, k)
			}
			if x >= 0 {
				// keep the farthest one
				sets[1^c] = make(map[int]int)
				sets[1^c][x] = xv
			}
			sets[c][id] = segs[id][2]
		} else {
			delete(sets[c], id)
		}
	}

	return dsu.Size()
}

type Point struct {
	pos   int
	id    int
	close int
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

func (this *DSU) Size() int {
	var ans int
	for i := 0; i < len(this.arr); i++ {
		if this.arr[i] == i {
			ans++
		}
	}
	return ans
}
