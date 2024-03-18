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
	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		a := readNNums(reader, n*m)
		res := solve(n, m, a)

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

func solve(n int, m int, a []int) int {
	people := make([]Person, n*m)

	for i := 0; i < n*m; i++ {
		people[i] = Person{i, a[i]}
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].sight < people[j].sight || people[i].sight == people[j].sight && people[i].id < people[j].id
	})

	// 1, 1, 2, 1
	// 这里 3号的位置 是 4，这个是确定的，但是它却是在1前面进入的
	var id int

	seats := make([]int, n*m)

	for i := 0; i < n*m; {
		j := i
		for i < n*m && people[i].sight == people[j].sight {
			i++
		}
		cnt := i - j
		pos := id % m
		if pos > 0 {
			tmp := min(cnt, m-pos)
			for k := tmp - 1; k >= 0; k-- {
				seats[people[j+k].id] = id
				id++
			}
			j += tmp
			cnt -= tmp
		}
		// 现在到头部了
		if cnt > 0 {
			for j+m <= i {
				for k := m - 1; k >= 0; k-- {
					seats[people[j+k].id] = id
					id++
				}
				j += m
			}
			if j < i {
				for k := i - 1; k >= j; k-- {
					seats[people[k].id] = id
					id++
				}
			}
		}
	}

	tr := NewSegTree(n * m)
	var res int

	for i := 0; i < n*m; i++ {
		j := seats[i]
		r, c := j/m, j%m

		res += tr.Get(r*m, r*m+c)
		tr.Update(j, 1)
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Person struct {
	id    int
	sight int
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p] += v
	for p > 1 {
		t.arr[p>>1] = t.arr[p] + t.arr[p^1]
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) int {
	var res int
	l += t.sz
	r += t.sz

	for l < r {
		if l&1 == 1 {
			res += t.arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += t.arr[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}
