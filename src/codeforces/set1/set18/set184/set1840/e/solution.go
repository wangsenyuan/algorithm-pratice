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
		s1 := readString(reader)
		s2 := readString(reader)
		x, p := readTwoNums(reader)
		queries := make([][]int, p)
		for i := 0; i < p; i++ {
			s, _ := reader.ReadBytes('\n')
			var y int
			pos := readInt(s, 0, &y)
			if y == 1 {
				queries[i] = make([]int, 2)
			} else if y == 2 {
				queries[i] = make([]int, 5)
			} else {
				queries[i] = make([]int, 1)
			}
			queries[i][0] = y
			for j := 1; j < len(queries[i]); j++ {
				pos = readInt(s, pos+1, &queries[i][j])
			}
		}
		res := solve(x, s1, s2, queries)
		for _, b := range res {
			if b {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
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

func solve(x int, s1 string, s2 string, queries [][]int) []bool {
	n := len(s1)
	ok := make([]bool, 4*n)

	bufs := [][]byte{
		[]byte(s1),
		[]byte(s2),
	}

	pullUp := func(i int) {
		ok[i] = ok[2*i+1] && ok[2*i+2]
	}

	var build func(i int, l int, r int)

	build = func(i int, l int, r int) {
		if l == r {
			ok[i] = bufs[0][l] == bufs[1][l]
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		pullUp(i)
	}

	build(0, 0, n-1)

	// block/un-block position p
	var block func(i int, l int, r int, p int, v bool)

	block = func(i int, l int, r int, p int, v bool) {
		if l == r {
			// l == p
			ok[i] = v || bufs[0][l] == bufs[1][l]
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			block(2*i+1, l, mid, p, v)
		} else {
			block(2*i+2, mid+1, r, p, v)
		}
		pullUp(i)
	}
	// position p already changed
	var update func(i int, l int, r int, p int)
	update = func(i int, l int, r int, p int) {
		if l == r {
			// l == p
			ok[i] = bufs[0][l] == bufs[1][l]
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			update(2*i+1, l, mid, p)
		} else {
			update(2*i+2, mid+1, r, p)
		}
		pullUp(i)
	}

	swap := func(a int, i int, b int, j int) {
		bufs[a-1][i-1], bufs[b-1][j-1] = bufs[b-1][j-1], bufs[a-1][i-1]
		update(0, 0, n-1, i-1)
		if i != j {
			update(0, 0, n-1, j-1)
		}
	}
	m := len(queries)
	unblock := make([]int, m+2)

	var res []bool

	for i := 1; i <= m; i++ {
		cur := queries[i-1]
		if unblock[i] > 0 {
			block(0, 0, n-1, unblock[i]-1, false)
		}
		if cur[0] == 1 {
			// do block
			block(0, 0, n-1, cur[1]-1, true)
			if i+x < len(unblock) {
				unblock[i+x] = cur[1]
			}
		} else if cur[0] == 2 {
			swap(cur[1], cur[2], cur[3], cur[4])
		} else {
			res = append(res, ok[0])
		}
	}

	return res
}
