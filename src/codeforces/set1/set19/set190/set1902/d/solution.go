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

	n, m := readTwoNums(reader)
	s := readString(reader)[:n]
	qs := make([][]int, m)
	for i := 0; i < m; i++ {
		qs[i] = readNNums(reader, 4)
	}

	res := solve(s, qs)

	var buf bytes.Buffer
	for _, x := range res {
		if x {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(s string, qs [][]int) []bool {
	n := len(s)
	ps := make([]Point, n+1)
	ps[0] = Point{0, 0}

	for i := 0; i < n; i++ {
		ps[i+1] = ps[i].Add(getValue(s[i] == 'R')-getValue(s[i] == 'L'), getValue(s[i] == 'U')-getValue(s[i] == 'D'))
	}

	mp := make(map[Point][]int)

	for i := 0; i <= n; i++ {
		if len(mp[ps[i]]) == 0 {
			mp[ps[i]] = make([]int, 0, 1)
		}
		mp[ps[i]] = append(mp[ps[i]], i)
	}

	check := func(pt Point, l int, r int) bool {
		if vs, ok := mp[pt]; !ok {
			return false
		} else {
			i := sort.SearchInts(vs, l)
			return i < len(vs) && vs[i] <= r
		}
	}

	ans := make([]bool, len(qs))

	for i, cur := range qs {
		x, y, l, r := cur[0], cur[1], cur[2], cur[3]
		nx := ps[r].x + ps[l-1].x - x
		ny := ps[r].y + ps[l-1].y - y

		ans[i] = check(Point{x, y}, 0, l-1) || check(Point{nx, ny}, l, r-1) || check(Point{x, y}, r, n)
	}

	return ans
}

type Point struct {
	x int
	y int
}

func (this Point) Add(dx int, dy int) Point {
	return Point{this.x + dx, this.y + dy}
}

func getValue(b bool) int {
	if b {
		return 1
	}
	return 0
}
