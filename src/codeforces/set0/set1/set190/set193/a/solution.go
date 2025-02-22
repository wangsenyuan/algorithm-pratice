package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	G := make([]string, n)
	for i := 0; i < n; i++ {
		G[i] = readString(reader)[:m]
	}
	res := solve(G)
	fmt.Println(res)
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

var dir = []int{-1, 0, 1, 0, -1}

func solve(G []string) int {
	// it is already connected
	m := len(G)
	n := len(G[0])
	blocks := make([][]int, 0, m*n)
	id := make([][]int, m)

	for i := 0; i < m; i++ {
		id[i] = make([]int, n)
		for j := 0; j < n; j++ {
			id[i][j] = -1
			if G[i][j] == '#' {
				id[i][j] = len(blocks)
				blocks = append(blocks, []int{i, j})
			}
		}
	}

	if len(blocks) <= 2 {
		return -1
	}

	if len(blocks) == 3 {
		return 1
	}

	uf := NewUF(len(blocks))

	check := func(i int) bool {
		var p int
		for j := 0; j < len(blocks); j++ {
			if i == j {
				continue
			}
			x, y := blocks[j][0], blocks[j][1]
			for k := 0; k < 4; k++ {
				u, v := x+dir[k], y+dir[k+1]
				if u >= 0 && u < m && v >= 0 && v < n && id[u][v] >= 0 && id[u][v] != i {
					uf.Union(j, id[u][v])
				}
			}
			p = j
		}
		p = uf.Find(p)
		ok := uf.cnt[p] < len(blocks)-1
		uf.Reset()
		return ok
	}

	for i := 0; i < len(blocks); i++ {
		if check(i) {
			return 1
		}
	}

	return 2
}

type UF struct {
	arr []int
	cnt []int
}

func NewUF(n int) *UF {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i]++
	}
	return &UF{arr, cnt}
}

func (uf *UF) Find(x int) int {
	if uf.arr[x] != x {
		uf.arr[x] = uf.Find(uf.arr[x])
	}
	return uf.arr[x]
}

func (uf *UF) Union(a, b int) bool {
	a = uf.Find(a)
	b = uf.Find(b)
	if a != b {
		if uf.cnt[a] < uf.cnt[b] {
			a, b = b, a
		}
		uf.cnt[a] += uf.cnt[b]
		uf.arr[b] = a
		return true
	}
	return false
}

func (uf *UF) Reset() {
	for i := 0; i < len(uf.arr); i++ {
		uf.arr[i] = i
		uf.cnt[i] = 1
	}
}
